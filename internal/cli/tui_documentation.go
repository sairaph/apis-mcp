package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/importer"
	"github.com/sairaph/apis-mcp/library"
)

type documentationState struct {
	snapshot *library.Snapshot
	owned    bool

	collections      []library.Collection
	apis             []library.API
	collectionCursor int
	catalogCursor    int
	focus            int
	filter           textinput.Model
	filtering        bool

	selected documentChoice
	frames   []documentFrame
	reader   documentReader
	importer *importForm
}

type documentChoice struct {
	name        string
	description string
	collections []string
	version     library.APIVersion
}

type documentEntry struct {
	path *library.Path
	page *library.Page
}

type documentFrame struct {
	path         string
	entries      []documentEntry
	cursor       int
	focus        int
	search       textinput.Model
	searching    bool
	results      []library.SearchHit
	resultCursor int
	loading      bool
}

type documentReader struct {
	result       library.ReadResult
	viewport     viewport.Model
	wrappedWidth int
	targetLine   int
}

type readPayload struct {
	result     library.ReadResult
	targetLine int
}

type importForm struct {
	inputs []textinput.Model
	focus  int
}

type pagesPayload struct {
	path    string
	entries []documentEntry
}

type reloadPayload struct {
	snapshot    *library.Snapshot
	collections []library.Collection
	apis        []library.API
	imported    *importer.Result
}

func newDocumentationState(runtime *bootstrap.Runtime) *documentationState {
	filter := newTextInput("filter API name or version", "")
	state := &documentationState{filter: filter}
	if runtime != nil {
		state.snapshot = runtime.Library
	}
	return state
}

func newTextInput(placeholder, value string) textinput.Model {
	input := textinput.New()
	input.Prompt = ""
	input.Placeholder = placeholder
	input.PromptStyle = styleTitle
	input.TextStyle = styleInput
	input.Cursor.Style = styleTitle
	input.SetValue(value)
	return input
}

func (d *documentationState) editing() bool {
	return d.filtering || d.importer != nil || len(d.frames) > 0 && d.currentFrame().searching
}

func (d *documentationState) close() {
	if d.owned && d.snapshot != nil {
		_ = d.snapshot.Close()
	}
	d.snapshot, d.owned = nil, false
}

func (d *documentationState) resize(width, height int) {
	d.filter.Width = max(8, width/3)
	if len(d.frames) > 0 {
		d.currentFrame().search.Width = max(8, width/3)
	}
	if d.importer != nil {
		for index := range d.importer.inputs {
			d.importer.inputs[index].Width = max(12, width-28)
		}
	}
	if d.reader.result.PageID != "" {
		d.resizeReader(max(10, width-4), max(3, height-8))
	}
}

func (d *documentationState) currentFrame() *documentFrame {
	return &d.frames[len(d.frames)-1]
}

func (d *documentationState) collectionID() string {
	if d.collectionCursor <= 0 || d.collectionCursor > len(d.collections) {
		return ""
	}
	return d.collections[d.collectionCursor-1].Collection
}

func (d *documentationState) choices() []documentChoice {
	collection := d.collectionID()
	query := strings.ToLower(strings.TrimSpace(d.filter.Value()))
	var choices []documentChoice
	for _, api := range d.apis {
		if collection != "" && !containsFold(api.Collections, collection) {
			continue
		}
		for _, version := range api.Versions {
			if query != "" && !strings.Contains(strings.ToLower(api.Name+" "+version.Version), query) {
				continue
			}
			choices = append(choices, documentChoice{
				name: api.Name, description: api.Description,
				collections: api.Collections, version: version,
			})
		}
	}
	return choices
}

func (d *documentationState) clampCatalog() {
	choices := d.choices()
	if len(choices) == 0 {
		d.catalogCursor = 0
	} else {
		d.catalogCursor = max(0, min(d.catalogCursor, len(choices)-1))
	}
}

func (m *model) updateDocumentation(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	d := m.docs
	switch m.currentScreen() {
	case screenDocumentReader:
		if key.String() == "esc" || key.String() == "backspace" {
			m.pop()
			return m, nil
		}
		var cmd tea.Cmd
		d.reader.viewport, cmd = d.reader.viewport.Update(key)
		return m, cmd
	case screenDocumentImport:
		return m.updateImport(key)
	case screenDocumentBrowser:
		return m.updateDocumentBrowser(key)
	default:
		return m.updateDocumentCatalog(key)
	}
}

