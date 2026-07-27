package importer

import (
	"bytes"
	"context"
	"encoding/json"
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
	sphinxBuilderOption     = regexp.MustCompile(`(?m)\bBUILDER\s*:\s*['"]([^'"]+)['"]`)
	sphinxFileSuffixOption  = regexp.MustCompile(`(?m)\bFILE_SUFFIX\s*:\s*['"]([^'"]*)['"]`)
	sphinxLinkSuffixOption  = regexp.MustCompile(`(?m)\bLINK_SUFFIX\s*:\s*['"]([^'"]*)['"]`)
	vitepressGenerator      = regexp.MustCompile(`(?i)^vitepress v[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9a-z.-]+)?(?:\+[0-9a-z.-]+)?$`)
	starlightGenerator      = regexp.MustCompile(`(?i)^starlight v[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9a-z.-]+)?(?:\+[0-9a-z.-]+)?$`)
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

type sphinxRuntime struct {
	builder    string
	fileSuffix string
	linkSuffix string
}

type starlightLocaleScope struct {
	language     string
	root         string
	excluded     []string
	multilingual bool
}

type starlightSitemapAlternate struct {
	Rel      string `xml:"rel,attr"`
	Language string `xml:"hreflang,attr"`
	Href     string `xml:"href,attr"`
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
		inventoryIdentity, _ := mkdocsInventoryURL(provenance)
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
			if finiteInventory == nil {
				if canonicalErr != nil || len(pages) != 0 || pending.depth >= 10 || !withinDocumentationPath(redirected.Path, scopePath) {
					return Result{}, fmt.Errorf("HTML crawl did not complete: invalid initial refresh on %s", provenance)
				}
				canonicalAlias, aliasErr := canonicalHTTPURL(provenance)
				if aliasErr != nil || canonicalAlias == canonicalRedirect || processed[canonicalRedirect] {
					return Result{}, fmt.Errorf("HTML crawl did not complete: refresh cycle on %s", provenance)
				}
				processed[provenance] = true
				seen[canonicalRedirect] = true
				queue = append([]pendingPage{{source: canonicalRedirect, depth: pending.depth + 1}}, queue...)
				continue
			}
			if canonicalErr != nil || !finiteInventory[canonicalRedirect] {
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
		if len(pages) == 0 && isFiniteInventoryFramework(rootFramework) {
			var siteRoot string
			var inventory []string
			var inventoryErr error
			switch {
			case isMkDocsFramework(rootFramework):
				siteRoot, inventoryErr = mkdocsSiteRoot(view, pageURL)
				if inventoryErr == nil {
					scopePath = siteRoot
					inventory, inventoryErr = htmlSitemapInventory(ctx, "MkDocs", pageURL, siteRoot, siteRoot, reader)
				}
			case rootFramework == "vitepress":
				siteRoot, inventoryErr = vitepressSiteRoot(document, pageURL)
				if inventoryErr == nil {
					scopePath = siteRoot
					inventory, inventoryErr = htmlSitemapInventory(ctx, "VitePress", pageURL, siteRoot, siteRoot, reader)
				}
				if inventoryErr == nil {
					canonicalStart, canonicalErr := mkdocsInventoryURL(provenance)
					if canonicalErr == nil && !containsString(inventory, canonicalStart) {
						alias, aliasErr := vitepressStartInventoryAlias(pageURL, siteRoot, inventory)
						if aliasErr == nil {
							inventoryAliases[canonicalStart] = alias
							inventoryIdentity = alias
						}
					}
				}
			case rootFramework == "nextra":
				var deploymentRoot string
				deploymentRoot, inventoryErr = nextraDeploymentRoot(document, pageURL)
				if inventoryErr == nil {
					siteRoot, inventoryErr = nextraDocumentationRoot(view, pageURL)
				}
				if inventoryErr == nil {
					scopePath = siteRoot
					inventory, inventoryErr = htmlSitemapInventory(ctx, "Nextra", pageURL, deploymentRoot, siteRoot, reader)
				}
				if inventoryErr == nil {
					canonicalStart, canonicalErr := mkdocsInventoryURL(provenance)
					if canonicalErr == nil && !containsString(inventory, canonicalStart) {
						alias, aliasErr := vitepressStartInventoryAlias(pageURL, siteRoot, inventory)
						if aliasErr == nil {
							inventoryAliases[canonicalStart] = alias
							inventoryIdentity = alias
						}
					}
				}
			case rootFramework == "astro-starlight":
				var sitemapURL *url.URL
				var localeScope starlightLocaleScope
				sitemapURL, inventoryErr = starlightSitemapURL(document, pageURL)
				if inventoryErr == nil {
					localeScope, inventoryErr = starlightLocale(document, pageURL, sitemapURL)
				}
				if inventoryErr == nil {
					siteRoot = normalizeDocumentationRoot(path.Dir(sitemapURL.Path))
					scopePath = localeScope.root
					inventory, inventoryErr = starlightSitemapInventory(ctx, pageURL, sitemapURL, siteRoot, localeScope, reader)
				}
				if inventoryErr == nil {
					canonicalStart, canonicalErr := mkdocsInventoryURL(provenance)
					if canonicalErr == nil && !containsString(inventory, canonicalStart) {
						alias, aliasErr := starlightCanonicalURL(document, pageURL)
						if aliasErr != nil || !containsString(inventory, alias) {
							alias, aliasErr = vitepressStartInventoryAlias(pageURL, localeScope.root, inventory)
						}
						if aliasErr == nil {
							inventoryAliases[canonicalStart] = alias
							inventoryIdentity = alias
						}
					}
				}
			case rootFramework == "sphinx":
				var runtime sphinxRuntime
				siteRoot, runtime, inventoryErr = sphinxSite(ctx, document, pageURL, reader)
				if inventoryErr == nil {
					scopePath = siteRoot
					inventory, inventoryErr = sphinxSearchInventory(ctx, pageURL, siteRoot, runtime, reader)
				}
				if inventoryErr == nil {
					canonicalStart, canonicalErr := mkdocsInventoryURL(provenance)
					if canonicalErr == nil && !containsString(inventory, canonicalStart) {
						alias, aliasErr := sphinxStartInventoryAlias(pageURL, siteRoot, runtime)
						if aliasErr == nil && containsString(inventory, alias) {
							inventoryAliases[canonicalStart] = alias
							inventoryIdentity = alias
						}
					}
				}
			}
			if inventoryErr != nil {
				return Result{}, fmt.Errorf("%s crawl did not complete: %w", rootFramework, inventoryErr)
			}
			finiteInventory = make(map[string]bool, len(inventory))
			for _, linked := range inventory {
				finiteInventory[linked] = true
			}
			canonicalProvenance := inventoryIdentity
			_, canonicalErr := canonicalHTTPURL(canonicalProvenance)
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
		if rootFramework == "astro-starlight" {
			canonical, canonicalErr := starlightCanonicalURL(document, pageURL)
			if canonicalErr != nil || canonical != inventoryIdentity {
				return Result{}, fmt.Errorf("astro-starlight crawl did not complete: page %s has an invalid canonical URL", provenance)
			}
		}
		title, markdown := htmlToMarkdown(view.content, pageURL, view.includeHeader)
		if (rootFramework == "vitepress" || rootFramework == "nextra" || rootFramework == "astro-starlight") && strings.TrimSpace(markdown) == "" {
			return Result{}, fmt.Errorf("%s crawl did not complete: page has no statically rendered content: %s", rootFramework, provenance)
		}
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
			if inventoryIdentity != "" {
				generatedInventory[inventoryIdentity] = true
			}
		}
		reportProgress(options, Progress{Stage: "page", Framework: rootFramework, URL: provenance, Pages: len(pages), Queued: len(queue)})

		var linkedPages []string
		if finiteInventory == nil {
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
	case "sphinx":
		view.content = firstHTMLAttribute(root, "role", "main")
		if view.content == nil {
			view.content = firstHTMLTag(root, "main")
		}
		if view.content == nil {
			view.content = firstHTMLClass(root, "body")
		}
		if view.content == nil {
			view.profileError = errors.New("Sphinx page has no main content container")
			return view
		}
		view.linkRoots = append(htmlClasses(root, "sphinxsidebar"), htmlClasses(root, "wy-menu-vertical")...)
	case "vitepress":
		if home := firstHTMLClass(root, "VPHome"); home != nil {
			view.content = home
		} else if doc := firstHTMLClass(root, "VPDoc"); doc != nil {
			view.content = firstHTMLClass(doc, "vp-doc")
		} else if doc := firstHTMLClass(root, "VPContentDoc"); doc != nil {
			view.content = firstHTMLClass(doc, "vt-doc")
		} else if page := firstHTMLClass(root, "VPContentPage"); page != nil {
			view.content = firstHTMLTag(page, "main")
		} else if page := firstHTMLClass(root, "marketing-layout"); page != nil {
			if shell := lastHTMLChild(page, "div"); shell != nil {
				view.content = firstHTMLChild(shell, "div")
			}
		} else if page := firstHTMLClass(root, "VPPage"); page != nil {
			view.content = page
		} else {
			view.content = nil
		}
		if view.content == nil {
			view.profileError = errors.New("VitePress page has no supported default-theme content container")
		}
		view.linkRoots = nil
	case "nextra":
		if content := firstHTMLAttribute(root, "data-pagefind-body", "true"); content != nil {
			view.content = content
			view.linkRoots = htmlClasses(root, "nextra-sidebar")
		} else if article := firstHTMLClass(root, "nextra-content"); article != nil {
			view.content = firstHTMLChild(article, "main")
			view.linkRoots = htmlClasses(root, "nextra-sidebar-container")
		}
		if view.content == nil {
			view.profileError = errors.New("Nextra page has no supported docs-theme content container")
		}
	case "astro-starlight":
		view.content = firstHTMLAttributePresent(root, "data-pagefind-body")
		if sidebar := firstHTMLAttribute(root, "id", "starlight__sidebar"); sidebar != nil {
			view.linkRoots = []*htmlNode{sidebar}
		} else {
			view.linkRoots = nil
		}
		if view.content == nil || firstHTMLTag(view.content, "h1") == nil || firstHTMLClass(view.content, "sl-markdown-content") == nil {
			view.profileError = errors.New("Astro Starlight page has no supported statically rendered content container")
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
			case vitepressGenerator.MatchString(generator):
				framework = "vitepress"
				return
			case starlightGenerator.MatchString(generator):
				framework = "astro-starlight"
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
	if framework == "" && looksLikeSphinx(root) {
		framework = "sphinx"
	}
	if framework == "" && looksLikeNextra(root) {
		framework = "nextra"
	}
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

func looksLikeSphinx(root *htmlNode) bool {
	if firstHTMLAttribute(root, "role", "main") == nil && firstHTMLTag(root, "main") == nil && firstHTMLClass(root, "body") == nil {
		return false
	}
	return htmlAssetSuffix(root, "script", "src", "_static/documentation_options.js") && htmlAssetSuffix(root, "script", "src", "_static/doctools.js")
}

func looksLikeNextra(root *htmlNode) bool {
	if !htmlAssetPathContains(root, "/_next/static/") {
		return false
	}
	if firstHTMLAttribute(root, "data-pagefind-body", "true") != nil && firstHTMLAttribute(root, "id", "nextra-skip-nav") != nil && len(htmlClasses(root, "nextra-sidebar")) > 0 {
		return true
	}
	article := firstHTMLClass(root, "nextra-content")
	return article != nil && firstHTMLChild(article, "main") != nil && len(htmlClasses(root, "nextra-sidebar-container")) > 0
}

func htmlAssetPathContains(root *htmlNode, marker string) bool {
	found := false
	walkHTML(root, func(node *htmlNode) {
		if found || node.tag != "script" && node.tag != "link" {
			return
		}
		value := node.attrs["src"]
		if value == "" {
			value = node.attrs["href"]
		}
		parsed, err := url.Parse(strings.TrimSpace(value))
		found = err == nil && strings.Contains(parsed.Path, marker)
	})
	return found
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

func htmlSitemapInventory(ctx context.Context, framework string, pageURL *url.URL, siteRoot, documentationRoot string, reader *sourceReader) ([]string, error) {
	sitemap := *pageURL
	sitemap.Path = path.Join(siteRoot, "sitemap.xml")
	sitemap.RawPath = ""
	sitemap.RawQuery = ""
	sitemap.Fragment = ""
	raw, provenance, err := reader.read(ctx, sitemap.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("fetch %s sitemap %s: %w", framework, sitemap.String(), err)
	}
	var document struct {
		Locations []string `xml:"url>loc"`
	}
	if err := xml.Unmarshal(raw, &document); err != nil {
		return nil, fmt.Errorf("parse %s sitemap %s: %w", framework, provenance, err)
	}
	if len(document.Locations) == 0 {
		return nil, fmt.Errorf("%s sitemap %s contains no URLs", framework, provenance)
	}
	if len(document.Locations) > maxHTMLSitemapURLs {
		return nil, fmt.Errorf("%s sitemap %s exceeds %d URLs", framework, provenance, maxHTMLSitemapURLs)
	}
	seen := make(map[string]bool, len(document.Locations))
	urls := make([]string, 0, len(document.Locations))
	for index, location := range document.Locations {
		linked, err := url.Parse(strings.TrimSpace(location))
		if err != nil || linked.Scheme != "http" && linked.Scheme != "https" || linked.Host == "" || linked.RawQuery != "" || !sameOrigin(pageURL, linked) || !withinDocumentationPath(linked.Path, siteRoot) || !likelyHTMLPath(linked.Path) {
			return nil, fmt.Errorf("%s sitemap %s contains invalid URL at position %d", framework, provenance, index)
		}
		linked.Fragment = ""
		canonical, err := canonicalHTTPURL(linked.String())
		if err != nil {
			return nil, fmt.Errorf("canonicalize %s sitemap URL %q: %w", framework, location, err)
		}
		if seen[canonical] {
			return nil, fmt.Errorf("%s sitemap %s repeats URL %s", framework, provenance, canonical)
		}
		seen[canonical] = true
		if withinDocumentationPath(linked.Path, documentationRoot) {
			urls = append(urls, canonical)
		}
	}
	if len(urls) == 0 {
		return nil, fmt.Errorf("%s sitemap %s contains no URLs under %s", framework, provenance, documentationRoot)
	}
	sort.Strings(urls)
	return urls, nil
}

func starlightSitemapURL(document *htmlNode, pageURL *url.URL) (*url.URL, error) {
	var found []*url.URL
	walkHTML(document, func(node *htmlNode) {
		if node.tag != "link" || !hasHTMLToken(node.attrs["rel"], "sitemap") {
			return
		}
		linked, err := pageURL.Parse(strings.TrimSpace(node.attrs["href"]))
		if err == nil && sameOrigin(pageURL, linked) && linked.RawQuery == "" && linked.Fragment == "" {
			found = append(found, linked)
		}
	})
	if len(found) != 1 {
		return nil, errors.New("Astro Starlight page has no unique same-origin sitemap index")
	}
	return found[0], nil
}

func starlightCanonicalURL(document *htmlNode, pageURL *url.URL) (string, error) {
	var found []string
	walkHTML(document, func(node *htmlNode) {
		if node.tag != "link" || !hasHTMLToken(node.attrs["rel"], "canonical") {
			return
		}
		linked, err := pageURL.Parse(strings.TrimSpace(node.attrs["href"]))
		if err != nil || !sameOrigin(pageURL, linked) || linked.RawQuery != "" || linked.Fragment != "" {
			return
		}
		canonical, err := canonicalHTTPURL(linked.String())
		if err == nil {
			found = append(found, canonical)
		}
	})
	if len(found) != 1 {
		return "", errors.New("Astro Starlight page has no unique same-origin canonical URL")
	}
	return found[0], nil
}

func starlightLocale(document *htmlNode, pageURL, sitemapURL *url.URL) (starlightLocaleScope, error) {
	siteRoot := normalizeDocumentationRoot(path.Dir(sitemapURL.Path))
	canonical, err := starlightCanonicalURL(document, pageURL)
	if err != nil {
		return starlightLocaleScope{}, err
	}
	canonicalURL, err := url.Parse(canonical)
	if err != nil || !withinDocumentationPath(canonicalURL.Path, siteRoot) {
		return starlightLocaleScope{}, errors.New("Astro Starlight canonical URL is outside the sitemap deployment root")
	}

	seenLanguages := make(map[string]bool)
	seenURLs := make(map[string]bool)
	var alternatePaths []string
	currentFound := false
	currentLanguage := ""
	var alternateErr error
	walkHTML(document, func(node *htmlNode) {
		if alternateErr != nil || node.tag != "link" || !hasHTMLToken(node.attrs["rel"], "alternate") {
			return
		}
		language := strings.ToLower(strings.TrimSpace(node.attrs["hreflang"]))
		if language == "x-default" {
			return
		}
		if language == "" || seenLanguages[language] {
			alternateErr = errors.New("Astro Starlight page has an invalid or repeated hreflang language")
			return
		}
		linked, parseErr := pageURL.Parse(strings.TrimSpace(node.attrs["href"]))
		if parseErr != nil || !sameOrigin(pageURL, linked) || linked.RawQuery != "" || linked.Fragment != "" || !withinDocumentationPath(linked.Path, siteRoot) || !likelyHTMLPath(linked.Path) {
			alternateErr = errors.New("Astro Starlight page has an invalid hreflang URL")
			return
		}
		alternate, canonicalErr := canonicalHTTPURL(linked.String())
		if canonicalErr != nil || seenURLs[alternate] {
			alternateErr = errors.New("Astro Starlight page has a repeated hreflang URL")
			return
		}
		seenLanguages[language] = true
		seenURLs[alternate] = true
		alternatePaths = append(alternatePaths, linked.Path)
		if alternate == canonical {
			currentFound = true
			currentLanguage = language
		}
	})
	if alternateErr != nil {
		return starlightLocaleScope{}, alternateErr
	}
	if !currentFound || len(alternatePaths) == 0 {
		return starlightLocaleScope{}, errors.New("Astro Starlight page has no valid hreflang alternate for its canonical URL")
	}

	suffix := commonPathSuffix(alternatePaths, siteRoot)
	if len(suffix) == 0 {
		return starlightLocaleScope{}, errors.New("Astro Starlight hreflang alternates have no common route suffix")
	}
	prefixes := make(map[string]bool)
	currentRoot := ""
	for _, alternatePath := range alternatePaths {
		segments := documentationPathSegments(alternatePath, siteRoot)
		prefix := segments[:len(segments)-len(suffix)]
		localeRoot := siteRoot
		if len(prefix) > 0 {
			localeRoot = normalizeDocumentationRoot(path.Join(siteRoot, path.Join(prefix...)))
		}
		prefixes[localeRoot] = true
		if path.Clean(alternatePath) == path.Clean(canonicalURL.Path) {
			currentRoot = localeRoot
		}
	}
	if currentRoot == "" {
		return starlightLocaleScope{}, errors.New("Astro Starlight canonical URL has no locale root")
	}
	var excluded []string
	for prefix := range prefixes {
		if prefix != currentRoot {
			excluded = append(excluded, prefix)
		}
	}
	sort.Strings(excluded)
	return starlightLocaleScope{language: currentLanguage, root: currentRoot, excluded: excluded, multilingual: len(alternatePaths) > 1}, nil
}

func documentationPathSegments(value, root string) []string {
	cleaned := strings.Trim(path.Clean(value), "/")
	root = strings.Trim(path.Clean(root), "/")
	if root != "" {
		cleaned = strings.TrimPrefix(cleaned, root)
		cleaned = strings.TrimPrefix(cleaned, "/")
	}
	if cleaned == "" {
		return nil
	}
	return strings.Split(cleaned, "/")
}

func commonPathSuffix(paths []string, root string) []string {
	var common []string
	for index, value := range paths {
		segments := documentationPathSegments(value, root)
		if index == 0 {
			common = append(common, segments...)
			continue
		}
		matched := 0
		for matched < len(common) && matched < len(segments) && common[len(common)-1-matched] == segments[len(segments)-1-matched] {
			matched++
		}
		common = common[len(common)-matched:]
	}
	return common
}

func starlightSitemapInventory(ctx context.Context, pageURL, sitemapURL *url.URL, siteRoot string, locale starlightLocaleScope, reader *sourceReader) ([]string, error) {
	raw, provenance, err := reader.read(ctx, sitemapURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("fetch Astro Starlight sitemap index %s: %w", sitemapURL.String(), err)
	}
	wantedIndex, _ := canonicalHTTPURL(sitemapURL.String())
	actualIndex, _ := canonicalHTTPURL(provenance)
	if wantedIndex == "" || actualIndex != wantedIndex {
		return nil, errors.New("Astro Starlight sitemap index redirected")
	}
	var index struct {
		XMLName  xml.Name
		Sitemaps []struct {
			Locations []string `xml:"loc"`
		} `xml:"sitemap"`
	}
	if err := xml.Unmarshal(raw, &index); err != nil || index.XMLName.Local != "sitemapindex" {
		return nil, fmt.Errorf("parse Astro Starlight sitemap index %s: invalid sitemap index", provenance)
	}
	if len(index.Sitemaps) == 0 || len(index.Sitemaps) > maxHTMLSitemapURLs {
		return nil, fmt.Errorf("Astro Starlight sitemap index %s has an invalid shard count", provenance)
	}

	seenShards := make(map[string]bool, len(index.Sitemaps))
	seenPages := make(map[string]bool)
	var inventory []string
	for shardIndex, entry := range index.Sitemaps {
		if len(entry.Locations) != 1 || strings.TrimSpace(entry.Locations[0]) == "" {
			return nil, fmt.Errorf("Astro Starlight sitemap index %s contains invalid shard record at position %d", provenance, shardIndex)
		}
		shardURL, parseErr := url.Parse(strings.TrimSpace(entry.Locations[0]))
		if parseErr != nil || shardURL.Scheme != "http" && shardURL.Scheme != "https" || shardURL.Host == "" || shardURL.RawQuery != "" || shardURL.Fragment != "" || !sameOrigin(pageURL, shardURL) || !withinDocumentationPath(shardURL.Path, siteRoot) {
			return nil, fmt.Errorf("Astro Starlight sitemap index %s contains invalid shard at position %d", provenance, shardIndex)
		}
		canonicalShard, canonicalErr := canonicalHTTPURL(shardURL.String())
		if canonicalErr != nil || seenShards[canonicalShard] {
			return nil, fmt.Errorf("Astro Starlight sitemap index %s repeats a shard", provenance)
		}
		seenShards[canonicalShard] = true
		shardRaw, shardProvenance, readErr := reader.read(ctx, canonicalShard, nil)
		if readErr != nil {
			return nil, fmt.Errorf("fetch Astro Starlight sitemap shard %s: %w", canonicalShard, readErr)
		}
		actualShard, _ := canonicalHTTPURL(shardProvenance)
		if actualShard != canonicalShard {
			return nil, fmt.Errorf("Astro Starlight sitemap shard %s redirected", canonicalShard)
		}
		var shard struct {
			XMLName xml.Name
			URLs    []struct {
				Locations  []string                    `xml:"loc"`
				Alternates []starlightSitemapAlternate `xml:"link"`
			} `xml:"url"`
		}
		if err := xml.Unmarshal(shardRaw, &shard); err != nil || shard.XMLName.Local != "urlset" || len(shard.URLs) == 0 {
			return nil, fmt.Errorf("parse Astro Starlight sitemap shard %s: invalid URL set", shardProvenance)
		}
		for pageIndex, page := range shard.URLs {
			if len(page.Locations) != 1 || strings.TrimSpace(page.Locations[0]) == "" {
				return nil, fmt.Errorf("Astro Starlight sitemap shard %s contains invalid URL record at position %d", shardProvenance, pageIndex)
			}
			location := page.Locations[0]
			linked, parseErr := url.Parse(strings.TrimSpace(location))
			if parseErr != nil || linked.Scheme != "http" && linked.Scheme != "https" || linked.Host == "" || linked.RawQuery != "" || linked.Fragment != "" || !sameOrigin(pageURL, linked) || !withinDocumentationPath(linked.Path, siteRoot) || !likelyHTMLPath(linked.Path) {
				return nil, fmt.Errorf("Astro Starlight sitemap shard %s contains invalid URL at position %d", shardProvenance, pageIndex)
			}
			canonicalPage, canonicalErr := canonicalHTTPURL(linked.String())
			if canonicalErr != nil || seenPages[canonicalPage] {
				return nil, fmt.Errorf("Astro Starlight sitemap repeats URL %s", canonicalPage)
			}
			seenPages[canonicalPage] = true
			if len(seenPages) > maxHTMLSitemapURLs {
				return nil, fmt.Errorf("Astro Starlight sitemap exceeds %d URLs", maxHTMLSitemapURLs)
			}
			selectedLocale, alternateErr := starlightSitemapRecordLocale(page.Alternates, canonicalPage, pageURL, siteRoot, locale)
			if alternateErr != nil {
				return nil, fmt.Errorf("Astro Starlight sitemap shard %s has invalid alternates at position %d: %w", shardProvenance, pageIndex, alternateErr)
			}
			if selectedLocale {
				inventory = append(inventory, canonicalPage)
			}
		}
	}
	if len(inventory) == 0 {
		return nil, fmt.Errorf("Astro Starlight sitemap contains no URLs under locale root %s", locale.root)
	}
	sort.Strings(inventory)
	return inventory, nil
}

func starlightSitemapRecordLocale(alternates []starlightSitemapAlternate, canonicalPage string, pageURL *url.URL, siteRoot string, locale starlightLocaleScope) (bool, error) {
	if len(alternates) == 0 && !locale.multilingual {
		parsed, err := url.Parse(canonicalPage)
		return err == nil && withinStarlightLocale(parsed.Path, locale), nil
	}
	if len(alternates) == 0 {
		return false, errors.New("multilingual URL record has no alternates")
	}
	seenLanguages := make(map[string]bool)
	seenURLs := make(map[string]bool)
	selected := ""
	matchedSelf := false
	for _, alternate := range alternates {
		language := strings.ToLower(strings.TrimSpace(alternate.Language))
		linked, err := url.Parse(strings.TrimSpace(alternate.Href))
		if !strings.EqualFold(strings.TrimSpace(alternate.Rel), "alternate") || language == "" || err != nil || linked.Scheme != "http" && linked.Scheme != "https" || linked.Host == "" || linked.RawQuery != "" || linked.Fragment != "" || !sameOrigin(pageURL, linked) || !withinDocumentationPath(linked.Path, siteRoot) || !likelyHTMLPath(linked.Path) {
			return false, errors.New("invalid hreflang alternate")
		}
		canonical, err := canonicalHTTPURL(linked.String())
		if err != nil || seenLanguages[language] {
			return false, errors.New("repeated hreflang language or URL")
		}
		seenLanguages[language] = true
		if language == "x-default" {
			continue
		}
		if seenURLs[canonical] {
			return false, errors.New("repeated hreflang language or URL")
		}
		seenURLs[canonical] = true
		matchedSelf = matchedSelf || canonical == canonicalPage
		if language == locale.language {
			selected = canonical
		}
	}
	if !matchedSelf {
		return false, errors.New("URL record does not identify its own locale")
	}
	parsed, err := url.Parse(canonicalPage)
	return err == nil && selected == canonicalPage && withinStarlightLocale(parsed.Path, locale), nil
}

func withinStarlightLocale(value string, locale starlightLocaleScope) bool {
	if !withinDocumentationPath(value, locale.root) {
		return false
	}
	for _, excluded := range locale.excluded {
		if withinDocumentationPath(value, excluded) {
			return false
		}
	}
	return true
}

func vitepressSiteRoot(document *htmlNode, pageURL *url.URL) (string, error) {
	roots := make(map[string]bool)
	walkHTML(document, func(node *htmlNode) {
		if node.tag != "link" || !hasHTMLToken(node.attrs["rel"], "stylesheet") {
			return
		}
		linked, err := pageURL.Parse(strings.TrimSpace(node.attrs["href"]))
		if err != nil || !sameOrigin(pageURL, linked) || path.Base(linked.Path) != "vp-icons.css" || linked.RawQuery != "" {
			return
		}
		root := normalizeDocumentationRoot(strings.TrimSuffix(linked.Path, "vp-icons.css"))
		if withinDocumentationPath(pageURL.Path, root) {
			roots[root] = true
		}
	})
	if len(roots) == 0 {
		return "", errors.New("VitePress page has no same-origin vp-icons.css site-root marker")
	}
	if len(roots) != 1 {
		return "", errors.New("VitePress page has conflicting site-root markers")
	}
	for root := range roots {
		return root, nil
	}
	panic("unreachable")
}

func nextraDeploymentRoot(document *htmlNode, pageURL *url.URL) (string, error) {
	roots := make(map[string]bool)
	walkHTML(document, func(node *htmlNode) {
		if node.tag != "script" && node.tag != "link" {
			return
		}
		value := node.attrs["src"]
		if value == "" {
			value = node.attrs["href"]
		}
		linked, err := pageURL.Parse(strings.TrimSpace(value))
		if err != nil || !sameOrigin(pageURL, linked) {
			return
		}
		index := strings.Index(linked.Path, "/_next/static/")
		if index >= 0 {
			roots[normalizeDocumentationRoot(linked.Path[:index+1])] = true
		}
	})
	if len(roots) != 1 {
		return "", errors.New("Nextra page has no unique same-origin _next/static deployment root")
	}
	for root := range roots {
		return root, nil
	}
	panic("unreachable")
}

func nextraDocumentationRoot(view htmlDocumentView, pageURL *url.URL) (string, error) {
	candidateRoot := documentationPathScope(strings.TrimSuffix(pageURL.Path, "/"))
	paths := []string{pageURL.Path}
	for _, navigation := range view.linkRoots {
		walkHTML(navigation, func(node *htmlNode) {
			if node.tag != "a" {
				return
			}
			linked, err := pageURL.Parse(strings.TrimSpace(node.attrs["href"]))
			if err == nil && sameOrigin(pageURL, linked) && linked.RawQuery == "" && likelyHTMLPath(linked.Path) && withinDocumentationPath(linked.Path, candidateRoot) {
				paths = append(paths, linked.Path)
			}
		})
	}
	if len(paths) < 2 {
		return "", errors.New("Nextra sidebar has no same-origin documentation routes")
	}
	root := commonDocumentationRoot(paths)
	if root == "" || root == "/" || !withinDocumentationPath(pageURL.Path, root) {
		return "", errors.New("Nextra sidebar does not define a scoped documentation root")
	}
	return root, nil
}

func hasHTMLToken(value, wanted string) bool {
	for _, token := range strings.Fields(value) {
		if strings.EqualFold(token, wanted) {
			return true
		}
	}
	return false
}

func vitepressStartInventoryAlias(pageURL *url.URL, siteRoot string, inventory []string) (string, error) {
	base := *pageURL
	base.RawPath = ""
	base.RawQuery = ""
	base.Fragment = ""
	paths := make(map[string]bool)
	add := func(value string) {
		candidate := base
		candidate.Path = value
		canonical, err := canonicalHTTPURL(candidate.String())
		if err == nil && withinDocumentationPath(candidate.Path, siteRoot) && containsString(inventory, canonical) {
			paths[canonical] = true
		}
	}
	if strings.HasSuffix(base.Path, "/index.html") {
		add(strings.TrimSuffix(base.Path, "index.html"))
	}
	if strings.HasSuffix(base.Path, "/index") {
		add(strings.TrimSuffix(base.Path, "index"))
	}
	if strings.HasSuffix(base.Path, "/") {
		add(path.Join(base.Path, "index.html"))
	} else if strings.HasSuffix(base.Path, ".html") {
		add(strings.TrimSuffix(base.Path, ".html"))
	} else {
		add(base.Path + ".html")
	}
	if len(paths) != 1 {
		return "", errors.New("VitePress starting page has no unique sitemap alias")
	}
	for candidate := range paths {
		return candidate, nil
	}
	panic("unreachable")
}

func mkdocsInventoryURL(value string) (string, error) {
	parsed, err := url.Parse(value)
	if err != nil {
		return "", err
	}
	parsed.RawQuery = ""
	return canonicalHTTPURL(parsed.String())
}

func sphinxSite(ctx context.Context, document *htmlNode, pageURL *url.URL, reader *sourceReader) (string, sphinxRuntime, error) {
	assetURL := htmlAssetURL(document, pageURL, "script", "src", "_static/documentation_options.js")
	if assetURL == nil || !sameOrigin(pageURL, assetURL) {
		return "", sphinxRuntime{}, errors.New("Sphinx page has no same-origin documentation_options.js")
	}
	const optionsPath = "_static/documentation_options.js"
	if !strings.HasSuffix(assetURL.Path, optionsPath) {
		return "", sphinxRuntime{}, errors.New("Sphinx documentation options path is invalid")
	}
	siteRoot := normalizeDocumentationRoot(strings.TrimSuffix(assetURL.Path, optionsPath))

	var advertisedRoot string
	walkHTML(document, func(node *htmlNode) {
		if advertisedRoot != "" {
			return
		}
		candidate := ""
		if node.tag == "html" {
			candidate = strings.TrimSpace(node.attrs["data-content_root"])
		}
		if node.tag == "script" && candidate == "" {
			candidate = strings.TrimSpace(node.attrs["data-url_root"])
		}
		if candidate == "" {
			return
		}
		resolved, err := pageURL.Parse(candidate)
		if err == nil && sameOrigin(pageURL, resolved) && resolved.RawQuery == "" && resolved.Fragment == "" {
			advertisedRoot = normalizeDocumentationRoot(resolved.Path)
		}
	})
	if advertisedRoot != "" && advertisedRoot != siteRoot {
		return "", sphinxRuntime{}, fmt.Errorf("Sphinx content root %s conflicts with asset root %s", advertisedRoot, siteRoot)
	}

	raw, provenance, err := reader.read(ctx, assetURL.String(), nil)
	if err != nil {
		return "", sphinxRuntime{}, fmt.Errorf("fetch Sphinx documentation options %s: %w", assetURL, err)
	}
	resolved, err := url.Parse(provenance)
	if err != nil || !sameOrigin(pageURL, resolved) || resolved.Path != assetURL.Path {
		return "", sphinxRuntime{}, errors.New("Sphinx documentation options redirected outside its expected URL")
	}
	runtime, err := parseSphinxRuntime(raw)
	if err != nil {
		return "", sphinxRuntime{}, err
	}
	return siteRoot, runtime, nil
}

func parseSphinxRuntime(raw []byte) (sphinxRuntime, error) {
	value := string(raw)
	builder := sphinxBuilderOption.FindStringSubmatch(value)
	if len(builder) != 2 || builder[1] != "html" && builder[1] != "dirhtml" {
		return sphinxRuntime{}, errors.New("Sphinx documentation options use an unsupported builder")
	}
	fileSuffix := sphinxFileSuffixOption.FindStringSubmatch(value)
	if len(fileSuffix) != 2 || invalidSphinxSuffix(fileSuffix[1]) {
		return sphinxRuntime{}, errors.New("Sphinx documentation options have an invalid file suffix")
	}
	linkSuffix := sphinxLinkSuffixOption.FindStringSubmatch(value)
	if len(linkSuffix) != 2 {
		linkSuffix = fileSuffix
	}
	if linkSuffix[1] != "/" && invalidSphinxSuffix(linkSuffix[1]) {
		return sphinxRuntime{}, errors.New("Sphinx documentation options have an invalid link suffix")
	}
	return sphinxRuntime{builder: builder[1], fileSuffix: fileSuffix[1], linkSuffix: linkSuffix[1]}, nil
}

func invalidSphinxSuffix(value string) bool {
	return strings.ContainsAny(value, `/\\?#`) || strings.IndexFunc(value, unicode.IsControl) >= 0
}

func sphinxSearchInventory(ctx context.Context, pageURL *url.URL, siteRoot string, runtime sphinxRuntime, reader *sourceReader) ([]string, error) {
	indexURL := *pageURL
	indexURL.Path = path.Join(siteRoot, "searchindex.js")
	indexURL.RawPath = ""
	indexURL.RawQuery = ""
	indexURL.Fragment = ""
	raw, provenance, err := reader.read(ctx, indexURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("fetch Sphinx search index %s: %w", indexURL.String(), err)
	}
	resolved, err := url.Parse(provenance)
	if err != nil || !sameOrigin(pageURL, resolved) || resolved.Path != indexURL.Path {
		return nil, errors.New("Sphinx search index redirected outside its expected URL")
	}
	docnames, err := parseSphinxSearchIndex(raw)
	if err != nil {
		return nil, err
	}
	if len(docnames) > maxHTMLSitemapURLs {
		return nil, fmt.Errorf("Sphinx search index exceeds %d documents", maxHTMLSitemapURLs)
	}
	rootURL := *pageURL
	rootURL.Path = siteRoot
	rootURL.RawPath = ""
	rootURL.RawQuery = ""
	rootURL.Fragment = ""
	seen := make(map[string]bool, len(docnames))
	urls := make([]string, 0, len(docnames))
	for index, docname := range docnames {
		linked, buildErr := sphinxDocumentURL(&rootURL, docname, runtime)
		if buildErr != nil {
			return nil, fmt.Errorf("Sphinx search index has invalid docname at position %d: %w", index, buildErr)
		}
		canonical, canonicalErr := canonicalHTTPURL(linked)
		linkedURL, parseErr := url.Parse(linked)
		if canonicalErr != nil || parseErr != nil || !sameOrigin(&rootURL, linkedURL) || !withinDocumentationPath(linkedURL.Path, siteRoot) {
			return nil, fmt.Errorf("Sphinx search index has invalid URL at position %d", index)
		}
		if seen[canonical] {
			return nil, fmt.Errorf("Sphinx search index repeats generated URL %s", canonical)
		}
		seen[canonical] = true
		urls = append(urls, canonical)
	}
	sort.Strings(urls)
	return urls, nil
}

func parseSphinxSearchIndex(raw []byte) ([]string, error) {
	value := strings.TrimSpace(string(raw))
	const prefix = "Search.setIndex("
	if !strings.HasPrefix(value, prefix) {
		return nil, errors.New("Sphinx search index uses an unsupported legacy format")
	}
	value = strings.TrimSpace(strings.TrimPrefix(value, prefix))
	if strings.HasSuffix(value, ");") {
		value = strings.TrimSpace(strings.TrimSuffix(value, ");"))
	} else if strings.HasSuffix(value, ")") {
		value = strings.TrimSpace(strings.TrimSuffix(value, ")"))
	} else {
		return nil, errors.New("Sphinx search index has an invalid wrapper")
	}
	var index struct {
		DocNames []string `json:"docnames"`
		Titles   []string `json:"titles"`
	}
	if err := json.Unmarshal([]byte(value), &index); err != nil {
		return nil, fmt.Errorf("parse Sphinx search index: %w", err)
	}
	if len(index.DocNames) == 0 {
		return nil, errors.New("Sphinx search index contains no docnames")
	}
	if len(index.Titles) != 0 && len(index.Titles) != len(index.DocNames) {
		return nil, errors.New("Sphinx search index title count does not match docnames")
	}
	return index.DocNames, nil
}

func sphinxDocumentURL(rootURL *url.URL, docname string, runtime sphinxRuntime) (string, error) {
	if docname == "" || strings.HasPrefix(docname, "/") || strings.Contains(docname, "\\") || strings.IndexFunc(docname, unicode.IsControl) >= 0 {
		return "", errors.New("unsafe document name")
	}
	for _, segment := range strings.Split(docname, "/") {
		if segment == "" || segment == "." || segment == ".." {
			return "", errors.New("unsafe document path segment")
		}
	}
	linked := *rootURL
	switch runtime.builder {
	case "html":
		linked.Path = path.Join(rootURL.Path, docname+runtime.fileSuffix)
	case "dirhtml":
		directory := ""
		if docname != "index" {
			directory = strings.TrimSuffix(docname, "/index")
		}
		linked.Path = normalizeDocumentationRoot(path.Join(rootURL.Path, directory))
	default:
		return "", errors.New("unsupported Sphinx builder")
	}
	linked.RawPath = ""
	return linked.String(), nil
}

func sphinxStartInventoryAlias(pageURL *url.URL, siteRoot string, runtime sphinxRuntime) (string, error) {
	aliased := *pageURL
	aliased.RawPath = ""
	aliased.RawQuery = ""
	aliased.Fragment = ""
	switch runtime.builder {
	case "html":
		if runtime.linkSuffix == "/" && aliased.Path != siteRoot {
			if !strings.HasSuffix(aliased.Path, "/") || !strings.HasPrefix(aliased.Path, siteRoot) {
				return "", errors.New("Sphinx HTML start URL does not use the configured link suffix")
			}
			relative := strings.TrimSuffix(strings.TrimPrefix(aliased.Path, siteRoot), "/")
			aliased.Path = path.Join(siteRoot, relative+runtime.fileSuffix)
		} else if strings.HasSuffix(aliased.Path, "/") {
			aliased.Path = path.Join(aliased.Path, "index"+runtime.fileSuffix)
		} else {
			if !strings.HasPrefix(aliased.Path, siteRoot) {
				return "", errors.New("Sphinx HTML start URL is outside the documentation root")
			}
			relative := strings.TrimPrefix(aliased.Path, siteRoot)
			if runtime.linkSuffix != "" && !strings.HasSuffix(relative, runtime.linkSuffix) {
				return "", errors.New("Sphinx HTML start URL does not use the configured link suffix")
			}
			relative = strings.TrimSuffix(relative, runtime.linkSuffix)
			aliased.Path = path.Join(siteRoot, relative+runtime.fileSuffix)
		}
	case "dirhtml":
		indexName := "index" + runtime.fileSuffix
		if path.Base(aliased.Path) != indexName {
			return "", errors.New("Sphinx dirhtml start URL is not an index alias")
		}
		aliased.Path = normalizeDocumentationRoot(path.Dir(aliased.Path))
	default:
		return "", errors.New("unsupported Sphinx builder")
	}
	if !withinDocumentationPath(aliased.Path, siteRoot) {
		return "", errors.New("Sphinx start alias is outside the documentation root")
	}
	return canonicalHTTPURL(aliased.String())
}

func htmlAssetURL(root *htmlNode, base *url.URL, tag, attribute, suffix string) *url.URL {
	var found *url.URL
	walkHTML(root, func(node *htmlNode) {
		if found != nil || node.tag != tag {
			return
		}
		linked, err := base.Parse(strings.TrimSpace(node.attrs[attribute]))
		if err == nil && strings.HasSuffix(linked.Path, suffix) {
			linked.RawQuery = ""
			linked.Fragment = ""
			found = linked
		}
	})
	return found
}

func isMkDocsFramework(framework string) bool {
	return framework == "mkdocs" || framework == "mkdocs-material"
}

func isFiniteInventoryFramework(framework string) bool {
	return isMkDocsFramework(framework) || framework == "sphinx" || framework == "vitepress" || framework == "nextra" || framework == "astro-starlight"
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

func firstHTMLTag(root *htmlNode, tag string) *htmlNode {
	var found *htmlNode
	walkHTML(root, func(node *htmlNode) {
		if found == nil && node.tag == tag {
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

func firstHTMLAttributePresent(root *htmlNode, attribute string) *htmlNode {
	var found *htmlNode
	walkHTML(root, func(node *htmlNode) {
		if found != nil {
			return
		}
		if _, exists := node.attrs[attribute]; exists {
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
	semanticAside := node.tag == "aside" && (hasHTMLClass(node, "footnote") || hasHTMLClass(node, "footnote-list") || hasHTMLClass(node, "starlight-aside") || strings.EqualFold(node.attrs["role"], "doc-footnote"))
	if htmlNodeHidden(node) || skippedHTMLTag(node.tag) && !(includeHeader && node.tag == "header") && !semanticAside {
		return
	}
	if node.tag == "" {
		return
	}
	if strings.EqualFold(node.attrs["role"], "tabpanel") && !htmlHasBlockChild(node) {
		writeMarkdownBlock(output, renderHTMLInline(node, base, nil))
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
		if strings.TrimSpace(code) == "" {
			return
		}
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
	case "dt":
		if hasHTMLClass(node, "sig") {
			writeMarkdownBlock(output, markdownCodeSpan(sphinxSignatureText(node)))
			return
		}
		writeMarkdownBlock(output, "**"+renderHTMLInline(node, base, nil)+"**")
		return
	case "aside":
		if semanticAside {
			label := firstHTMLClass(node, "label")
			if label != nil {
				writeMarkdownBlock(output, "**"+cleanInline(htmlNodeText(label))+"**")
			}
			for _, child := range node.children {
				if child != label {
					renderHTMLBlocks(output, child, base, includeHeader)
				}
			}
			return
		}
	}
	for _, child := range node.children {
		renderHTMLBlocks(output, child, base, includeHeader)
	}
}

func sphinxSignatureText(node *htmlNode) string {
	var output strings.Builder
	var render func(*htmlNode)
	render = func(current *htmlNode) {
		if current.tag == "a" && (hasHTMLClass(current, "headerlink") || hasHTMLClass(current, "viewcode-link") || cleanInline(htmlNodeText(current)) == "[source]") {
			return
		}
		if current.tag == "" {
			output.WriteString(current.text)
			return
		}
		if current.tag == "br" {
			output.WriteByte(' ')
			return
		}
		for _, child := range current.children {
			render(child)
		}
	}
	render(node)
	return cleanInline(output.String())
}

func markdownCodeSpan(value string) string {
	longest := 0
	for index := 0; index < len(value); {
		if value[index] != '`' {
			index++
			continue
		}
		end := index
		for end < len(value) && value[end] == '`' {
			end++
		}
		longest = max(longest, end-index)
		index = end
	}
	fence := strings.Repeat("`", longest+1)
	if strings.HasPrefix(value, "`") || strings.HasSuffix(value, "`") || strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") {
		value = " " + value + " "
	}
	return fence + value + fence
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
			if hasHTMLClass(current, "hash-link") || hasHTMLClass(current, "headerlink") || hasHTMLClass(current, "header-anchor") || hasHTMLClass(current, "sl-anchor-link") {
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
		if htmlListItemHasBlocks(item) {
			var blocks strings.Builder
			for _, child := range item.children {
				if !excluded[child] {
					renderHTMLBlocks(&blocks, child, base, false)
				}
			}
			lines := strings.Split(cleanMarkdown(blocks.String()), "\n")
			indent := strings.Repeat("  ", depth)
			continuation := indent + strings.Repeat(" ", len(marker))
			if len(lines) > 0 && lines[0] != "" {
				output.WriteString(indent + marker + lines[0] + "\n")
				for _, line := range lines[1:] {
					if line == "" {
						output.WriteByte('\n')
					} else {
						output.WriteString(continuation + line + "\n")
					}
				}
			}
		} else {
			output.WriteString(strings.Repeat("  ", depth) + marker + renderHTMLInline(item, base, excluded) + "\n")
		}
		for _, child := range nested {
			if rendered := renderHTMLList(child, base, depth+1); rendered != "" {
				output.WriteString(rendered + "\n")
			}
		}
	}
	return strings.TrimRight(output.String(), "\n")
}

func htmlListItemHasBlocks(item *htmlNode) bool {
	for _, child := range item.children {
		switch child.tag {
		case "h1", "h2", "h3", "h4", "h5", "h6", "p", "pre", "table", "article", "section", "div", "aside":
			return true
		}
	}
	return false
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

func lastHTMLChild(node *htmlNode, tag string) *htmlNode {
	for index := len(node.children) - 1; index >= 0; index-- {
		if node.children[index].tag == tag {
			return node.children[index]
		}
	}
	return nil
}

func htmlNodeHidden(node *htmlNode) bool {
	_, hidden := node.attrs["hidden"]
	_, pagefindIgnored := node.attrs["data-pagefind-ignore"]
	_, algoliaExcluded := node.attrs["data-algolia-exclude"]
	starlightBanner := pagefindIgnored && hasHTMLClass(node, "sl-banner")
	starlightPromotion := algoliaExcluded && node.parent != nil && hasHTMLClass(node.parent, "hide-when-toc-is-visible")
	starlightMobileDuplicate := hasHTMLClass(node, "mobile-only") && hasHTMLClass(node, "not-content") && node.parent != nil && hasHTMLClass(node.parent, "hero")
	return hidden || starlightBanner || starlightPromotion || starlightMobileDuplicate
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
		if language := strings.TrimSpace(current.attrs["data-language"]); language != "" {
			return language
		}
		for _, class := range strings.Fields(current.attrs["class"]) {
			if strings.HasPrefix(class, "language-") {
				return strings.TrimPrefix(class, "language-")
			}
			if strings.HasPrefix(class, "highlight-") {
				language := strings.TrimPrefix(class, "highlight-")
				if language != "" && language != "default" && language != "none" {
					return language
				}
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
	value := output.String()
	if hasHTMLClass(node, "ec-line") {
		return strings.TrimRight(value, "\n") + "\n"
	}
	return value
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
