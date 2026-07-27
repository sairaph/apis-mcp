package importer

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	xhtml "golang.org/x/net/html"
)

const (
	DefaultMaxHTMLPages = 50
	DefaultMaxHTMLDepth = 3
	maxHTMLTreeDepth    = 128
	maxHTMLNodes        = 50_000
	maxHTMLSitemapURLs  = 20_000
)

var (
	htmlMetaTag             = regexp.MustCompile(`(?is)<meta\b[^>]*>`)
	mkdocsMaterialGenerator = regexp.MustCompile(`(?:^|,\s*)mkdocs-material(?:-|$)`)
	mkdocsGenerator         = regexp.MustCompile(`(?:^|,\s*)mkdocs(?:-[0-9]+(?:\.[0-9]+)*)?(?:\s*(?:,|$))`)
	mkdocsScopeURL          = regexp.MustCompile(`__md_scope\s*=\s*new\s+URL\(\s*["']([^"']*)["']`)
	mkdocsBaseURL           = regexp.MustCompile(`\bbase_url\s*=\s*["']([^"']*)["']`)
	mkdocsVersionComment    = regexp.MustCompile(`(?im)^\s*MkDocs\s+version\s*:\s*[0-9]+(?:\.[0-9]+)+(?:[A-Za-z0-9.+-]*)?\s*$`)
	mkdocsBuildComment      = regexp.MustCompile(`(?im)^\s*Docs Build Date UTC\s*:`)
)

type htmlNode struct {
	tag          string
	attrs        map[string]string
	text         string
	children     []*htmlNode
	parent       *htmlNode
	htmlDocument bool
}

type crawledPage struct {
	title       string
	description string
	source      string
	markdown    string
}

type htmlDocumentView struct {
	framework     string
	content       *htmlNode
	linkRoots     []*htmlNode
	includeHeader bool
	profileError  error
}

