package library

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

type Source struct {
	Name     string
	FS       fs.FS
	Official bool
}

type catalog struct {
	fingerprint string
	families    []familyRecord
	documents   []documentRecord
}

type familyRecord struct {
	ID          string
	Name        string
	Description string
	Collections []string
}

type documentRecord struct {
	DocID       string
	FamilyID    string
	Name        string
	Version     string
	Description string
	Collections []string
	SourceRoot  string
	Location    string
	Pages       []pageRecord
}

type pageRecord struct {
	PageID       string
	Title        string
	Description  string
	Path         string
	Source       string
	HTTPMethods  []string
	APIEndpoints []string
	OperationIDs []string
	Body         string
	RelativeFile string
	Location     string
	explicitID   bool
}

type manifestFront struct {
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Collections []string `yaml:"collections"`
	SourceRoot  string   `yaml:"source_root"`
}

type pageFront struct {
	Title        string   `yaml:"title"`
	PageID       string   `yaml:"page_id"`
	Path         *string  `yaml:"path"`
	Description  string   `yaml:"description"`
	Source       string   `yaml:"source"`
	HTTPMethods  []string `yaml:"http_methods"`
	APIEndpoints []string `yaml:"api_endpoints"`
	OperationIDs []string `yaml:"operation_ids"`
}

type sourceFile struct {
	path string
	raw  []byte
}

func loadCatalog(sources []Source) (result *catalog, returnErr error) {
	defer func() {
		returnErr = errors.Join(returnErr, closeSources(sources))
	}()
	type sourceDocuments struct {
		official  bool
		documents []documentRecord
	}
	var loaded []sourceDocuments
	userFamilies := make(map[string]bool)
	hash := sha256.New()

	for _, source := range sources {
		if source.FS == nil {
			continue
		}
		files, err := markdownFiles(source)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			fmt.Fprintf(hash, "%s\x00%s\x00", source.Name, file.path)
			hash.Write(file.raw)
			hash.Write([]byte{0})
		}
		docs, err := loadSource(source.Name, files)
		if err != nil {
			return nil, err
		}
		locations := make(map[string][]string)
		for _, doc := range docs {
			locations[doc.DocID] = append(locations[doc.DocID], doc.Location)
			if !source.Official {
				userFamilies[doc.FamilyID] = true
			}
		}
		if err := validateUniqueDocuments(locations); err != nil {
			return nil, err
		}
		loaded = append(loaded, sourceDocuments{official: source.Official, documents: docs})
	}

	var documents []documentRecord
	locations := make(map[string][]string)
	for _, source := range loaded {
		for _, doc := range source.documents {
			// A user-defined family replaces its official counterpart as one unit.
			if source.official && userFamilies[doc.FamilyID] {
				continue
			}
			locations[doc.DocID] = append(locations[doc.DocID], doc.Location)
			documents = append(documents, doc)
		}
	}
	if err := validateUniqueDocuments(locations); err != nil {
		return nil, err
	}

	sort.Slice(documents, func(i, j int) bool { return documents[i].DocID < documents[j].DocID })
	families, err := deriveFamilies(documents)
	if err != nil {
		return nil, err
	}
	return &catalog{
		fingerprint: hex.EncodeToString(hash.Sum(nil)),
		families:    families,
		documents:   documents,
	}, nil
}

func closeSources(sources []Source) error {
	var result error
	for _, source := range sources {
		if closer, ok := source.FS.(io.Closer); ok {
			result = errors.Join(result, closer.Close())
		}
	}
	return result
}

func validateUniqueDocuments(locations map[string][]string) error {
	for docID, found := range locations {
		if len(found) < 2 {
			continue
		}
		sort.Strings(found)
		return &ValidationError{
			Message:   fmt.Sprintf("duplicate doc_id %q", docID),
			Locations: found,
		}
	}
	return nil
}

func markdownFiles(source Source) ([]sourceFile, error) {
	var files []sourceFile
	err := fs.WalkDir(source.FS, ".", func(name string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("walk %s at %s: %w", source.Name, name, walkErr)
		}
		if entry.IsDir() || !strings.EqualFold(path.Ext(name), ".md") {
			return nil
		}
		raw, err := fs.ReadFile(source.FS, name)
		if err != nil {
			return fmt.Errorf("read %s:%s: %w", source.Name, name, err)
		}
		files = append(files, sourceFile{path: name, raw: raw})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool { return files[i].path < files[j].path })
	return files, nil
}