func (m *model) updateDocumentCatalog(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	d := m.docs
	if d.filtering {
		switch key.String() {
		case "esc", "enter":
			d.filtering = false
			d.filter.Blur()
			d.clampCatalog()
			return m, nil
		case "tab", "shift+tab":
			d.filtering = false
			d.filter.Blur()
			d.focus = 1
			return m, nil
		}
		var cmd tea.Cmd
		d.filter, cmd = d.filter.Update(key)
		d.clampCatalog()
		return m, cmd
	}
	switch key.String() {
	case "tab", "shift+tab", "left", "right", "h", "l":
		d.focus = 1 - d.focus
	case "/":
		d.filtering = true
		return m, d.filter.Focus()
	case "up", "k":
		if d.focus == 0 {
			d.collectionCursor = (d.collectionCursor - 1 + len(d.collections) + 1) % (len(d.collections) + 1)
			d.catalogCursor = 0
		} else if choices := d.choices(); len(choices) > 0 {
			d.catalogCursor = (d.catalogCursor - 1 + len(choices)) % len(choices)
		}
	case "down", "j":
		if d.focus == 0 {
			d.collectionCursor = (d.collectionCursor + 1) % (len(d.collections) + 1)
			d.catalogCursor = 0
		} else if choices := d.choices(); len(choices) > 0 {
			d.catalogCursor = (d.catalogCursor + 1) % len(choices)
		}
	case "enter":
		if d.focus == 0 {
			d.focus = 1
			return m, nil
		}
		choices := d.choices()
		if len(choices) == 0 {
			return m, nil
		}
		d.selected = choices[d.catalogCursor]
		return m, m.openDocumentPath("")
	case "i":
		m.openImport()
	case "r":
		return m, m.reloadLibraryCmd()
	}
	return m, nil
}

func (m *model) updateDocumentBrowser(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	d, frame := m.docs, m.docs.currentFrame()
	if frame.searching {
		switch key.String() {
		case "esc":
			frame.searching = false
			frame.search.Blur()
			return m, nil
		case "enter":
			query := strings.TrimSpace(frame.search.Value())
			if query == "" {
				m.status, m.severity = "Enter terms to search this documentation path", severityWarning
				return m, nil
			}
			frame.searching = false
			frame.search.Blur()
			return m, m.searchDocumentCmd(query, frame.path)
		}
		var cmd tea.Cmd
		frame.search, cmd = frame.search.Update(key)
		return m, cmd
	}
	switch key.String() {
	case "esc", "backspace":
		d.frames = d.frames[:len(d.frames)-1]
		m.pop()
	case "tab", "shift+tab", "left", "right", "h", "l":
		frame.focus = 1 - frame.focus
	case "s", "/":
		frame.searching = true
		frame.focus = 1
		return m, frame.search.Focus()
	case "up", "k":
		if frame.focus == 0 && len(frame.entries) > 0 {
			frame.cursor = (frame.cursor - 1 + len(frame.entries)) % len(frame.entries)
		} else if frame.focus == 1 && len(frame.results) > 0 {
			frame.resultCursor = (frame.resultCursor - 1 + len(frame.results)) % len(frame.results)
		}
	case "down", "j":
		if frame.focus == 0 && len(frame.entries) > 0 {
			frame.cursor = (frame.cursor + 1) % len(frame.entries)
		} else if frame.focus == 1 && len(frame.results) > 0 {
			frame.resultCursor = (frame.resultCursor + 1) % len(frame.results)
		}
	case "enter":
		if frame.focus == 1 {
			if len(frame.results) > 0 {
				hit := frame.results[frame.resultCursor]
				return m, m.readDocumentCmd(hit.PageID, hit.Line)
			}
			frame.searching = true
			return m, frame.search.Focus()
		}
		if len(frame.entries) == 0 {
			return m, nil
		}
		entry := frame.entries[frame.cursor]
		if entry.path != nil {
			return m, m.openDocumentPath(entry.path.Path)
		}
		return m, m.readDocumentCmd(entry.page.PageID)
	case "i":
		m.openImport()
	case "r":
		return m, m.reloadLibraryCmd()
	}
	return m, nil
}