// ImportHTML crawls a bounded set of static, same-origin HTML pages and
// publishes their useful document content as canonical Markdown.
func ImportHTML(ctx context.Context, name, version, source string, options Options) (Result, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return Result{}, err
	}
	if strings.TrimSpace(name) == "" || strings.TrimSpace(version) == "" {
		return Result{}, errors.New("HTML import requires API name and version")
	}
	start, err := url.Parse(strings.TrimSpace(source))
	if err != nil || start.Host == "" || start.Scheme != "http" && start.Scheme != "https" {
		return Result{}, errors.New("HTML import source must be an HTTP(S) URL")
	}
	start.Fragment = ""
	scopePath := "/"
	if options.HTMLScope == "path" {
		scopePath = documentationPathScope(start.Path)
	}
	var finiteInventory map[string]bool

	// A same-origin link must not escape through an HTTP redirect either.
	client := *options.HTTPClient
	previousRedirectPolicy := client.CheckRedirect
	client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if !sameOrigin(start, request.URL) || options.HTMLScope == "path" && !withinDocumentationPath(request.URL.Path, scopePath) {
			return errors.New("HTML crawl redirect crosses the source origin")
		}
		if finiteInventory != nil {
			canonical, err := canonicalHTTPURL(request.URL.String())
			if err != nil || !finiteInventory[canonical] {
				return errors.New("HTML crawl redirect crosses the finite inventory")
			}
		}
		if previousRedirectPolicy != nil {
			return previousRedirectPolicy(request, via)
		}
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return nil
	}
	options.HTTPClient = &client
	reader := newSourceReader(options)

	type pendingPage struct {
		source string
		depth  int
	}
	queue := []pendingPage{{source: start.String()}}
	seen := map[string]bool{start.String(): true}
	processed := make(map[string]bool)
	pageCapacity := options.MaxHTMLPages
	if pageCapacity < 0 {
		pageCapacity = DefaultMaxHTMLPages
	}
	pages := make([]crawledPage, 0, pageCapacity)
	rootSource := start.String()
	rootFramework := ""
	fetchedInventory := make(map[string]bool)
	generatedInventory := make(map[string]bool)
	inventoryAliases := make(map[string]string)
	crawlRequests := 0
	crawlLimited := false
	var crawlErrors []error
	for len(queue) > 0 && (options.MaxHTMLPages < 0 || crawlRequests < options.MaxHTMLPages) {
		if err := ctx.Err(); err != nil {
			return Result{}, err
		}
		pending := queue[0]
		queue = queue[1:]
		crawlRequests++
		raw, provenance, readErr := reader.read(ctx, pending.source, nil)
		if readErr != nil {
			if len(pages) == 0 {
				return Result{}, readErr
			}
			if rootFramework != "" && rootFramework != "unknown" {
				crawlErrors = append(crawlErrors, fmt.Errorf("fetch framework page %s: %w", pending.source, readErr))
			}
			continue
		}
		pageURL, parseErr := url.Parse(provenance)
		if parseErr != nil || !sameOrigin(start, pageURL) {
			continue
		}
		if finiteInventory != nil {
			canonicalPending, canonicalErr := canonicalHTTPURL(pending.source)
			if canonicalErr == nil && finiteInventory[canonicalPending] {
				fetchedInventory[canonicalPending] = true
			}
			canonicalProvenance, canonicalErr := canonicalHTTPURL(provenance)
			if canonicalErr != nil || !finiteInventory[canonicalProvenance] || !withinDocumentationPath(pageURL.Path, scopePath) {
				return Result{}, fmt.Errorf("%s crawl did not complete: inventory entry %s redirected outside its finite inventory", rootFramework, pending.source)
			}
			if canonicalErr == nil && canonicalPending != canonicalProvenance {
				inventoryAliases[canonicalPending] = canonicalProvenance
			}
		}
		if processed[provenance] {
			continue
		}
		redirected, refreshErr := htmlRefreshRedirect(raw, pageURL)
		if refreshErr != nil {
			return Result{}, fmt.Errorf("%s crawl did not complete: invalid refresh on %s: %w", rootFramework, provenance, refreshErr)
		}
		if redirected != nil {
			canonicalRedirect, canonicalErr := canonicalHTTPURL(redirected.String())
			if finiteInventory == nil || canonicalErr != nil || !finiteInventory[canonicalRedirect] {
				return Result{}, fmt.Errorf("%s crawl did not complete: page %s redirects outside its finite inventory", rootFramework, provenance)
			}
			canonicalAlias, canonicalErr := canonicalHTTPURL(provenance)
			if canonicalErr != nil || canonicalAlias == canonicalRedirect {
				return Result{}, fmt.Errorf("%s crawl did not complete: page %s has a self-referential refresh", rootFramework, provenance)
			}
			inventoryAliases[canonicalAlias] = canonicalRedirect
			seen[provenance] = true
			processed[provenance] = true
			reportProgress(options, Progress{Stage: "redirect", Framework: rootFramework, URL: provenance, Pages: len(pages), Queued: len(queue)})
			continue
		}
		document, parseErr := parseHTML(raw)
		if parseErr != nil {
			return Result{}, fmt.Errorf("parse HTML page %s: %w", provenance, parseErr)
		}
		if !isHTMLDocument(document) {
			if len(pages) == 0 {
				return Result{}, errors.New("source is not a static HTML document")
			}
			if rootFramework != "" && rootFramework != "unknown" {
				crawlErrors = append(crawlErrors, fmt.Errorf("framework page is not HTML: %s", provenance))
			}
			continue
		}
		if len(pages) == 0 {
			rootSource = provenance
		}
		view := inspectHTMLDocument(document)
		if len(pages) == 0 {
			rootFramework = view.framework
			reportProgress(options, Progress{Stage: "detected", Framework: rootFramework, URL: provenance})
		} else if rootFramework != "" && view.framework != rootFramework {
			crawlErrors = append(crawlErrors, fmt.Errorf("framework changed from %s to %s at %s", rootFramework, view.framework, provenance))
			continue
		}
		if view.profileError != nil {
			return Result{}, fmt.Errorf("%s crawl did not complete: %w", rootFramework, view.profileError)
		}
		seen[provenance] = true
		processed[provenance] = true
		if len(pages) == 0 && isMkDocsFramework(rootFramework) {
			siteRoot, inventoryErr := mkdocsSiteRoot(view, pageURL)
			if inventoryErr == nil {
				scopePath = siteRoot
			}
			var inventory []string
			if inventoryErr == nil {
				inventory, inventoryErr = mkdocsSitemapInventory(ctx, pageURL, siteRoot, reader)
			}
			if inventoryErr != nil {
				return Result{}, fmt.Errorf("%s crawl did not complete: %w", rootFramework, inventoryErr)
			}
			finiteInventory = make(map[string]bool, len(inventory))
			for _, linked := range inventory {
				finiteInventory[linked] = true
			}
			canonicalProvenance, canonicalErr := mkdocsInventoryURL(provenance)
			if canonicalErr != nil || !containsString(inventory, canonicalProvenance) {
				return Result{}, fmt.Errorf("%s crawl did not complete: finite inventory does not contain starting page %s", rootFramework, provenance)
			}
			seen[canonicalProvenance] = true
			fetchedInventory[canonicalProvenance] = true
			if options.MaxHTMLDepth == 0 {
				crawlLimited = len(inventory) > 1
			} else {
				for _, linked := range inventory {
					if !seen[linked] {
						seen[linked] = true
						queue = append(queue, pendingPage{source: linked, depth: 1})
					}
				}
			}
		}
		title, markdown := htmlToMarkdown(view.content, pageURL, view.includeHeader)
		if title == "" {
			title = pageTitleFromURL(pageURL)
		}
		if strings.TrimSpace(markdown) == "" {
			markdown = "# " + title
		}
		pages = append(pages, crawledPage{
			title: title, description: htmlMetadataDescription(document),
			source: provenance, markdown: markdown,
		})
		if finiteInventory != nil {
			canonicalGenerated, canonicalErr := mkdocsInventoryURL(provenance)
			if canonicalErr == nil {
				generatedInventory[canonicalGenerated] = true
			}
		}
		reportProgress(options, Progress{Stage: "page", Framework: rootFramework, URL: provenance, Pages: len(pages), Queued: len(queue)})

		var linkedPages []string
		if !isMkDocsFramework(rootFramework) {
			linkedPages = htmlPageLinks(view.linkRoots, pageURL, options.HTMLScope, scopePath)
		}
		if options.MaxHTMLDepth >= 0 && pending.depth >= options.MaxHTMLDepth {
			for _, linked := range linkedPages {
				if !seen[linked] {
					crawlLimited = true
					break
				}
			}
			continue
		}
		for _, linked := range linkedPages {
			if seen[linked] {
				continue
			}
			seen[linked] = true
			queue = append(queue, pendingPage{source: linked, depth: pending.depth + 1})
		}
	}
	if len(crawlErrors) > 0 {
		return Result{}, fmt.Errorf("%s crawl did not complete: %w", rootFramework, errors.Join(crawlErrors...))
	}
	if len(finiteInventory) > 0 && !crawlLimited && len(queue) == 0 && len(fetchedInventory) != len(finiteInventory) {
		return Result{}, fmt.Errorf("%s crawl did not complete: finite inventory advertised %d entries but %d were fetched", rootFramework, len(finiteInventory), len(fetchedInventory))
	}
	if len(finiteInventory) > 0 && !crawlLimited && len(queue) == 0 {
		if err := validateInventoryAliases(inventoryAliases, generatedInventory); err != nil {
			return Result{}, fmt.Errorf("%s crawl did not complete: %w", rootFramework, err)
		}
	}
	if len(pages) == 0 {
		return Result{}, errors.New("HTML crawl produced no pages")
	}

	description := pages[0].description
	truncated := crawlLimited || len(queue) > 0
	reportProgress(options, Progress{Stage: "publishing", Framework: rootFramework, Pages: len(pages), Truncated: truncated})
	result, err := publish(ctx, options, name, version, func(stage string) error {
		metadata := manifest{
			Name: name, Version: version, Description: description, Collections: options.Collections,
			SourceRoot: rootSource, SourceType: "html", ImportedFrom: rootSource,
		}
		if err := writeCanonicalFile(stage, "_index.md", metadata, "This document set was generated from static HTML pages."); err != nil {
			return err
		}
		for _, page := range pages {
			front := pageFront{
				Title: page.title, PageID: stableID("page-", page.source), Path: "documentation",
				Description: page.description, Source: page.source,
				SourceType: "html", ImportedFrom: rootSource,
			}
			filename := stableID("", page.source) + ".md"
			if err := writeCanonicalFile(stage, filepath.Join("documentation", filename), front, page.markdown); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind, result.Framework, result.Source, result.Pages = "html", rootFramework, rootSource, len(pages)
	result.Truncated = truncated
	reportProgress(options, Progress{Stage: "completed", Framework: rootFramework, Pages: len(pages), Truncated: truncated})
	return result, nil
}

func htmlMetadataDescription(root *htmlNode) string {
	var description string
	walkHTML(root, func(node *htmlNode) {
		if description != "" || node.tag != "meta" {
			return
		}
		name := strings.ToLower(strings.TrimSpace(node.attrs["name"]))
		property := strings.ToLower(strings.TrimSpace(node.attrs["property"]))
		if name == "description" || name == "twitter:description" || property == "og:description" || property == "twitter:description" {
			description = strings.TrimSpace(node.attrs["content"])
		}
	})
	return description
}

func parseHTML(raw []byte) (*htmlNode, error) {
	isDocument, err := inspectHTMLSource(raw)
	if err != nil {
		return nil, err
	}
	parsed, err := xhtml.Parse(bytes.NewReader(raw))
	if err != nil {
		if strings.Contains(err.Error(), "open stack of elements exceeds") {
			return nil, fmt.Errorf("HTML document exceeds maximum tree depth of %d: %w", maxHTMLTreeDepth, err)
		}
		return nil, err
	}
	root := &htmlNode{tag: "document", htmlDocument: isDocument}
	nodes := 1
	var convert func(*xhtml.Node, *htmlNode, int) error
	convert = func(source *xhtml.Node, parent *htmlNode, depth int) error {
		if depth > maxHTMLTreeDepth {
			return fmt.Errorf("HTML document exceeds maximum tree depth of %d", maxHTMLTreeDepth)
		}
		for current := source; current != nil; current = current.NextSibling {
			var converted *htmlNode
			switch current.Type {
			case xhtml.ElementNode:
				attributes := make(map[string]string, len(current.Attr))
				for _, attribute := range current.Attr {
					attributes[strings.ToLower(attribute.Key)] = attribute.Val
				}
				converted = &htmlNode{tag: strings.ToLower(current.Data), attrs: attributes, parent: parent}
			case xhtml.TextNode:
				if current.Data == "" {
					continue
				}
				converted = &htmlNode{text: current.Data, parent: parent}
			case xhtml.CommentNode:
				converted = &htmlNode{tag: "#comment", text: current.Data, parent: parent}
			default:
				if err := convert(current.FirstChild, parent, depth); err != nil {
					return err
				}
				continue
			}
			nodes++
			if nodes > maxHTMLNodes {
				return fmt.Errorf("HTML document exceeds maximum node count of %d", maxHTMLNodes)
			}
			parent.children = append(parent.children, converted)
			if err := convert(current.FirstChild, converted, depth+1); err != nil {
				return err
			}
		}
		return nil
	}
	if err := convert(parsed.FirstChild, root, 1); err != nil {
		return nil, err
	}
	return root, nil
}

func inspectHTMLSource(raw []byte) (bool, error) {
	tokenizer := xhtml.NewTokenizer(bytes.NewReader(raw))
	nodes := 1
	isDocument := false
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case xhtml.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return isDocument, nil
			}
			return false, tokenizer.Err()
		case xhtml.DoctypeToken:
			if strings.EqualFold(tokenizer.Token().Data, "html") {
				isDocument = true
			}
		case xhtml.StartTagToken, xhtml.SelfClosingTagToken:
			token := tokenizer.Token()
			switch strings.ToLower(token.Data) {
			case "html", "head", "body", "title", "main", "article":
				isDocument = true
			}
			nodes++
		case xhtml.TextToken:
			if len(tokenizer.Text()) > 0 {
				nodes++
			}
		}
		if nodes > maxHTMLNodes {
			return false, fmt.Errorf("HTML document exceeds maximum node count of %d", maxHTMLNodes)
		}
	}
}