func loadSource(sourceName string, files []sourceFile) ([]documentRecord, error) {
	manifestByRoot := make(map[string]sourceFile)
	var roots []string
	for _, file := range files {
		if path.Base(file.path) != "_index.md" {
			continue
		}
		root := path.Dir(file.path)
		manifestByRoot[root] = file
		roots = append(roots, root)
	}
	sort.Strings(roots)

	documents := make([]documentRecord, 0, len(roots))
	for _, root := range roots {
		manifest := manifestByRoot[root]
		location := sourceLocation(sourceName, manifest.path)
		var front manifestFront
		if _, err := parseFrontmatter(manifest.raw, &front); err != nil {
			return nil, &ValidationError{Message: fmt.Sprintf("%s: invalid manifest: %v", location, err)}
		}
		front.Name = strings.TrimSpace(front.Name)
		front.Version = strings.TrimSpace(front.Version)
		if front.Name == "" || front.Version == "" {
			return nil, &ValidationError{Message: fmt.Sprintf("%s: manifest requires name and version", location)}
		}
		familyID := slug(front.Name)
		versionID := slug(front.Version)
		if familyID == "" || versionID == "" {
			return nil, &ValidationError{Message: fmt.Sprintf("%s: name and version must contain letters or digits", location)}
		}
		collections, err := cleanCollections(front.Collections)
		if err != nil {
			return nil, &ValidationError{Message: fmt.Sprintf("%s: %v", location, err)}
		}
		doc := documentRecord{
			DocID:       familyID + "-" + versionID,
			FamilyID:    familyID,
			Name:        front.Name,
			Version:     front.Version,
			Description: strings.TrimSpace(front.Description),
			Collections: collections,
			SourceRoot:  strings.TrimSpace(front.SourceRoot),
			Location:    location,
		}

		for _, file := range files {
			if file.path == manifest.path || ownerRoot(file.path, roots) != root {
				continue
			}
			rel, ok := underRoot(file.path, root)
			if !ok {
				continue
			}
			page, err := loadPage(sourceName, rel, file)
			if err != nil {
				return nil, err
			}
			doc.Pages = append(doc.Pages, page)
		}
		if err := assignPageIDs(&doc); err != nil {
			return nil, err
		}
		sort.Slice(doc.Pages, func(i, j int) bool { return doc.Pages[i].PageID < doc.Pages[j].PageID })
		documents = append(documents, doc)
	}
	return documents, nil
}

func loadPage(sourceName, relative string, file sourceFile) (pageRecord, error) {
	location := sourceLocation(sourceName, file.path)
	var front pageFront
	body, err := parseFrontmatter(file.raw, &front)
	if err != nil {
		return pageRecord{}, &ValidationError{Message: fmt.Sprintf("%s: invalid page: %v", location, err)}
	}
	front.Title = strings.TrimSpace(front.Title)
	if front.Title == "" {
		return pageRecord{}, &ValidationError{Message: fmt.Sprintf("%s: page frontmatter requires title", location)}
	}
	pagePath := path.Dir(relative)
	if pagePath == "." {
		pagePath = ""
	}
	if front.Path != nil {
		pagePath, err = canonicalPath(*front.Path)
		if err != nil {
			return pageRecord{}, &ValidationError{Message: fmt.Sprintf("%s: invalid path: %v", location, err)}
		}
	}
	pageID := strings.TrimSpace(front.PageID)
	explicitID := pageID != ""
	if !explicitID {
		pageID = slug(front.Title)
	}
	if !validSlug(pageID) {
		return pageRecord{}, &ValidationError{Message: fmt.Sprintf("%s: invalid page_id %q", location, pageID)}
	}
	return pageRecord{
		PageID:       pageID,
		Title:        front.Title,
		Description:  strings.TrimSpace(front.Description),
		Path:         pagePath,
		Source:       strings.TrimSpace(front.Source),
		HTTPMethods:  cleanStrings(front.HTTPMethods),
		APIEndpoints: cleanStrings(front.APIEndpoints),
		OperationIDs: cleanStrings(front.OperationIDs),
		Body:         body,
		RelativeFile: relative,
		Location:     location,
		explicitID:   explicitID,
	}, nil
}

func assignPageIDs(doc *documentRecord) error {
	groups := make(map[string][]int)
	for i := range doc.Pages {
		groups[doc.Pages[i].PageID] = append(groups[doc.Pages[i].PageID], i)
	}
	for base, indexes := range groups {
		if len(indexes) < 2 {
			continue
		}
		var explicit []string
		for _, index := range indexes {
			if doc.Pages[index].explicitID {
				explicit = append(explicit, doc.Pages[index].Location)
			}
		}
		if len(explicit) > 1 {
			sort.Strings(explicit)
			return &ValidationError{
				Message:   fmt.Sprintf("duplicate explicit page_id %q in %s", base, doc.DocID),
				Locations: explicit,
			}
		}
		for _, index := range indexes {
			if doc.Pages[index].explicitID {
				continue
			}
			sum := sha256.Sum256([]byte(doc.Pages[index].RelativeFile))
			doc.Pages[index].PageID = base + "-" + hex.EncodeToString(sum[:4])
		}
	}
	seen := make(map[string]string)
	for _, page := range doc.Pages {
		if previous, ok := seen[page.PageID]; ok {
			return &ValidationError{
				Message:   fmt.Sprintf("duplicate page_id %q in %s", page.PageID, doc.DocID),
				Locations: []string{previous, page.Location},
			}
		}
		seen[page.PageID] = page.Location
	}
	return nil
}

