package importer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
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
)

type htmlNode struct {
	tag          string
	attrs        map[string]string
	text         string
	children     []*htmlNode
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

	// A same-origin link must not escape through an HTTP redirect either.
	client := *options.HTTPClient
	previousRedirectPolicy := client.CheckRedirect
	client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if !sameOrigin(start, request.URL) || options.HTMLScope == "path" && !withinDocumentationPath(request.URL.Path, scopePath) {
			return errors.New("HTML crawl redirect crosses the source origin")
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
	crawlLimited := false
	var crawlErrors []error
	for len(queue) > 0 && (options.MaxHTMLPages < 0 || len(pages) < options.MaxHTMLPages) {
		if err := ctx.Err(); err != nil {
			return Result{}, err
		}
		pending := queue[0]
		queue = queue[1:]
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
		if processed[provenance] {
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
		seen[provenance] = true
		processed[provenance] = true
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
		reportProgress(options, Progress{Stage: "page", Framework: rootFramework, URL: provenance, Pages: len(pages), Queued: len(queue)})

		linkedPages := htmlPageLinks(view.linkRoots, pageURL, options.HTMLScope, scopePath)
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
				converted = &htmlNode{tag: strings.ToLower(current.Data), attrs: attributes}
			case xhtml.TextNode:
				if current.Data == "" {
					continue
				}
				converted = &htmlNode{text: current.Data}
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
	if view.framework != "docusaurus" {
		return view
	}
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
	return view
}

func detectHTMLFramework(root *htmlNode) string {
	framework := ""
	walkHTML(root, func(node *htmlNode) {
		if framework != "" {
			return
		}
		if node.tag == "meta" && strings.EqualFold(strings.TrimSpace(node.attrs["name"]), "generator") && strings.HasPrefix(strings.ToLower(strings.TrimSpace(node.attrs["content"])), "docusaurus") {
			framework = "docusaurus"
			return
		}
		if node.attrs["id"] == "__docusaurus" || hasHTMLClass(node, "docs-doc-page") && hasHTMLClass(node, "plugin-docs") {
			framework = "docusaurus"
		}
	})
	return framework
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
			if hasHTMLClass(current, "hash-link") {
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
	for _, class := range strings.Fields(node.attrs["class"]) {
		if strings.HasPrefix(class, "language-") {
			return strings.TrimPrefix(class, "language-")
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