func (m *model) openDocumentPath(path string) tea.Cmd {
	search := newTextInput("search this path", "")
	m.docs.frames = append(m.docs.frames, documentFrame{path: path, search: search, loading: true})
	m.push(screenDocumentBrowser)
	snapshot, docID := m.docs.snapshot, m.docs.selected.version.DocID
	return m.startOperation("pages", "Opening documentation hierarchy", true, func(ctx context.Context, id int64) tea.Msg {
		entries, err := loadAllPages(ctx, snapshot, docID, path)
		return asyncMsg{id: id, kind: "pages", value: pagesPayload{path: path, entries: entries}, err: err}
	})
}

func loadAllPages(ctx context.Context, snapshot *library.Snapshot, docID, path string) ([]documentEntry, error) {
	var entries []documentEntry
	for page := 1; ; page++ {
		result, err := snapshot.Pages(ctx, library.PagesRequest{DocID: docID, Path: path, Page: page})
		if err != nil {
			return nil, err
		}
		for index := range result.Paths {
			item := result.Paths[index]
			entries = append(entries, documentEntry{path: &item})
		}
		for index := range result.Pages {
			item := result.Pages[index]
			entries = append(entries, documentEntry{page: &item})
		}
		if page >= result.TotalPages {
			break
		}
	}
	return entries, nil
}

func (d *documentationState) applyPages(payload pagesPayload) {
	frame := d.currentFrame()
	frame.path, frame.entries, frame.loading = payload.path, payload.entries, false
	frame.cursor = 0
}

func (m *model) searchDocumentCmd(query, path string) tea.Cmd {
	snapshot, docID := m.docs.snapshot, m.docs.selected.version.DocID
	return m.startOperation("search", "Searching documentation", true, func(ctx context.Context, id int64) tea.Msg {
		var hits []library.SearchHit
		var total int
		for page := 1; ; page++ {
			result, err := snapshot.Search(ctx, library.SearchRequest{DocID: docID, Query: query, Path: path, Page: page})
			if err != nil {
				return asyncMsg{id: id, kind: "search", err: err}
			}
			hits = append(hits, result.Hits...)
			total = result.Total
			if page >= result.TotalPages {
				return asyncMsg{id: id, kind: "search", value: library.SearchResult{
					DocID: docID, Query: query, Path: path,
					Pagination: library.Pagination{Page: 1, Total: total, TotalPages: result.TotalPages}, Hits: hits,
				}}
			}
		}
	})
}

func (m *model) readDocumentCmd(pageID string, targetLine ...int) tea.Cmd {
	snapshot, docID := m.docs.snapshot, m.docs.selected.version.DocID
	target := 0
	if len(targetLine) > 0 {
		target = targetLine[0]
	}
	return m.startOperation("read", "Opening documentation page", true, func(ctx context.Context, id int64) tea.Msg {
		result, err := readFullDocument(ctx, snapshot, docID, pageID)
		return asyncMsg{id: id, kind: "read", value: readPayload{result: result, targetLine: target}, err: err}
	})
}

func readFullDocument(ctx context.Context, snapshot *library.Snapshot, docID, pageID string) (library.ReadResult, error) {
	result, err := snapshot.Read(ctx, library.ReadRequest{DocID: docID, PageID: pageID})
	if err != nil || !result.Truncated {
		return result, err
	}
	var markdown strings.Builder
	markdown.WriteString(result.Markdown)
	end := result.Lines[1]
	for end < result.TotalLines {
		chunk, readErr := snapshot.Read(ctx, library.ReadRequest{DocID: docID, PageID: pageID, Lines: []int{end + 1, result.TotalLines}})
		if readErr != nil {
			return library.ReadResult{}, readErr
		}
		if chunk.Lines[1] <= end {
			return library.ReadResult{}, fmt.Errorf("documentation reader made no progress after line %d", end)
		}
		markdown.WriteString(chunk.Markdown)
		end = chunk.Lines[1]
	}
	result.Markdown = markdown.String()
	result.Lines[1], result.Truncated = end, false
	return result, nil
}