func htmlRefreshRedirect(raw []byte, base *url.URL) (*url.URL, error) {
	tokenizer := xhtml.NewTokenizer(bytes.NewReader(raw))
	inNoScript := false
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case xhtml.ErrorToken:
			return nil, nil
		case xhtml.StartTagToken, xhtml.SelfClosingTagToken:
			token := tokenizer.Token()
			if strings.EqualFold(token.Data, "noscript") {
				inNoScript = tokenType == xhtml.StartTagToken
				continue
			}
			if strings.EqualFold(token.Data, "meta") {
				if target, err := refreshTargetFromMeta(token, base); target != nil || err != nil {
					return target, err
				}
			}
		case xhtml.EndTagToken:
			if strings.EqualFold(tokenizer.Token().Data, "noscript") {
				inNoScript = false
			}
		case xhtml.TextToken:
			if !inNoScript {
				continue
			}
			for _, tag := range htmlMetaTag.FindAll(tokenizer.Text(), -1) {
				nested := xhtml.NewTokenizer(bytes.NewReader(tag))
				if nestedType := nested.Next(); nestedType == xhtml.StartTagToken || nestedType == xhtml.SelfClosingTagToken {
					if target, err := refreshTargetFromMeta(nested.Token(), base); target != nil || err != nil {
						return target, err
					}
				}
			}
		}
	}
}

