package importer

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

const (
	DefaultMaxHTMLPages = 50
	DefaultMaxHTMLDepth = 3
	maxHTMLPages        = 200
	maxHTMLDepth        = 10
	maxHTMLTreeDepth    = 128
	maxHTMLNodes        = 50_000
)

type htmlNode struct {
	tag      string
	attrs    map[string]string
	text     string
	children []*htmlNode
}

type crawledPage struct {
	title       string
	description string
	source      string
	markdown    string
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

	// A same-origin link must not escape through an HTTP redirect either.
	client := *options.HTTPClient
	previousRedirectPolicy := client.CheckRedirect
	client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if !sameOrigin(start, request.URL) {
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
	pages := make([]crawledPage, 0, options.MaxHTMLPages)
	rootSource := start.String()
	for len(queue) > 0 && len(pages) < options.MaxHTMLPages {
		pending := queue[0]
		queue = queue[1:]
		raw, provenance, readErr := reader.read(ctx, pending.source, nil)
		if readErr != nil {
			if len(pages) == 0 {
				return Result{}, readErr
			}
			continue
		}
		pageURL, parseErr := url.Parse(provenance)
		if parseErr != nil || !sameOrigin(start, pageURL) {
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
			continue
		}
		if len(pages) == 0 {
			rootSource = provenance
		}
		seen[provenance] = true
		title, markdown := htmlToMarkdown(document, pageURL)
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

		if pending.depth >= options.MaxHTMLDepth {
			continue
		}
		for _, linked := range htmlPageLinks(document, pageURL) {
			if len(seen) >= options.MaxHTMLPages || seen[linked] {
				continue
			}
			seen[linked] = true
			queue = append(queue, pendingPage{source: linked, depth: pending.depth + 1})
		}
	}
	if len(pages) == 0 {
		return Result{}, errors.New("HTML crawl produced no pages")
	}

	description := pages[0].description
	result, err := publish(ctx, options, name, version, func(stage string) error {
		metadata := manifest{
			Name: name, Version: version, Description: description,
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
	result.Kind, result.Source, result.Pages = "html", rootSource, len(pages)
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
	root := &htmlNode{tag: "document"}
	stack := []*htmlNode{root}
	nodes := 1
	text := string(raw)
	lower := strings.ToLower(text)
	for position := 0; position < len(text); {
		if text[position] != '<' {
			next := strings.IndexByte(text[position:], '<')
			if next < 0 {
				next = len(text) - position
			}
			if err := appendHTMLText(stack[len(stack)-1], text[position:position+next], &nodes); err != nil {
				return nil, err
			}
			position += next
			continue
		}
		if strings.HasPrefix(text[position:], "<!--") {
			end := strings.Index(text[position+4:], "-->")
			if end < 0 {
				break
			}
			position += end + 7
			continue
		}
		end := htmlTagEnd(text, position+1)
		if end < 0 {
			if err := appendHTMLText(stack[len(stack)-1], text[position:], &nodes); err != nil {
				return nil, err
			}
			break
		}
		rawTag := strings.TrimSpace(text[position+1 : end])
		position = end + 1
		if rawTag == "" || rawTag[0] == '!' || rawTag[0] == '?' {
			continue
		}
		if rawTag[0] == '/' {
			name := htmlTagName(rawTag[1:])
			for index := len(stack) - 1; index > 0; index-- {
				if stack[index].tag == name {
					stack = stack[:index]
					break
				}
			}
			continue
		}
		selfClosing := strings.HasSuffix(strings.TrimSpace(rawTag), "/")
		name := htmlTagName(rawTag)
		if name == "" {
			continue
		}
		if len(stack) > maxHTMLTreeDepth {
			return nil, fmt.Errorf("HTML document exceeds maximum tree depth of %d", maxHTMLTreeDepth)
		}
		nodes++
		if nodes > maxHTMLNodes {
			return nil, fmt.Errorf("HTML document exceeds maximum node count of %d", maxHTMLNodes)
		}
		node := &htmlNode{tag: name, attrs: htmlAttributes(rawTag[len(name):])}
		parent := stack[len(stack)-1]
		parent.children = append(parent.children, node)
		if selfClosing || htmlVoidElement(name) {
			continue
		}
		if name == "script" || name == "style" {
			closing := "</" + name
			relative := strings.Index(lower[position:], closing)
			if relative < 0 {
				if err := appendHTMLText(node, text[position:], &nodes); err != nil {
					return nil, err
				}
				break
			}
			if err := appendHTMLText(node, text[position:position+relative], &nodes); err != nil {
				return nil, err
			}
			position += relative
			continue
		}
		stack = append(stack, node)
	}
	return root, nil
}

func htmlTagEnd(value string, start int) int {
	var quote byte
	for index := start; index < len(value); index++ {
		character := value[index]
		if quote != 0 {
			if character == quote {
				quote = 0
			}
			continue
		}
		if character == '\'' || character == '"' {
			quote = character
		} else if character == '>' {
			return index
		}
	}
	return -1
}

func htmlTagName(value string) string {
	value = strings.TrimSpace(value)
	end := 0
	for end < len(value) && (unicode.IsLetter(rune(value[end])) || unicode.IsDigit(rune(value[end])) || value[end] == '-' || value[end] == ':') {
		end++
	}
	return strings.ToLower(value[:end])
}

func htmlAttributes(value string) map[string]string {
	attributes := make(map[string]string)
	for position := 0; position < len(value); {
		for position < len(value) && (unicode.IsSpace(rune(value[position])) || value[position] == '/') {
			position++
		}
		start := position
		for position < len(value) && !unicode.IsSpace(rune(value[position])) && value[position] != '=' && value[position] != '/' {
			position++
		}
		if start == position {
			position++
			continue
		}
		name := strings.ToLower(value[start:position])
		for position < len(value) && unicode.IsSpace(rune(value[position])) {
			position++
		}
		attributeValue := ""
		if position < len(value) && value[position] == '=' {
			position++
			for position < len(value) && unicode.IsSpace(rune(value[position])) {
				position++
			}
			if position < len(value) && (value[position] == '\'' || value[position] == '"') {
				quote := value[position]
				position++
				start = position
				for position < len(value) && value[position] != quote {
					position++
				}
				attributeValue = value[start:position]
				if position < len(value) {
					position++
				}
			} else {
				start = position
				for position < len(value) && !unicode.IsSpace(rune(value[position])) && value[position] != '/' {
					position++
				}
				attributeValue = value[start:position]
			}
		}
		attributes[name] = html.UnescapeString(attributeValue)
	}
	return attributes
}

func appendHTMLText(parent *htmlNode, value string, nodes *int) error {
	if value != "" {
		*nodes = *nodes + 1
		if *nodes > maxHTMLNodes {
			return fmt.Errorf("HTML document exceeds maximum node count of %d", maxHTMLNodes)
		}
		parent.children = append(parent.children, &htmlNode{text: html.UnescapeString(value)})
	}
	return nil
}

func htmlVoidElement(name string) bool {
	switch name {
	case "area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr":
		return true
	default:
		return false
	}
}

func isHTMLDocument(root *htmlNode) bool {
	found := false
	walkHTML(root, func(node *htmlNode) {
		if node.tag == "html" || node.tag == "body" || node.tag == "title" || node.tag == "main" || node.tag == "article" {
			found = true
		}
	})
	return found
}

func walkHTML(node *htmlNode, visit func(*htmlNode)) {
	visit(node)
	for _, child := range node.children {
		walkHTML(child, visit)
	}
}

func htmlToMarkdown(root *htmlNode, base *url.URL) (string, string) {
	title := ""
	walkHTML(root, func(node *htmlNode) {
		if title == "" && node.tag == "title" {
			title = cleanInline(htmlNodeText(node))
		}
	})
	var output strings.Builder
	renderHTMLBlocks(&output, root, base)
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

func renderHTMLBlocks(output *strings.Builder, node *htmlNode, base *url.URL) {
	if skippedHTMLTag(node.tag) {
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
		language := ""
		if child := firstHTMLChild(node, "code"); child != nil {
			for _, class := range strings.Fields(child.attrs["class"]) {
				if strings.HasPrefix(class, "language-") {
					language = strings.TrimPrefix(class, "language-")
				}
			}
		}
		code := strings.Trim(strings.ReplaceAll(htmlNodeText(node), "\r\n", "\n"), "\n")
		writeMarkdownBlock(output, "```"+language+"\n"+code+"\n```")
		return
	case "ul", "ol":
		writeMarkdownBlock(output, renderHTMLList(node, base, 0))
		return
	case "table":
		writeMarkdownBlock(output, renderHTMLTable(node, base))
		return
	case "a", "code":
		writeMarkdownBlock(output, renderHTMLInline(node, base, nil))
		return
	}
	for _, child := range node.children {
		renderHTMLBlocks(output, child, base)
	}
}

func renderHTMLInline(node *htmlNode, base *url.URL, excluded map[*htmlNode]bool) string {
	var output strings.Builder
	var render func(*htmlNode)
	render = func(current *htmlNode) {
		if excluded[current] || skippedHTMLTag(current.tag) {
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
			label := cleanInline(htmlNodeText(current))
			target := resolveHTMLReference(current.attrs["href"], base)
			if label != "" && target != "" {
				fmt.Fprintf(&output, "[%s](%s)", label, target)
				return
			}
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
			output.WriteString(renderHTMLList(child, base, depth+1))
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

func skippedHTMLTag(tag string) bool {
	switch tag {
	case "script", "style", "noscript", "svg", "canvas", "template", "nav", "header", "footer", "aside", "form":
		return true
	default:
		return false
	}
}

func htmlPageLinks(root *htmlNode, base *url.URL) []string {
	unique := make(map[string]bool)
	walkHTML(root, func(node *htmlNode) {
		if node.tag != "a" {
			return
		}
		linked, err := base.Parse(strings.TrimSpace(node.attrs["href"]))
		if err != nil || linked.Host == "" || linked.Scheme != "http" && linked.Scheme != "https" || !sameOrigin(base, linked) {
			return
		}
		linked.Fragment = ""
		if likelyHTMLPath(linked.Path) {
			unique[linked.String()] = true
		}
	})
	links := make([]string, 0, len(unique))
	for linked := range unique {
		links = append(links, linked)
	}
	sort.Strings(links)
	return links
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