func (d *documentationState) openReader(result library.ReadResult, targetLine, width, height int) {
	d.reader.result = result
	d.reader.targetLine = targetLine
	d.reader.viewport = viewport.New(max(10, width), max(3, height))
	d.reader.viewport.MouseWheelEnabled = false
	d.reader.wrappedWidth = 0
	d.resizeReader(width, height)
}

func (d *documentationState) resizeReader(width, height int) {
	d.reader.viewport.Width = max(10, width)
	d.reader.viewport.Height = max(3, height)
	if d.reader.wrappedWidth == width {
		return
	}
	content := strings.TrimSpace(safeMultiline(d.reader.result.Markdown))
	physical := strings.Split(content, "\n")
	wrapped := make([]string, 0, len(physical))
	targetOffset := 0
	for index, line := range physical {
		if index+1 == d.reader.targetLine {
			targetOffset = len(wrapped)
		}
		wrapped = append(wrapped, wrapLines(line, max(10, width))...)
	}
	d.reader.viewport.SetContent(strings.Join(wrapped, "\n"))
	if d.reader.targetLine > 0 {
		d.reader.viewport.SetYOffset(max(0, targetOffset-height/3))
	}
	d.reader.wrappedWidth = width
}

func (m *model) reloadLibraryCmd() tea.Cmd {
	if m.runtime == nil {
		m.status, m.severity = "Documentation runtime is unavailable", severityError
		return nil
	}
	runtime := m.runtime
	return m.startOperation("reload", "Rebuilding and reopening documentation", true, func(ctx context.Context, id int64) tea.Msg {
		if err := runtime.RebuildLibrary(ctx); err != nil {
			return asyncMsg{id: id, kind: "reload", err: err}
		}
		snapshot, err := library.Open(ctx, libraryOptions(runtime))
		if err != nil {
			return asyncMsg{id: id, kind: "reload", err: err}
		}
		collections, apis, err := loadCatalog(ctx, snapshot)
		if err != nil {
			_ = snapshot.Close()
			return asyncMsg{id: id, kind: "reload", err: err}
		}
		if err := ctx.Err(); err != nil {
			_ = snapshot.Close()
			return asyncMsg{id: id, kind: "reload", err: err}
		}
		m.registerSnapshot(id, snapshot)
		return asyncMsg{id: id, kind: "reload", value: reloadPayload{snapshot: snapshot, collections: collections, apis: apis}}
	})
}

func (m *model) applyReload(payload reloadPayload) {
	d := m.docs
	if d.owned && d.snapshot != nil {
		_ = d.snapshot.Close()
	}
	d.snapshot, d.owned = payload.snapshot, true
	d.collections, d.apis = payload.collections, payload.apis
	d.frames, d.reader, d.importer = nil, documentReader{}, nil
	d.collectionCursor, d.catalogCursor, d.focus = 0, 0, 1
	m.stacks[contextDocumentation] = []screenID{screenDocumentation}
	if m.context == contextDocumentation {
		m.stack = []screenID{screenDocumentation}
	}
	if payload.imported != nil {
		m.status = fmt.Sprintf("Imported %s %s (%d pages) and reopened generation %s", payload.imported.Name, payload.imported.Version, payload.imported.Pages, d.snapshot.Fingerprint())
	} else {
		m.status = "Documentation rebuilt and reopened at generation " + d.snapshot.Fingerprint()
	}
	m.severity = severitySuccess
}

func (m *model) openImport() {
	if m.runtime == nil {
		m.status, m.severity = "Documentation import requires an open runtime", severityError
		return
	}
	values := []struct{ placeholder, value string }{
		{"markdown, openapi, llms, html, or docsify", "markdown"},
		{"API name (not needed for markdown)", ""},
		{"version (not needed for markdown)", ""},
		{"directory, file, or URL", ""},
		{"HTML max pages", strconv.Itoa(importer.DefaultMaxHTMLPages)},
		{"HTML max depth", strconv.Itoa(importer.DefaultMaxHTMLDepth)},
	}
	form := &importForm{focus: 0}
	for _, value := range values {
		form.inputs = append(form.inputs, newTextInput(value.placeholder, value.value))
	}
	m.docs.importer = form
	m.docs.resize(m.width, m.height)
	m.push(screenDocumentImport)
	_ = form.inputs[0].Focus()
}