func refreshTargetFromMeta(token xhtml.Token, base *url.URL) (*url.URL, error) {
	attributes := make(map[string]string, len(token.Attr))
	for _, attribute := range token.Attr {
		attributes[strings.ToLower(attribute.Key)] = strings.TrimSpace(attribute.Val)
	}
	if !strings.EqualFold(attributes["http-equiv"], "refresh") {
		return nil, nil
	}
	content := attributes["content"]
	parts := strings.SplitN(content, ";", 2)
	if len(parts) != 2 || strings.TrimSpace(parts[0]) != "0" {
		return nil, errors.New("refresh meta tag must use a zero-second delay")
	}
	index := strings.Index(strings.ToLower(parts[1]), "url=")
	if index < 0 {
		return nil, errors.New("refresh meta tag has no URL")
	}
	target := strings.Trim(strings.TrimSpace(parts[1][index+4:]), "'\"")
	if target == "" {
		return nil, errors.New("refresh meta tag has an empty URL")
	}
	resolved, err := base.Parse(target)
	if err != nil || !sameOrigin(base, resolved) || !likelyHTMLPath(resolved.Path) {
		return nil, errors.New("refresh target is not a same-origin HTML page")
	}
	resolved.Fragment = ""
	return resolved, nil
}

func isHTMLDocument(root *htmlNode) bool {
	return root != nil && root.htmlDocument
}

func walkHTML(node *htmlNode, visit func(*htmlNode)) {
	visit(node)
	for _, child := range node.children {
		walkHTML(child, visit)
	}
}

func inspectHTMLDocument(root *htmlNode) htmlDocumentView {
	view := htmlDocumentView{content: root, linkRoots: []*htmlNode{root}}
	view.framework = detectHTMLFramework(root)
	switch view.framework {
	case "docusaurus":
		if content := firstHTMLClass(root, "theme-doc-markdown"); content != nil {
			view.content = content
			view.includeHeader = true
		} else if content := firstHTMLClassPrefix(root, "generatedIndexPage_"); content != nil {
			view.content = content
			view.includeHeader = true
		}
		view.linkRoots = append(htmlClasses(root, "theme-doc-sidebar-menu"), htmlClasses(root, "pagination-nav")...)
		if len(view.linkRoots) == 0 {
			view.linkRoots = []*htmlNode{view.content}
		}
	case "mkdocs-material":
		content := firstHTMLClass(root, "md-content__inner")
		if content == nil {
			view.profileError = errors.New("MkDocs Material page has no md-content__inner content container")
			return view
		}
		view.content = content
		view.linkRoots = htmlClasses(root, "md-nav--primary")
		if len(view.linkRoots) == 0 {
			view.profileError = errors.New("MkDocs Material page has no primary navigation")
		}
	case "mkdocs":
		content := firstHTMLAttribute(root, "role", "main")
		if content == nil {
			view.profileError = errors.New("MkDocs page has no main content container")
			return view
		}
		view.content = content
		if navigation := firstHTMLAttribute(root, "id", "navbar-collapse"); navigation != nil {
			view.linkRoots = []*htmlNode{navigation}
		} else {
			view.linkRoots = htmlClasses(root, "wy-menu-vertical")
		}
	}
	return view
}

func detectHTMLFramework(root *htmlNode) string {
	framework := ""
	hasMainContent := firstHTMLAttribute(root, "role", "main") != nil
	walkHTML(root, func(node *htmlNode) {
		if framework != "" {
			return
		}
		if node.tag == "meta" && strings.EqualFold(strings.TrimSpace(node.attrs["name"]), "generator") {
			generator := strings.ToLower(strings.TrimSpace(node.attrs["content"]))
			switch {
			case mkdocsMaterialGenerator.MatchString(generator):
				framework = "mkdocs-material"
				return
			case mkdocsGenerator.MatchString(generator):
				framework = "mkdocs"
				return
			case strings.HasPrefix(generator, "docusaurus"):
				framework = "docusaurus"
				return
			}
		}
		if node.tag == "#comment" && hasMainContent && mkdocsVersionComment.MatchString(node.text) && mkdocsBuildComment.MatchString(node.text) {
			framework = "mkdocs"
			return
		}
		if node.attrs["id"] == "__docusaurus" || hasHTMLClass(node, "docs-doc-page") && hasHTMLClass(node, "plugin-docs") {
			framework = "docusaurus"
		}
	})
	if framework == "" && looksLikeMkDocsBuiltIn(root) {
		framework = "mkdocs"
	}
	return framework
}

