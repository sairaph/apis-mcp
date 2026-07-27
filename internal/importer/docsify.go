package importer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

var (
	docsifyGitHubBrowseURL = regexp.MustCompile(`https://github\.com/([A-Za-z0-9_.-]+)/([A-Za-z0-9_.-]+)/(?:blob|tree)/([A-Za-z0-9_.+@=!$&(),;%-]+)/([^'"` + "`" + `\s]*)`)
	docsifyGitHubRawURL    = regexp.MustCompile(`https://raw\.githubusercontent\.com/([A-Za-z0-9_.-]+)/([A-Za-z0-9_.-]+)/([A-Za-z0-9_.+@=!$&(),;%-]+)/([^'"` + "`" + `\s]*)`)
	docsifyJSDelivrURL     = regexp.MustCompile(`https://cdn\.jsdelivr\.net/gh/([A-Za-z0-9_.-]+)/([A-Za-z0-9_.-]+)@([A-Za-z0-9_.+@=!$&(),;%-]+)/([^'"` + "`" + `\s]*)`)
	docsifyBasePath        = regexp.MustCompile(`\bbasePath\s*:`)
	docsifyStaticBasePath  = regexp.MustCompile(`\bbasePath\s*:\s*['"]([^'"]+)['"]`)
	docsifyBasePathToken   = regexp.MustCompile(`\bbasePath\b`)
	docsifyObjectConfig    = regexp.MustCompile(`window\.\$docsify\s*=\s*\{`)
	docsifyConfigAssign    = regexp.MustCompile(`window\.\$docsify\s*=`)
	docsifyConfigMutation  = regexp.MustCompile(`window\.\$docsify\s*\.`)
	docsifyConfigSpread    = regexp.MustCompile(`\.\.\.`)
	docsifyObjectID        = regexp.MustCompile(`^[0-9a-fA-F]{40}(?:[0-9a-fA-F]{24})?$`)
)

const (
	docsifyMaxSelections    = 64
	docsifyMaxRefCandidates = 32
)

type docsifyGitHubSelection struct {
	owner string
	repo  string
	ref   string
	path  string
	exact bool
}

type docsifyGitHubFile struct {
	owner  string
	repo   string
	ref    string
	commit string
	path   string
	size   int64
}

type docsifyRepositoryKey struct {
	owner string
	repo  string
	ref   string
}

type docsifyGitHubSnapshot struct {
	commit string
	tree   string
}

type docsifyDocument struct {
	title       string
	description string
	source      string
	identity    string
	output      string
	pagePath    string
	body        string
	front       map[string]any
}