func (m *model) updateImport(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	form := m.docs.importer
	switch key.String() {
	case "esc":
		m.docs.importer = nil
		m.pop()
		return m, nil
	case "tab", "shift+tab":
		if form.focus < len(form.inputs) {
			form.inputs[form.focus].Blur()
		}
		direction := 1
		if key.String() == "shift+tab" {
			direction = -1
		}
		form.focus = (form.focus + direction + len(form.inputs) + 1) % (len(form.inputs) + 1)
		if form.focus < len(form.inputs) {
			return m, form.inputs[form.focus].Focus()
		}
		return m, nil
	case "enter":
		if form.focus == len(form.inputs) {
			return m, m.importDocumentationCmd()
		}
		form.inputs[form.focus].Blur()
		form.focus = (form.focus + 1) % (len(form.inputs) + 1)
		if form.focus < len(form.inputs) {
			return m, form.inputs[form.focus].Focus()
		}
		return m, nil
	}
	if form.focus >= len(form.inputs) {
		return m, nil
	}
	var cmd tea.Cmd
	form.inputs[form.focus], cmd = form.inputs[form.focus].Update(key)
	return m, cmd
}

func (m *model) importDocumentationCmd() tea.Cmd {
	form := m.docs.importer
	values := make([]string, len(form.inputs))
	for index := range form.inputs {
		values[index] = strings.TrimSpace(form.inputs[index].Value())
	}
	kind, name, version, source := strings.ToLower(values[0]), values[1], values[2], values[3]
	maxPages, pagesErr := strconv.Atoi(values[4])
	maxDepth, depthErr := strconv.Atoi(values[5])
	if source == "" || pagesErr != nil || depthErr != nil {
		m.status, m.severity = "Source and valid numeric HTML limits are required", severityError
		return nil
	}
	if kind != "markdown" && (name == "" || version == "") {
		m.status, m.severity = "API name and version are required for this import type", severityError
		return nil
	}
	runtime := m.runtime
	return m.startOperation("import", "Importing documentation", false, func(ctx context.Context, id int64) tea.Msg {
		options := importer.Options{LibraryRoot: runtime.Paths.Library, Rebuild: runtime.RebuildLibrary}
		var result importer.Result
		var err error
		switch kind {
		case "markdown":
			result, err = importer.ImportMarkdown(ctx, source, options)
		case "openapi":
			result, err = importer.ImportOpenAPI(ctx, name, version, source, options)
		case "llms":
			result, err = importer.ImportLLMSTxt(ctx, name, version, source, options)
		case "html":
			options.MaxHTMLPages, options.MaxHTMLDepth = maxPages, maxDepth
			result, err = importer.ImportHTML(ctx, name, version, source, options)
		case "docsify":
			options.HTMLLimitsSet, options.MaxHTMLPages, options.MaxHTMLDepth = true, -1, -1
			result, err = importer.ImportDocsify(ctx, name, version, source, options)
		default:
			err = fmt.Errorf("unknown import type %q", kind)
		}
		if err != nil {
			return asyncMsg{id: id, kind: "import", err: err}
		}
		snapshot, err := library.Open(ctx, libraryOptions(runtime))
		if err != nil {
			return asyncMsg{id: id, kind: "import", err: err}
		}
		collections, apis, err := loadCatalog(ctx, snapshot)
		if err != nil {
			_ = snapshot.Close()
			return asyncMsg{id: id, kind: "import", err: err}
		}
		if err := ctx.Err(); err != nil {
			_ = snapshot.Close()
			return asyncMsg{id: id, kind: "import", err: err}
		}
		m.registerSnapshot(id, snapshot)
		return asyncMsg{id: id, kind: "import", value: reloadPayload{snapshot: snapshot, collections: collections, apis: apis, imported: &result}}
	})
}

func (m *model) viewDocumentation(width, height int) []string {
	switch m.currentScreen() {
	case screenDocumentBrowser:
		return m.viewDocumentBrowser(width, height)
	case screenDocumentReader:
		return m.viewDocumentReader(width, height)
	case screenDocumentImport:
		return m.viewDocumentImport(width, height)
	default:
		return m.viewDocumentCatalog(width, height)
	}
}