func looksLikeMkDocsBuiltIn(root *htmlNode) bool {
	if firstHTMLAttribute(root, "role", "main") == nil {
		return false
	}
	if firstHTMLAttribute(root, "id", "navbar-collapse") != nil && htmlAssetSuffix(root, "link", "href", "css/base.css") && htmlAssetSuffix(root, "script", "src", "js/base.js") {
		return true
	}
	if len(htmlClasses(root, "wy-menu-vertical")) == 0 || !htmlAssetSuffix(root, "script", "src", "js/theme.js") {
		return false
	}
	foundRuntime := false
	walkHTML(root, func(node *htmlNode) {
		if foundRuntime || node.tag != "script" {
			return
		}
		script := htmlNodeText(node)
		foundRuntime = strings.Contains(script, "mkdocs_page_name") || mkdocsBaseURL.MatchString(script)
	})
	return foundRuntime
}

func htmlAssetSuffix(root *htmlNode, tag, attribute, suffix string) bool {
	found := false
	walkHTML(root, func(node *htmlNode) {
		if found || node.tag != tag {
			return
		}
		value := strings.SplitN(strings.TrimSpace(node.attrs[attribute]), "?", 2)[0]
		found = strings.HasSuffix(value, suffix)
	})
	return found
}

func mkdocsSiteRoot(view htmlDocumentView, pageURL *url.URL) (string, error) {
	document := view.content
	for document.parent != nil {
		document = document.parent
	}
	var materialRoot, baseRoot string
	walkHTML(document, func(node *htmlNode) {
		if node.tag != "script" {
			return
		}
		script := htmlNodeText(node)
		match := mkdocsScopeURL.FindStringSubmatch(script)
		if len(match) == 2 {
			resolved, err := pageURL.Parse(match[1])
			if err == nil && sameOrigin(pageURL, resolved) && resolved.RawQuery == "" {
				materialRoot = normalizeDocumentationRoot(resolved.Path)
			}
		}
		match = mkdocsBaseURL.FindStringSubmatch(script)
		if len(match) == 2 {
			resolved, err := pageURL.Parse(match[1])
			if err == nil && sameOrigin(pageURL, resolved) && resolved.RawQuery == "" {
				baseRoot = normalizeDocumentationRoot(resolved.Path)
			}
		}
	})
	if materialRoot != "" {
		return materialRoot, nil
	}
	if baseRoot != "" {
		return baseRoot, nil
	}

	paths := []string{pageURL.Path}
	for _, root := range view.linkRoots {
		walkHTML(root, func(node *htmlNode) {
			if node.tag != "a" || hasHTMLAncestorClass(node, "md-nav--secondary") {
				return
			}
			linked, err := pageURL.Parse(strings.TrimSpace(node.attrs["href"]))
			if err != nil || !sameOrigin(pageURL, linked) || linked.RawQuery != "" || !likelyHTMLPath(linked.Path) {
				return
			}
			paths = append(paths, linked.Path)
		})
	}
	common := commonDocumentationRoot(paths)
	if common == "" {
		return "", errors.New("MkDocs primary navigation has no same-origin document root")
	}
	return common, nil
}

func hasHTMLAncestorClass(node *htmlNode, class string) bool {
	for current := node.parent; current != nil; current = current.parent {
		if hasHTMLClass(current, class) {
			return true
		}
	}
	return false
}

func commonDocumentationRoot(paths []string) string {
	var common []string
	for index, value := range paths {
		root := normalizeDocumentationRoot(value)
		segments := strings.Split(strings.Trim(root, "/"), "/")
		if root == "/" {
			segments = nil
		}
		if index == 0 {
			common = append(common, segments...)
			continue
		}
		limit := len(common)
		if len(segments) < limit {
			limit = len(segments)
		}
		matched := 0
		for matched < limit && common[matched] == segments[matched] {
			matched++
		}
		common = common[:matched]
	}
	if len(paths) == 0 {
		return ""
	}
	if len(common) == 0 {
		return "/"
	}
	return "/" + strings.Join(common, "/") + "/"
}

func normalizeDocumentationRoot(value string) string {
	directory := strings.HasSuffix(value, "/")
	cleaned := path.Clean("/" + strings.TrimPrefix(value, "/"))
	if !directory && path.Ext(cleaned) != "" {
		cleaned = path.Dir(cleaned)
	}
	if cleaned == "/" {
		return "/"
	}
	return strings.TrimSuffix(cleaned, "/") + "/"
}