func deriveFamilies(documents []documentRecord) ([]familyRecord, error) {
	byID := make(map[string]familyRecord)
	locations := make(map[string]string)
	for _, doc := range documents {
		family, ok := byID[doc.FamilyID]
		if !ok {
			byID[doc.FamilyID] = familyRecord{
				ID: doc.FamilyID, Name: doc.Name, Description: doc.Description,
				Collections: append([]string(nil), doc.Collections...),
			}
			locations[doc.FamilyID] = doc.Location
			continue
		}
		if family.Name != doc.Name || family.Description != doc.Description || !equalStrings(family.Collections, doc.Collections) {
			return nil, &ValidationError{
				Message:   fmt.Sprintf("versions of API family %q disagree on family metadata", doc.Name),
				Locations: []string{locations[doc.FamilyID], doc.Location},
			}
		}
	}
	families := make([]familyRecord, 0, len(byID))
	for _, family := range byID {
		families = append(families, family)
	}
	sort.Slice(families, func(i, j int) bool {
		left, right := strings.ToLower(families[i].Name), strings.ToLower(families[j].Name)
		if left == right {
			return families[i].ID < families[j].ID
		}
		return left < right
	})
	return families, nil
}

func parseFrontmatter(raw []byte, target any) (string, error) {
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return "", fmt.Errorf("YAML frontmatter must start on the first line")
	}
	closing := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			closing = i
			break
		}
	}
	if closing < 0 {
		return "", fmt.Errorf("YAML frontmatter is missing its closing delimiter")
	}
	if err := yaml.Unmarshal([]byte(strings.Join(lines[1:closing], "\n")), target); err != nil {
		return "", err
	}
	bodyLines := lines[closing+1:]
	if len(bodyLines) > 0 && bodyLines[0] == "" {
		bodyLines = bodyLines[1:]
	}
	return strings.Join(bodyLines, "\n"), nil
}

func ownerRoot(file string, roots []string) string {
	owner := ""
	ownerLength := -1
	for _, root := range roots {
		if _, ok := underRoot(file, root); ok && len(root) > ownerLength {
			owner = root
			ownerLength = len(root)
		}
	}
	return owner
}

func underRoot(file, root string) (string, bool) {
	if root == "." {
		return file, true
	}
	prefix := root + "/"
	if !strings.HasPrefix(file, prefix) {
		return "", false
	}
	return strings.TrimPrefix(file, prefix), true
}

func sourceLocation(sourceName, name string) string {
	if sourceName == "" {
		return name
	}
	if filepath.IsAbs(sourceName) {
		return filepath.Join(sourceName, filepath.FromSlash(name))
	}
	return sourceName + ":" + name
}

func canonicalPath(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", nil
	}
	if strings.Contains(value, "\\") || strings.HasPrefix(value, "/") || strings.HasSuffix(value, "/") {
		return "", fmt.Errorf("must be a relative slash-separated path without leading or trailing slashes")
	}
	clean := path.Clean(value)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") || clean != value || strings.Contains(value, "//") {
		return "", fmt.Errorf("must be canonical and may not contain dot segments")
	}
	for _, segment := range strings.Split(value, "/") {
		if segment == "" || segment == "." || segment == ".." {
			return "", fmt.Errorf("contains an invalid path segment")
		}
	}
	return value, nil
}

func slug(value string) string {
	var out strings.Builder
	dash := false
	for _, r := range strings.ToLower(strings.TrimSpace(value)) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			if dash && out.Len() > 0 {
				out.WriteByte('-')
			}
			out.WriteRune(r)
			dash = false
			continue
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			// IDs intentionally remain portable ASCII; non-ASCII separates words.
			dash = true
			continue
		}
		dash = true
	}
	return strings.Trim(out.String(), "-")
}

func validSlug(value string) bool {
	return value != "" && slug(value) == value
}

func cleanCollections(values []string) ([]string, error) {
	values = cleanStrings(values)
	for _, value := range values {
		if !validCollection(value) {
			return nil, fmt.Errorf("invalid collection %q", value)
		}
	}
	return values, nil
}

func validCollection(value string) bool {
	if value == "" || strings.HasPrefix(value, "_") || strings.HasSuffix(value, "_") {
		return false
	}
	previousUnderscore := false
	for _, r := range value {
		if r == '_' {
			if previousUnderscore {
				return false
			}
			previousUnderscore = true
			continue
		}
		if r < 'a' || r > 'z' {
			if r < '0' || r > '9' {
				return false
			}
		}
		previousUnderscore = false
	}
	return true
}

func cleanStrings(values []string) []string {
	seen := make(map[string]struct{})
	out := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}

func equalStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