func (m *model) viewDocumentCatalog(width, height int) []string {
	d := m.docs
	collectionStart, collectionEnd := visibleWindow(d.collectionCursor, len(d.collections)+1, max(1, height-1))
	collections := make([]string, 0, collectionEnd-collectionStart)
	for index := collectionStart; index < collectionEnd; index++ {
		line := "  All documentation"
		if index > 0 {
			collection := d.collections[index-1]
			line = fmt.Sprintf("  %-18s %d", safeLine(collection.Name), collection.APICount)
		}
		collections = append(collections, selectedLine(line, max(1, width), d.focus == 0 && d.collectionCursor == index))
	}
	choices := d.choices()
	start, end := visibleWindow(d.catalogCursor, len(choices), max(1, height-2))
	filter := "  / " + d.filter.View()
	if !d.filtering {
		filter = styleDim.Render("  / filter: ") + d.filter.Value()
	}
	catalog := []string{filter}
	for index := start; index < end; index++ {
		choice := choices[index]
		line := fmt.Sprintf("  %-24s %-12s %4d pages", safeLine(choice.name), safeLine(choice.version.Version), choice.version.Pages)
		catalog = append(catalog, selectedLine(line, max(1, width), d.focus == 1 && index == d.catalogCursor))
	}
	preview := []string{styleDim.Render("  Select an API version to browse its paths and pages.")}
	if len(choices) > 0 {
		choice := choices[d.catalogCursor]
		preview = []string{
			"  " + styleTitle.Render(safeLine(choice.name)),
			fmt.Sprintf("  Version %s · %d pages", safeLine(choice.version.Version), choice.version.Pages),
			"  " + styleDim.Render(safeLine(strings.Join(choice.collections, ", "))),
			"",
		}
		preview = append(preview, wrapLines(safeMultiline(choice.description), max(10, width/3-4))...)
	}
	if m.width >= 120 {
		left := 25
		middle := min(55, (width-left-2)*3/5)
		right := width - left - middle - 2
		return joinPanes(height,
			pane("Collections", d.focus == 0, left, height, collections),
			pane("API / version catalog", d.focus == 1, middle, height, catalog),
			pane("Selected", false, right, height, preview),
		)
	}
	if m.width >= 80 {
		left := 28
		return joinPanes(height, pane("Collections", d.focus == 0, left, height, collections), pane("API / version catalog", d.focus == 1, width-left-1, height, catalog))
	}
	if d.focus == 0 {
		return pane("Collections", true, width, height, collections)
	}
	return pane("API / version catalog", true, width, height, catalog)
}

func (m *model) viewDocumentBrowser(width, height int) []string {
	frame := m.docs.currentFrame()
	navigation := []string{"  " + styleDim.Render(safeLine(m.documentationBreadcrumb()))}
	start, end := visibleWindow(frame.cursor, len(frame.entries), max(1, height-3))
	for index := start; index < end; index++ {
		entry := frame.entries[index]
		line := ""
		if entry.path != nil {
			line = fmt.Sprintf("  / %-30s %d pages", safeLine(strings.TrimPrefix(entry.path.Path, frame.path+"/")), entry.path.NestedPages)
		} else {
			line = "  · " + safeLine(entry.page.Title)
		}
		navigation = append(navigation, selectedLine(line, max(1, width), frame.focus == 0 && index == frame.cursor))
	}
	search := []string{"  " + frame.search.View()}
	if !frame.searching {
		search[0] = styleDim.Render("  s search: ") + frame.search.Value()
	}
	start, end = visibleWindow(frame.resultCursor, len(frame.results), max(1, height-3))
	for index := start; index < end; index++ {
		hit := frame.results[index]
		search = append(search, selectedLine(fmt.Sprintf("  %-28s :%d  %s", safeLine(hit.Title), hit.Line, safeLine(hit.Snippet)), max(1, width), frame.focus == 1 && index == frame.resultCursor))
	}
	detail := []string{"  " + styleDim.Render("Choose a path, page, or search result.")}
	if frame.focus == 0 && len(frame.entries) > 0 {
		entry := frame.entries[frame.cursor]
		if entry.path != nil {
			detail = []string{"  " + styleTitle.Render(safeLine(entry.path.Path)), fmt.Sprintf("  %d nested pages", entry.path.NestedPages)}
		} else {
			detail = []string{"  " + styleTitle.Render(safeLine(entry.page.Title)), "", "  " + styleDim.Render(safeLine(entry.page.Description))}
		}
	} else if frame.focus == 1 && len(frame.results) > 0 {
		hit := frame.results[frame.resultCursor]
		detail = []string{"  " + styleTitle.Render(safeLine(hit.Title)), fmt.Sprintf("  %s · line %d", safeLine(displayPath(hit.Path)), hit.Line), ""}
		detail = append(detail, wrapLines(safeMultiline(hit.Snippet), max(10, width/3-4))...)
	}
	if m.width >= 120 {
		left := 45
		middle := 45
		return joinPanes(height, pane("Paths and pages", frame.focus == 0, left, height, navigation), pane("Search", frame.focus == 1, middle, height, search), pane("Selected", false, width-left-middle-2, height, detail))
	}
	if m.width >= 80 {
		left := width / 2
		return joinPanes(height, pane("Paths and pages", frame.focus == 0, left, height, navigation), pane("Search", frame.focus == 1, width-left-1, height, search))
	}
	if frame.focus == 0 {
		return pane("Paths and pages", true, width, height, navigation)
	}
	return pane("Search", true, width, height, search)
}