func mkdocsSitemapInventory(ctx context.Context, pageURL *url.URL, siteRoot string, reader *sourceReader) ([]string, error) {
	sitemap := *pageURL
	sitemap.Path = path.Join(siteRoot, "sitemap.xml")
	sitemap.RawPath = ""
	sitemap.RawQuery = ""
	sitemap.Fragment = ""
	raw, provenance, err := reader.read(ctx, sitemap.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("fetch MkDocs sitemap %s: %w", sitemap.String(), err)
	}
	var document struct {
		Locations []string `xml:"url>loc"`
	}
	if err := xml.Unmarshal(raw, &document); err != nil {
		return nil, fmt.Errorf("parse MkDocs sitemap %s: %w", provenance, err)
	}
	if len(document.Locations) == 0 {
		return nil, fmt.Errorf("MkDocs sitemap %s contains no URLs", provenance)
	}
	if len(document.Locations) > maxHTMLSitemapURLs {
		return nil, fmt.Errorf("MkDocs sitemap %s exceeds %d URLs", provenance, maxHTMLSitemapURLs)
	}
	seen := make(map[string]bool, len(document.Locations))
	urls := make([]string, 0, len(document.Locations))
	for index, location := range document.Locations {
		linked, err := url.Parse(strings.TrimSpace(location))
		if err != nil || linked.Scheme != "http" && linked.Scheme != "https" || linked.Host == "" || linked.RawQuery != "" || !sameOrigin(pageURL, linked) || !withinDocumentationPath(linked.Path, siteRoot) || !likelyHTMLPath(linked.Path) {
			return nil, fmt.Errorf("MkDocs sitemap %s contains invalid URL at position %d", provenance, index)
		}
		linked.Fragment = ""
		canonical, err := canonicalHTTPURL(linked.String())
		if err != nil {
			return nil, fmt.Errorf("canonicalize MkDocs sitemap URL %q: %w", location, err)
		}
		if seen[canonical] {
			return nil, fmt.Errorf("MkDocs sitemap %s repeats URL %s", provenance, canonical)
		}
		seen[canonical] = true
		urls = append(urls, canonical)
	}
	sort.Strings(urls)
	return urls, nil
}

func mkdocsInventoryURL(value string) (string, error) {
	parsed, err := url.Parse(value)
	if err != nil {
		return "", err
	}
	parsed.RawQuery = ""
	return canonicalHTTPURL(parsed.String())
}

func isMkDocsFramework(framework string) bool {
	return framework == "mkdocs" || framework == "mkdocs-material"
}

func containsString(values []string, wanted string) bool {
	index := sort.SearchStrings(values, wanted)
	return index < len(values) && values[index] == wanted
}

func validateInventoryAliases(aliases map[string]string, generated map[string]bool) error {
	for alias, target := range aliases {
		seen := map[string]bool{alias: true}
		current := target
		for !generated[current] {
			if seen[current] {
				return fmt.Errorf("inventory alias cycle starts at %s", alias)
			}
			seen[current] = true
			next, exists := aliases[current]
			if !exists {
				return fmt.Errorf("inventory alias %s does not terminate at generated content", alias)
			}
			current = next
		}
	}
	return nil
}

func firstHTMLClass(root *htmlNode, class string) *htmlNode {
	var found *htmlNode
	walkHTML(root, func(node *htmlNode) {
		if found == nil && hasHTMLClass(node, class) {
			found = node
		}
	})
	return found
}

func firstHTMLAttribute(root *htmlNode, attribute, value string) *htmlNode {
	var found *htmlNode
	walkHTML(root, func(node *htmlNode) {
		if found == nil && strings.EqualFold(strings.TrimSpace(node.attrs[attribute]), value) {
			found = node
		}
	})
	return found
}

func firstHTMLClassPrefix(root *htmlNode, prefix string) *htmlNode {
	var found *htmlNode
	walkHTML(root, func(node *htmlNode) {
		if found != nil {
			return
		}
		for _, class := range strings.Fields(node.attrs["class"]) {
			if strings.HasPrefix(class, prefix) {
				found = node
				return
			}
		}
	})
	return found
}

func htmlClasses(root *htmlNode, class string) []*htmlNode {
	var found []*htmlNode
	walkHTML(root, func(node *htmlNode) {
		if hasHTMLClass(node, class) {
			found = append(found, node)
		}
	})
	return found
}

func hasHTMLClass(node *htmlNode, wanted string) bool {
	for _, class := range strings.Fields(node.attrs["class"]) {
		if class == wanted {
			return true
		}
	}
	return false
}

func htmlToMarkdown(root *htmlNode, base *url.URL, includeHeader bool) (string, string) {
	title := ""
	walkHTML(root, func(node *htmlNode) {
		if title == "" && node.tag == "h1" {
			title = renderHTMLInline(node, base, nil)
		}
		if title == "" && node.tag == "title" {
			title = cleanInline(htmlNodeText(node))
		}
	})
	var output strings.Builder
	renderHTMLBlocks(&output, root, base, includeHeader)
	markdown := cleanMarkdown(output.String())
	if title == "" {
		for _, line := range strings.Split(markdown, "\n") {
			if strings.HasPrefix(line, "# ") {
				title = strings.TrimSpace(strings.TrimPrefix(line, "# "))
				break
			}
		}
	}
	return title, markdown
}