// ImportDocsify pins and exhausts the Git-backed source Markdown advertised by
// a Docsify shell. It never executes the shell or its plugins.
func ImportDocsify(ctx context.Context, name, version, source string, options Options) (Result, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return Result{}, err
	}
	if strings.TrimSpace(name) == "" || strings.TrimSpace(version) == "" {
		return Result{}, errors.New("Docsify import requires API name and version")
	}
	shellURL, err := url.Parse(strings.TrimSpace(source))
	if err != nil || shellURL.Host == "" || shellURL.Scheme != "http" && shellURL.Scheme != "https" {
		return Result{}, errors.New("Docsify import source must be an HTTP(S) URL")
	}
	reader := newSourceReader(options)
	rawShell, provenance, err := reader.read(ctx, shellURL.String(), nil)
	if err != nil {
		return Result{}, err
	}
	document, err := parseHTML(rawShell)
	if err != nil || !isHTMLDocument(document) || !looksLikeDocsify(document) {
		return Result{}, errors.New("source is not a supported Docsify shell")
	}
	finalShell, err := url.Parse(provenance)
	if err != nil {
		return Result{}, fmt.Errorf("parse Docsify shell URL: %w", err)
	}
	// HTTP does not carry fragments, but query-based Docsify configuration does.
	finalShell.Fragment = shellURL.Fragment
	selections, err := docsifyGitHubSources(document, finalShell)
	if err != nil {
		return Result{}, err
	}
	reportProgress(options, Progress{Stage: "catalog", Framework: "docsify", URL: provenance, Queued: len(selections)})
	files, err := docsifyGitHubInventory(ctx, selections, reader)
	if err != nil {
		return Result{}, fmt.Errorf("Docsify inventory did not complete: %w", err)
	}
	if len(files) == 0 {
		return Result{}, errors.New("Docsify inventory contains no Markdown files")
	}
	totalFiles := len(files)
	truncated := false
	if options.MaxHTMLPages >= 0 && len(files) > options.MaxHTMLPages {
		files = files[:options.MaxHTMLPages]
		truncated = true
	}

	documents := make([]docsifyDocument, 0, len(files))
	for index, file := range files {
		if err := ctx.Err(); err != nil {
			return Result{}, err
		}
		sourceURL := docsifyRawGitHubURL(file)
		raw, resolved, readErr := reader.readFromOrigin(ctx, sourceURL, nil, "https://raw.githubusercontent.com")
		if readErr != nil {
			return Result{}, fmt.Errorf("fetch Docsify Markdown %s: %w", file.path, readErr)
		}
		if resolved != sourceURL {
			return Result{}, fmt.Errorf("Docsify Markdown source redirected: %s", sourceURL)
		}
		raw = bytes.TrimPrefix(raw, []byte{0xef, 0xbb, 0xbf})
		if !utf8.Valid(raw) || bytes.IndexByte(raw, 0) >= 0 {
			return Result{}, fmt.Errorf("Docsify Markdown %s is not valid UTF-8 text", file.path)
		}
		body := strings.ReplaceAll(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\r", "\n")
		front := make(map[string]any)
		if strings.HasPrefix(body, "---\n") {
			candidate := make(map[string]any)
			if parsedBody, parseErr := splitFrontmatter([]byte(body), &candidate); parseErr == nil {
				body = parsedBody
				front = candidate
			}
		}
		title := ""
		if existing, ok := front["title"].(string); ok {
			title = strings.TrimSpace(existing)
		}
		if title == "" {
			title = docsifyHeading(body)
		}
		if title == "" {
			title = docsifyTitle(file.path)
		}
		description := ""
		if existing, ok := front["description"].(string); ok {
			description = strings.TrimSpace(existing)
		}
		identity := file.owner + "/" + file.repo + "/" + file.ref + "/" + file.path
		outputPath := file.path
		if strings.EqualFold(path.Base(outputPath), "_index.md") {
			outputPath = path.Join(path.Dir(outputPath), "_index.docsify-"+strings.TrimPrefix(stableID("source-", identity), "source-")+".md")
		}
		output := filepath.Join("documentation", SafeSlug(file.owner), SafeSlug(file.repo), file.commit, filepath.FromSlash(outputPath))
		pagePath := filepath.ToSlash(filepath.Dir(output))
		documents = append(documents, docsifyDocument{
			title: title, description: description, source: resolved, identity: identity, output: output, pagePath: pagePath, body: body, front: front,
		})
		reportProgress(options, Progress{Stage: "page", Framework: "docsify", URL: resolved, Pages: len(documents), Queued: len(files) - index - 1, Truncated: truncated})
	}
	if err := validateDocsifyOutputPaths(documents); err != nil {
		return Result{}, err
	}

	result, err := publish(ctx, options, name, version, func(stage string) error {
		metadata := manifest{
			Name: name, Version: version, Collections: options.Collections,
			SourceRoot: provenance, SourceType: "docsify", ImportedFrom: provenance, Sources: totalFiles,
		}
		manifestBody := "This document set was imported from a complete Docsify source inventory."
		if truncated {
			manifestBody = "This document set was imported from a deliberately truncated Docsify source inventory."
		}
		if err := writeCanonicalFile(stage, "_index.md", metadata, manifestBody); err != nil {
			return err
		}
		for _, document := range documents {
			front := document.front
			if upstream, ok := front["source"]; ok && fmt.Sprint(upstream) != document.source {
				front["upstream_source"] = upstream
			}
			front["title"] = document.title
			if document.description != "" {
				front["description"] = document.description
			} else {
				delete(front, "description")
			}
			front["page_id"] = stableID("page-", document.identity)
			front["path"] = document.pagePath
			front["source"] = document.source
			front["source_type"] = "docsify"
			front["imported_from"] = provenance
			if err := writeDocsifyFile(stage, document.output, front, document.body); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind, result.Framework, result.Source = "docsify", "docsify", provenance
	result.Pages, result.Sources, result.Truncated = len(documents), totalFiles, truncated
	return result, nil
}

func docsifyGitHubSources(document *htmlNode, shellURL *url.URL) ([]docsifyGitHubSelection, error) {
	queryBasePath := strings.TrimSpace(shellURL.Query().Get("basePath"))
	var scripts strings.Builder
	walkHTML(document, func(node *htmlNode) {
		if node.tag == "script" && node.attrs["src"] == "" {
			script := htmlNodeText(node)
			if !strings.Contains(script, "window.$docsify") {
				return
			}
			scripts.WriteString(script)
			scripts.WriteByte('\n')
		}
	})
	if queryBasePath != "" {
		if len(shellURL.Query()["basePath"]) != 1 {
			return nil, errors.New("Docsify requires exactly one GitHub-backed basePath query")
		}
		if err := docsifyValidateQuerySources(shellURL.Query()); err != nil {
			return nil, err
		}
		selection, err := docsifyGitHubBasePath(queryBasePath)
		if err != nil {
			return nil, fmt.Errorf("Docsify basePath has no complete GitHub inventory: %w", err)
		}
		return []docsifyGitHubSelection{selection}, nil
	}
	code := docsifyJavaScriptWithoutComments(scripts.String())
	if len(docsifyConfigAssign.FindAllStringIndex(code, -1)) != 1 || !docsifyObjectConfig.MatchString(code) || docsifyConfigMutation.MatchString(code) || docsifyConfigSpread.MatchString(code) {
		return nil, errors.New("dynamic Docsify configuration has no complete static inventory")
	}
	var selections []docsifyGitHubSelection
	basePathTokens := docsifyBasePathToken.FindAllStringIndex(code, -1)
	staticBasePaths := docsifyStaticBasePath.FindAllStringSubmatch(code, -1)
	if len(basePathTokens) > 0 {
		if len(basePathTokens) != 1 || len(staticBasePaths) != 1 || !docsifyBasePath.MatchString(code) {
			return nil, errors.New("dynamic Docsify basePath requires an explicit GitHub-backed basePath query")
		}
		selection, err := docsifyGitHubBasePath(staticBasePaths[0][1])
		if err != nil {
			return nil, fmt.Errorf("Docsify basePath has no complete GitHub inventory: %w", err)
		}
		selections = append(selections, selection)
	}

	for _, match := range docsifyGitHubBrowseURL.FindAllStringSubmatch(code, -1) {
		selections = append(selections, docsifyURLSelection(match[1], match[2], match[3], match[4]))
	}
	for _, match := range docsifyGitHubRawURL.FindAllStringSubmatch(code, -1) {
		selections = append(selections, docsifyURLSelection(match[1], match[2], match[3], match[4]))
	}
	for _, match := range docsifyJSDelivrURL.FindAllStringSubmatch(code, -1) {
		mappedPath := match[4]
		if strings.Contains(mappedPath, "$1") {
			mappedPath = strings.TrimSuffix(strings.Split(mappedPath, "$1")[0], "/")
		}
		selections = append(selections, docsifyURLSelection(match[1], match[2], match[3], mappedPath))
	}
	selections = uniqueDocsifySelections(selections)
	if len(selections) > docsifyMaxSelections {
		return nil, fmt.Errorf("Docsify shell advertises more than %d GitHub source roots", docsifyMaxSelections)
	}
	rootFound := false
	for _, selection := range selections {
		rootFound = rootFound || !selection.exact
	}
	if len(selections) == 0 || !rootFound {
		return nil, errors.New("Docsify shell has no uniquely enumerable GitHub source root")
	}
	return selections, nil
}

func docsifyValidateQuerySources(query url.Values) error {
	for _, key := range []string{"homepage", "loadSidebar", "loadNavbar", "coverpage"} {
		for _, value := range query[key] {
			normalized := strings.ReplaceAll(value, "\\", "/")
			parsed, err := url.Parse(normalized)
			if err != nil || parsed.IsAbs() || parsed.Host != "" || strings.HasPrefix(normalized, "/") {
				return fmt.Errorf("Docsify %s is outside the enumerable basePath", key)
			}
			for _, component := range strings.Split(normalized, "/") {
				if component == ".." {
					return fmt.Errorf("Docsify %s escapes the enumerable basePath", key)
				}
			}
		}
	}
	return nil
}

func docsifyJavaScriptWithoutComments(value string) string {
	var output strings.Builder
	var quote rune
	escaped, lineComment, blockComment := false, false, false
	runes := []rune(value)
	for index, character := range runes {
		next := rune(0)
		if index+1 < len(runes) {
			next = runes[index+1]
		}
		switch {
		case lineComment:
			if character == '\n' {
				lineComment = false
				output.WriteRune(character)
			}
		case blockComment:
			if character == '*' && next == '/' {
				blockComment = false
			}
		case quote != 0:
			output.WriteRune(character)
			if escaped {
				escaped = false
			} else if character == '\\' {
				escaped = true
			} else if character == quote {
				quote = 0
			}
		case character == '/' && next == '/':
			lineComment = true
		case character == '/' && next == '*':
			blockComment = true
		case character == '\'' || character == '"' || character == '`':
			quote = character
			output.WriteRune(character)
		default:
			output.WriteRune(character)
		}
	}
	return output.String()
}

func docsifyGitHubBasePath(value string) (docsifyGitHubSelection, error) {
	parsed, err := url.Parse(value)
	if err != nil || parsed.RawQuery != "" || parsed.Fragment != "" {
		return docsifyGitHubSelection{}, errors.New("invalid basePath URL")
	}
	segments := strings.Split(strings.Trim(parsed.Path, "/"), "/")
	switch strings.ToLower(parsed.Hostname()) {
	case "raw.githubusercontent.com":
		if len(segments) < 3 {
			return docsifyGitHubSelection{}, errors.New("raw GitHub basePath requires owner, repository, and ref")
		}
		return docsifySelection(segments[0], segments[1], segments[2], strings.Join(segments[3:], "/")), nil
	case "github.com":
		if len(segments) < 4 || segments[2] != "tree" && segments[2] != "blob" {
			return docsifyGitHubSelection{}, errors.New("GitHub basePath requires a tree or blob URL")
		}
		return docsifySelection(segments[0], segments[1], segments[3], strings.Join(segments[4:], "/")), nil
	default:
		return docsifyGitHubSelection{}, errors.New("basePath host is not GitHub")
	}
}

func docsifySelection(owner, repo, ref, selectedPath string) docsifyGitHubSelection {
	selectedPath = strings.Trim(strings.TrimSpace(selectedPath), "/")
	extension := strings.ToLower(path.Ext(selectedPath))
	return docsifyGitHubSelection{
		owner: owner, repo: repo, ref: ref, path: selectedPath,
		exact: extension == ".md",
	}
}

func docsifyURLSelection(owner, repo, ref, selectedPath string) docsifyGitHubSelection {
	if decoded, err := url.PathUnescape(ref); err == nil {
		ref = decoded
	}
	if decoded, err := url.PathUnescape(selectedPath); err == nil {
		selectedPath = decoded
	}
	return docsifySelection(owner, repo, ref, selectedPath)
}

func uniqueDocsifySelections(values []docsifyGitHubSelection) []docsifyGitHubSelection {
	seen := make(map[string]bool)
	var unique []docsifyGitHubSelection
	for _, value := range values {
		key := value.owner + "\x00" + value.repo + "\x00" + value.ref + "\x00" + value.path + fmt.Sprint(value.exact)
		if !seen[key] {
			seen[key] = true
			unique = append(unique, value)
		}
	}
	sort.Slice(unique, func(i, j int) bool {
		left := unique[i].owner + "/" + unique[i].repo + "/" + unique[i].ref + "/" + unique[i].path
		right := unique[j].owner + "/" + unique[j].repo + "/" + unique[j].ref + "/" + unique[j].path
		return left < right
	})
	return unique
}

func docsifyGitHubInventory(ctx context.Context, selections []docsifyGitHubSelection, reader *sourceReader) ([]docsifyGitHubFile, error) {
	var snapshots map[docsifyRepositoryKey]docsifyGitHubSnapshot
	var err error
	selections, snapshots, err = docsifyResolveSelectionRefs(ctx, selections, reader)
	if err != nil {
		return nil, err
	}
	grouped := make(map[docsifyRepositoryKey][]docsifyGitHubSelection)
	for _, selection := range selections {
		key := docsifyRepositoryKey{selection.owner, selection.repo, selection.ref}
		grouped[key] = append(grouped[key], selection)
	}
	var files []docsifyGitHubFile
	seenFiles := make(map[string]bool)
	for _, key := range sortedRepositoryKeys(grouped) {
		snapshot := snapshots[key]
		treeURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", key.owner, key.repo, snapshot.tree)
		rawTree, _, err := reader.readFromOrigin(ctx, treeURL, nil, "https://api.github.com")
		if err != nil {
			return nil, fmt.Errorf("fetch GitHub tree %s/%s@%s: %w", key.owner, key.repo, snapshot.commit, err)
		}
		var tree struct {
			SHA       string `json:"sha"`
			Truncated *bool  `json:"truncated"`
			Entries   []struct {
				Path string `json:"path"`
				Type string `json:"type"`
				Size int64  `json:"size"`
			} `json:"tree"`
		}
		if err := json.Unmarshal(rawTree, &tree); err != nil || tree.SHA != snapshot.tree || tree.Truncated == nil || *tree.Truncated {
			return nil, fmt.Errorf("GitHub tree %s/%s@%s is malformed or truncated", key.owner, key.repo, snapshot.commit)
		}
		matched := make([]bool, len(grouped[key]))
		for _, entry := range tree.Entries {
			if entry.Type != "blob" || !isMarkdownPath(entry.Path) {
				continue
			}
			if !safeRepositoryPath(entry.Path) {
				for _, selection := range grouped[key] {
					if selection.exact && entry.Path == selection.path || !selection.exact && withinRepositoryPath(entry.Path, selection.path) {
						return nil, fmt.Errorf("GitHub selection %s/%s@%s contains non-portable Markdown path %q", key.owner, key.repo, key.ref, entry.Path)
					}
				}
				continue
			}
			selected := false
			for index, selection := range grouped[key] {
				if selection.exact && entry.Path == selection.path || !selection.exact && withinRepositoryPath(entry.Path, selection.path) {
					matched[index] = true
					selected = true
				}
			}
			identity := strings.ToLower(key.owner) + "/" + strings.ToLower(key.repo) + "/" + snapshot.commit + "/" + entry.Path
			if selected && !seenFiles[identity] {
				seenFiles[identity] = true
				files = append(files, docsifyGitHubFile{owner: key.owner, repo: key.repo, ref: key.ref, commit: snapshot.commit, path: entry.Path, size: entry.Size})
			}
		}
		for index, ok := range matched {
			if !ok {
				return nil, fmt.Errorf("GitHub selection %s/%s@%s/%s contains no Markdown", key.owner, key.repo, key.ref, grouped[key][index].path)
			}
		}
	}
	if len(files) > defaultMaxFiles {
		return nil, fmt.Errorf("Docsify inventory exceeds %d Markdown files", defaultMaxFiles)
	}
	sort.Slice(files, func(i, j int) bool {
		left := files[i].owner + "/" + files[i].repo + "/" + files[i].ref + "/" + files[i].path
		right := files[j].owner + "/" + files[j].repo + "/" + files[j].ref + "/" + files[j].path
		return left < right
	})
	return files, nil
}

func docsifyResolveSelectionRefs(ctx context.Context, selections []docsifyGitHubSelection, reader *sourceReader) ([]docsifyGitHubSelection, map[docsifyRepositoryKey]docsifyGitHubSnapshot, error) {
	snapshots := make(map[docsifyRepositoryKey]docsifyGitHubSnapshot)
	invalid := make(map[docsifyRepositoryKey]bool)
	resolved := make([]docsifyGitHubSelection, 0, len(selections))
	for _, selection := range selections {
		if !validGitHubComponent(selection.owner) || !validGitHubComponent(selection.repo) || !validGitHubRef(selection.ref) || !safeRepositoryPath(selection.path) {
			return nil, nil, fmt.Errorf("invalid GitHub source selection %q", selection.owner+"/"+selection.repo+"@"+selection.ref+"/"+selection.path)
		}
		parts := []string(nil)
		if selection.path != "" {
			parts = strings.Split(selection.path, "/")
		}
		maxConsumed := len(parts)
		if selection.exact && maxConsumed > 0 {
			maxConsumed--
		}
		if maxConsumed+1 > docsifyMaxRefCandidates {
			return nil, nil, fmt.Errorf("GitHub source %s/%s@%s requires more than %d ref candidates", selection.owner, selection.repo, selection.ref, docsifyMaxRefCandidates)
		}
		matched := false
		for consumed := maxConsumed; consumed >= 0; consumed-- {
			candidate := selection.ref
			if consumed > 0 {
				candidate += "/" + strings.Join(parts[:consumed], "/")
			}
			key := docsifyRepositoryKey{selection.owner, selection.repo, candidate}
			snapshot, exists := snapshots[key]
			if !exists && !invalid[key] {
				commitURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s", selection.owner, selection.repo, url.PathEscape(candidate))
				rawCommit, _, readErr := reader.readFromOrigin(ctx, commitURL, nil, "https://api.github.com")
				if readErr != nil {
					if strings.Contains(readErr.Error(), "HTTP 404") || strings.Contains(readErr.Error(), "HTTP 422") {
						invalid[key] = true
						continue
					}
					return nil, nil, fmt.Errorf("resolve GitHub ref %s/%s@%s: %w", selection.owner, selection.repo, candidate, readErr)
				}
				var commit struct {
					SHA    string `json:"sha"`
					Commit struct {
						Tree struct {
							SHA string `json:"sha"`
						} `json:"tree"`
					} `json:"commit"`
				}
				if jsonErr := json.Unmarshal(rawCommit, &commit); jsonErr != nil || !docsifyObjectID.MatchString(commit.SHA) || !docsifyObjectID.MatchString(commit.Commit.Tree.SHA) {
					return nil, nil, fmt.Errorf("GitHub ref %s/%s@%s has no immutable tree", selection.owner, selection.repo, candidate)
				}
				snapshot = docsifyGitHubSnapshot{commit: commit.SHA, tree: commit.Commit.Tree.SHA}
				snapshots[key] = snapshot
			}
			if invalid[key] {
				continue
			}
			selection.ref = candidate
			selection.path = strings.Join(parts[consumed:], "/")
			selection.exact = strings.EqualFold(path.Ext(selection.path), ".md")
			resolved = append(resolved, selection)
			matched = true
			break
		}
		if !matched {
			return nil, nil, fmt.Errorf("GitHub source %s/%s@%s has no resolvable ref", selection.owner, selection.repo, selection.ref)
		}
	}
	return uniqueDocsifySelections(resolved), snapshots, nil
}

func sortedRepositoryKeys[V any](values map[docsifyRepositoryKey]V) []docsifyRepositoryKey {
	keys := make([]docsifyRepositoryKey, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].owner+"/"+keys[i].repo+"/"+keys[i].ref < keys[j].owner+"/"+keys[j].repo+"/"+keys[j].ref
	})
	return keys
}

func validGitHubComponent(value string) bool {
	if value == "" {
		return false
	}
	for _, character := range value {
		if !(character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' || character >= '0' && character <= '9' || strings.ContainsRune("._-", character)) {
			return false
		}
	}
	return true
}

func validGitHubRef(value string) bool {
	if !safeRepositoryPath(value) || value == "@" || strings.HasPrefix(value, ".") || strings.HasSuffix(value, ".") || strings.Contains(value, "..") || strings.Contains(value, "@{") {
		return false
	}
	for _, component := range strings.Split(value, "/") {
		if component == "" || strings.HasSuffix(strings.ToLower(component), ".lock") {
			return false
		}
		for _, character := range component {
			if unicode.IsControl(character) || unicode.IsSpace(character) || strings.ContainsRune(`~^:?*[\`, character) {
				return false
			}
		}
	}
	return true
}

func safeRepositoryPath(value string) bool {
	if value == "" {
		return true
	}
	if strings.Contains(value, "\\") || strings.HasPrefix(value, "/") || path.Clean(value) != value || value == ".." || strings.HasPrefix(value, "../") {
		return false
	}
	for _, component := range strings.Split(value, "/") {
		if component == "" || strings.TrimRight(component, " .") != component {
			return false
		}
		for _, character := range component {
			if unicode.IsControl(character) || strings.ContainsRune(`<>:"|?*`, character) {
				return false
			}
		}
		base := strings.ToUpper(strings.SplitN(component, ".", 2)[0])
		if base == "CON" || base == "PRN" || base == "AUX" || base == "NUL" || len(base) == 4 && (strings.HasPrefix(base, "COM") || strings.HasPrefix(base, "LPT")) && base[3] >= '1' && base[3] <= '9' {
			return false
		}
	}
	return true
}

func withinRepositoryPath(value, root string) bool {
	return root == "" || value == root || strings.HasPrefix(value, root+"/")
}

func isMarkdownPath(value string) bool {
	return strings.EqualFold(path.Ext(value), ".md")
}

func docsifyRawGitHubURL(file docsifyGitHubFile) string {
	segments := strings.Split(file.path, "/")
	for index := range segments {
		segments[index] = url.PathEscape(segments[index])
	}
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", file.owner, file.repo, file.commit, strings.Join(segments, "/"))
}

func docsifyTitle(value string) string {
	name := strings.TrimSuffix(path.Base(value), path.Ext(value))
	if strings.EqualFold(name, "readme") {
		if parent := path.Base(path.Dir(value)); parent != "." && parent != "/" {
			name = parent
		} else {
			name = "Overview"
		}
	}
	words := strings.FieldsFunc(name, func(character rune) bool { return character == '-' || character == '_' || unicode.IsSpace(character) })
	for index, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			words[index] = string(runes)
		}
	}
	if len(words) == 0 {
		return "Documentation"
	}
	return strings.Join(words, " ")
}

func docsifyHeading(body string) string {
	heading := firstHeading(body)
	if badge := strings.Index(heading, " [!["); badge >= 0 {
		heading = heading[:badge]
	}
	if directive := strings.Index(heading, "<!--"); directive >= 0 {
		heading = heading[:directive]
	}
	return strings.TrimSpace(heading)
}

func validateDocsifyOutputPaths(documents []docsifyDocument) error {
	seen := make(map[string]string)
	for _, document := range documents {
		key := strings.ToLower(filepath.ToSlash(filepath.Clean(document.output)))
		if previous, exists := seen[key]; exists {
			return fmt.Errorf("Docsify sources %s and %s collide at output path %s", previous, document.source, document.output)
		}
		seen[key] = document.source
	}
	return nil
}

func writeDocsifyFile(root, relative string, front map[string]any, body string) error {
	raw, err := yaml.Marshal(front)
	if err != nil {
		return err
	}
	content := append([]byte("---\n"), raw...)
	content = append(content, []byte("---\n\n")...)
	content = append(content, strings.ReplaceAll(body, "\r\n", "\n")...)
	if len(content) == 0 || content[len(content)-1] != '\n' {
		content = append(content, '\n')
	}
	return writeFile(root, relative, content)
}