func (m *model) viewDocumentReader(width, height int) []string {
	d := m.docs
	d.resizeReader(max(10, width-4), max(3, height-2))
	header := fmt.Sprintf("  %s  %s", styleTitle.Render(safeLine(d.reader.result.Title)), styleDim.Render(safeLine(m.documentationBreadcrumb())))
	rows := []string{fixed(header, width), fixed(fmt.Sprintf("  lines %d-%d of %d  ·  %3.0f%%", d.reader.result.Lines[0], d.reader.result.Lines[1], d.reader.result.TotalLines, d.reader.viewport.ScrollPercent()*100), width)}
	rows = append(rows, strings.Split(d.reader.viewport.View(), "\n")...)
	for len(rows) < height {
		rows = append(rows, "")
	}
	return rows[:height]
}

func (m *model) viewDocumentImport(width, height int) []string {
	form := m.docs.importer
	labels := []string{"Import type", "API name", "Version", "Source", "HTML max pages", "HTML max depth"}
	rows := []string{"  " + styleTitle.Render("Import documentation"), "  " + styleDim.Render("The imported source is validated, indexed, and reopened in this application."), ""}
	for index, input := range form.inputs {
		line := fmt.Sprintf("  %-18s %s", labels[index], input.View())
		rows = append(rows, selectedLine(line, width, form.focus == index))
	}
	rows = append(rows, "", selectedLine("  Import and reopen library", width, form.focus == len(form.inputs)))
	for len(rows) < height {
		rows = append(rows, "")
	}
	return rows[:height]
}

func (m *model) documentationBreadcrumb() string {
	parts := []string{"Documentation"}
	if m.docs.selected.name != "" {
		parts = append(parts, m.docs.selected.name+" "+m.docs.selected.version.Version)
	}
	for _, frame := range m.docs.frames {
		if frame.path != "" {
			parts = append(parts, frame.path)
		}
	}
	if m.currentScreen() == screenDocumentReader {
		parts = append(parts, m.docs.reader.result.Title)
	}
	return strings.Join(parts, " / ")
}

func (m *model) documentationFooter() string {
	if m.docs.filtering {
		return "type filter  enter finish  esc clear focus; then 1-4/?/q"
	}
	if len(m.docs.frames) > 0 && m.docs.currentFrame().searching {
		return "type search  enter run  esc stop editing; then 1-4/?/q"
	}
	switch m.currentScreen() {
	case screenDocumentReader:
		return "j/k or arrows scroll  pgup/pgdn  esc back  ? help"
	case screenDocumentImport:
		return "tab fields  type to edit  enter advance/run  esc back"
	case screenDocumentBrowser:
		return "tab panes  j/k navigate  enter open  s search  i import  r rebuild/reopen  esc back"
	default:
		return "tab panes  j/k navigate  / filter  enter browse  i import  r rebuild/reopen  1-4 contexts  ? help  q quit"
	}
}

func containsFold(values []string, wanted string) bool {
	for _, value := range values {
		if strings.EqualFold(value, wanted) {
			return true
		}
	}
	return false
}