func renderHTMLBlocks(output *strings.Builder, node *htmlNode, base *url.URL, includeHeader bool) {
	if htmlNodeHidden(node) || skippedHTMLTag(node.tag) && !(includeHeader && node.tag == "header") {
		return
	}
	switch node.tag {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		level := int(node.tag[1] - '0')
		writeMarkdownBlock(output, strings.Repeat("#", level)+" "+renderHTMLInline(node, base, nil))
		return
	case "p":
		writeMarkdownBlock(output, renderHTMLInline(node, base, nil))
		return
	case "pre":
		language := htmlCodeLanguage(node)
		if child := firstHTMLChild(node, "code"); child != nil {
			if language == "" {
				language = htmlCodeLanguage(child)
			}
		}
		code := strings.Trim(strings.ReplaceAll(htmlCodeText(node), "\r\n", "\n"), "\n")
		writeMarkdownBlock(output, "```"+language+"\n"+code+"\n```")
		return
	case "ul", "ol":
		writeMarkdownBlock(output, renderHTMLList(node, base, 0))
		return
	case "table":
		writeMarkdownBlock(output, renderHTMLTable(node, base))
		return
	case "a":
		if htmlHasBlockChild(node) {
			target := resolveHTMLReference(node.attrs["href"], base)
			for _, child := range node.children {
				if len(child.tag) == 2 && child.tag[0] == 'h' && child.tag[1] >= '1' && child.tag[1] <= '6' && target != "" {
					level := int(child.tag[1] - '0')
					label := renderHTMLInline(child, base, nil)
					writeMarkdownBlock(output, strings.Repeat("#", level)+" ["+label+"]("+target+")")
					continue
				}
				renderHTMLBlocks(output, child, base, includeHeader)
			}
			return
		}
		writeMarkdownBlock(output, renderHTMLInline(node, base, nil))
		return
	case "code":
		writeMarkdownBlock(output, renderHTMLInline(node, base, nil))
		return
	}
	for _, child := range node.children {
		renderHTMLBlocks(output, child, base, includeHeader)
	}
}

func renderHTMLInline(node *htmlNode, base *url.URL, excluded map[*htmlNode]bool) string {
	var output strings.Builder
	var render func(*htmlNode)
	render = func(current *htmlNode) {
		if excluded[current] || htmlNodeHidden(current) || skippedHTMLTag(current.tag) {
			return
		}
		if current.tag == "" {
			output.WriteString(current.text)
			return
		}
		switch current.tag {
		case "br":
			output.WriteString("  \n")
			return
		case "code":
			output.WriteByte('`')
			output.WriteString(strings.ReplaceAll(cleanInline(htmlNodeText(current)), "`", "\\`"))
			output.WriteByte('`')
			return
		case "strong", "b":
			output.WriteString("**")
			for _, child := range current.children {
				render(child)
			}
			output.WriteString("**")
			return
		case "em", "i":
			output.WriteByte('*')
			for _, child := range current.children {
				render(child)
			}
			output.WriteByte('*')
			return
		case "a":
			if hasHTMLClass(current, "hash-link") || hasHTMLClass(current, "headerlink") {
				return
			}
			if hasHTMLClass(current, "toclink") {
				for _, child := range current.children {
					render(child)
				}
				return
			}
			label := cleanInline(htmlNodeText(current))
			target := resolveHTMLReference(current.attrs["href"], base)
			if label != "" && target != "" {
				fmt.Fprintf(&output, "[%s](%s)", label, target)
				return
			}
		case "img":
			label := cleanInline(current.attrs["alt"])
			target := resolveHTMLReference(current.attrs["src"], base)
			if label != "" && target != "" {
				fmt.Fprintf(&output, "![%s](%s)", label, target)
			}
			return
		}
		for _, child := range current.children {
			render(child)
		}
	}
	render(node)
	return cleanInline(output.String())
}

func renderHTMLList(node *htmlNode, base *url.URL, depth int) string {
	var output strings.Builder
	itemNumber := 0
	for _, item := range node.children {
		if item.tag != "li" {
			continue
		}
		itemNumber++
		nested := make([]*htmlNode, 0, len(item.children))
		excluded := make(map[*htmlNode]bool)
		for _, child := range item.children {
			if child.tag == "ul" || child.tag == "ol" {
				nested = append(nested, child)
				excluded[child] = true
			}
		}
		marker := "- "
		if node.tag == "ol" {
			marker = fmt.Sprintf("%d. ", itemNumber)
		}
		output.WriteString(strings.Repeat("  ", depth) + marker + renderHTMLInline(item, base, excluded) + "\n")
		for _, child := range nested {
			if rendered := renderHTMLList(child, base, depth+1); rendered != "" {
				output.WriteString(rendered + "\n")
			}
		}
	}
	return strings.TrimRight(output.String(), "\n")
}

func renderHTMLTable(node *htmlNode, base *url.URL) string {
	var rows [][]string
	var collectRows func(*htmlNode)
	collectRows = func(current *htmlNode) {
		if current != node && current.tag == "table" {
			return
		}
		if current.tag == "tr" {
			var cells []string
			for _, child := range current.children {
				if child.tag == "th" || child.tag == "td" {
					cell := strings.ReplaceAll(renderHTMLInline(child, base, nil), "|", "\\|")
					cells = append(cells, strings.ReplaceAll(cell, "\n", " "))
				}
			}
			if len(cells) > 0 {
				rows = append(rows, cells)
			}
			return
		}
		for _, child := range current.children {
			collectRows(child)
		}
	}
	collectRows(node)
	if len(rows) == 0 {
		return ""
	}
	columns := 0
	for _, row := range rows {
		if len(row) > columns {
			columns = len(row)
		}
	}
	for index := range rows {
		for len(rows[index]) < columns {
			rows[index] = append(rows[index], "")
		}
	}
	var output strings.Builder
	writeTableRow := func(row []string) {
		output.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}
	writeTableRow(rows[0])
	delimiter := make([]string, columns)
	for index := range delimiter {
		delimiter[index] = "---"
	}
	writeTableRow(delimiter)
	for _, row := range rows[1:] {
		writeTableRow(row)
	}
	return strings.TrimRight(output.String(), "\n")
}

func writeMarkdownBlock(output *strings.Builder, block string) {
	block = strings.TrimSpace(block)
	if block == "" {
		return
	}
	if output.Len() > 0 {
		output.WriteString("\n\n")
	}
	output.WriteString(block)
}

func cleanMarkdown(value string) string {
	lines := strings.Split(strings.ReplaceAll(value, "\r", ""), "\n")
	for index := range lines {
		lines[index] = strings.TrimRightFunc(lines[index], unicode.IsSpace)
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func cleanInline(value string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(value), " "))
}

func htmlNodeText(node *htmlNode) string {
	if node.tag == "" {
		return node.text
	}
	var output strings.Builder
	for _, child := range node.children {
		output.WriteString(htmlNodeText(child))
	}
	return output.String()
}

func firstHTMLChild(node *htmlNode, tag string) *htmlNode {
	for _, child := range node.children {
		if child.tag == tag {
			return child
		}
	}
	return nil
}

func htmlNodeHidden(node *htmlNode) bool {
	_, hidden := node.attrs["hidden"]
	return hidden
}

func htmlHasBlockChild(node *htmlNode) bool {
	for _, child := range node.children {
		switch child.tag {
		case "h1", "h2", "h3", "h4", "h5", "h6", "p", "pre", "ul", "ol", "table", "article", "section", "div":
			return true
		}
	}
	return false
}

func htmlCodeLanguage(node *htmlNode) string {
	for current := node; current != nil; current = current.parent {
		if current != node && (current.tag == "article" || current.tag == "main" || current.tag == "section" || current.tag == "body" || current.tag == "html" || current.tag == "document" || hasHTMLClass(current, "md-content__inner")) {
			return ""
		}
		for _, class := range strings.Fields(current.attrs["class"]) {
			if strings.HasPrefix(class, "language-") {
				return strings.TrimPrefix(class, "language-")
			}
		}
	}
	return ""
}

func htmlCodeText(node *htmlNode) string {
	if node.tag == "" {
		return node.text
	}
	if node.tag == "br" {
		return "\n"
	}
	var output strings.Builder
	for _, child := range node.children {
		output.WriteString(htmlCodeText(child))
	}
	return output.String()
}

func skippedHTMLTag(tag string) bool {
	switch tag {
	case "script", "style", "noscript", "svg", "canvas", "template", "nav", "header", "footer", "aside", "form":
		return true
	default:
		return false
	}
}

func htmlPageLinks(roots []*htmlNode, base *url.URL, scope, scopePath string) []string {
	unique := make(map[string]bool)
	for _, root := range roots {
		walkHTML(root, func(node *htmlNode) {
			if node.tag != "a" {
				return
			}
			linked, err := base.Parse(strings.TrimSpace(node.attrs["href"]))
			if err != nil || linked.Host == "" || linked.Scheme != "http" && linked.Scheme != "https" || !sameOrigin(base, linked) || scope == "path" && !withinDocumentationPath(linked.Path, scopePath) {
				return
			}
			linked.Fragment = ""
			if likelyHTMLPath(linked.Path) {
				unique[linked.String()] = true
			}
		})
	}
	links := make([]string, 0, len(unique))
	for linked := range unique {
		links = append(links, linked)
	}
	sort.Strings(links)
	return links
}

func documentationPathScope(value string) string {
	directory := strings.HasSuffix(value, "/")
	value = path.Clean("/" + strings.TrimPrefix(value, "/"))
	if value == "/" {
		return "/"
	}
	if directory {
		return value + "/"
	}
	if path.Ext(value) != "" {
		return strings.TrimSuffix(path.Dir(value), "/") + "/"
	}
	segments := strings.Split(strings.Trim(value, "/"), "/")
	if len(segments) == 1 {
		return "/" + segments[0] + "/"
	}
	return "/" + strings.Join(segments[:len(segments)-1], "/") + "/"
}

func withinDocumentationPath(value, scope string) bool {
	cleaned := path.Clean("/" + strings.TrimPrefix(value, "/"))
	return scope == "/" || cleaned == strings.TrimSuffix(scope, "/") || strings.HasPrefix(cleaned+"/", scope)
}

func likelyHTMLPath(value string) bool {
	extension := strings.ToLower(path.Ext(value))
	switch extension {
	case "", ".html", ".htm", ".xhtml", ".php", ".asp", ".aspx":
		return true
	default:
		return false
	}
}

func sameOrigin(first, second *url.URL) bool {
	return strings.EqualFold(first.Scheme, second.Scheme) &&
		strings.EqualFold(first.Hostname(), second.Hostname()) && originPort(first) == originPort(second)
}

func originPort(value *url.URL) string {
	if value.Port() != "" {
		return value.Port()
	}
	if strings.EqualFold(value.Scheme, "http") {
		return "80"
	}
	if strings.EqualFold(value.Scheme, "https") {
		return "443"
	}
	return ""
}

func resolveHTMLReference(reference string, base *url.URL) string {
	reference = strings.TrimSpace(reference)
	if reference == "" {
		return ""
	}
	resolved, err := base.Parse(reference)
	if err != nil {
		return ""
	}
	if resolved.Scheme != "http" && resolved.Scheme != "https" {
		return ""
	}
	return resolved.String()
}

func pageTitleFromURL(source *url.URL) string {
	name := strings.TrimSuffix(path.Base(strings.TrimSuffix(source.Path, "/")), path.Ext(source.Path))
	name = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(name, "-", " "), "_", " "))
	if name == "" || name == "." {
		return "Documentation"
	}
	return name
}
