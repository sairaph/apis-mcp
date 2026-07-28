package cli

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/httpcall"
	"github.com/sairaph/apis-mcp/internal/sessions"
)

const requestFieldCount = 11

type requestState struct {
	method   textinput.Model
	url      textinput.Model
	headers  textarea.Model
	body     textarea.Model
	timeout  textinput.Model
	retries  textinput.Model
	jsonPath textinput.Model

	sessions              []sessions.SessionInfo
	sessionCursor         int
	allowLarge            bool
	largeDownloadsAllowed bool
	focus                 int
	formOffset            int

	result        *httpcall.Result
	resultErr     error
	response      viewport.Model
	renderedWidth int
}

func (r *requestState) setLargeDownloadPolicy(allowed bool) {
	r.largeDownloadsAllowed = allowed
	if !allowed {
		r.allowLarge = false
	}
}

type requestPayload struct {
	result      httpcall.Result
	sessions    []sessions.SessionInfo
	sessionsErr error
	err         error
}

func newRequestState() *requestState {
	method := newTextInput("GET", "GET")
	url := newTextInput("https://api.example.com/resource", "")
	timeout := newTextInput("30", "30")
	retries := newTextInput("default", "")
	jsonPath := newTextInput("optional, for example $.data", "")
	headers := textarea.New()
	headers.Prompt = ""
	headers.Placeholder = "Accept: application/json\nAuthorization: Bearer ..."
	headers.ShowLineNumbers = false
	headers.FocusedStyle.CursorLine = headers.FocusedStyle.Base
	body := textarea.New()
	body.Prompt = ""
	body.Placeholder = "JSON value, text/file path, or blank"
	body.ShowLineNumbers = false
	body.FocusedStyle.CursorLine = body.FocusedStyle.Base
	response := viewport.New(30, 10)
	response.MouseWheelEnabled = false
	state := &requestState{
		method: method, url: url, headers: headers, body: body,
		timeout: timeout, retries: retries, jsonPath: jsonPath,
		focus: 0, response: response,
	}
	state.focusField(0)
	return state
}

func (r *requestState) editing() bool {
	switch r.focus {
	case 0, 1, 3, 4, 5, 6, 7:
		return true
	default:
		return false
	}
}

func (r *requestState) blur() {
	r.method.Blur()
	r.url.Blur()
	r.headers.Blur()
	r.body.Blur()
	r.timeout.Blur()
	r.retries.Blur()
	r.jsonPath.Blur()
}

func (r *requestState) focusField(index int) tea.Cmd {
	r.blur()
	r.focus = max(-1, min(index, requestFieldCount-1))
	switch r.focus {
	case 0:
		return r.method.Focus()
	case 1:
		return r.url.Focus()
	case 3:
		return r.headers.Focus()
	case 4:
		return r.body.Focus()
	case 5:
		return r.timeout.Focus()
	case 6:
		return r.retries.Focus()
	case 7:
		return r.jsonPath.Focus()
	}
	return nil
}

func (r *requestState) resize(width, height int) {
	fieldWidth := max(12, width/2-20)
	r.method.Width = 10
	r.url.Width = fieldWidth
	r.timeout.Width = 8
	r.retries.Width = 8
	r.jsonPath.Width = fieldWidth
	r.headers.SetWidth(fieldWidth)
	r.headers.SetHeight(3)
	r.body.SetWidth(fieldWidth)
	r.body.SetHeight(4)
	r.response.Width = max(10, width/2-4)
	r.response.Height = max(3, height-9)
	if r.result != nil {
		r.setResponseContent(r.response.Width)
	}
}

func (r *requestState) setSessions(items []sessions.SessionInfo, selectedID string) {
	if selectedID == "" && r.sessionCursor > 0 && r.sessionCursor <= len(r.sessions) {
		selectedID = r.sessions[r.sessionCursor-1].ID
	}
	r.sessions = items
	r.sessionCursor = 0
	for index, item := range items {
		if item.ID == selectedID {
			r.sessionCursor = index + 1
			break
		}
	}
}

func (m *model) updateRequest(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	r := m.request
	switch key.String() {
	case "esc":
		if r.focus >= 0 {
			r.blur()
			r.focus = -1
			m.status, m.severity = "Request workspace unfocused; tab resumes editing", severityInfo
		}
		return m, nil
	case "tab", "shift+tab":
		direction := 1
		if key.String() == "shift+tab" {
			direction = -1
		}
		next := r.focus + direction
		if r.focus < 0 {
			next = 0
		}
		next = (next + requestFieldCount) % requestFieldCount
		return m, r.focusField(next)
	case "ctrl+enter", "ctrl+s":
		return m, m.prepareRequest()
	}
	if r.focus == 2 {
		count := len(r.sessions) + 1
		switch key.String() {
		case "up", "left", "k", "h":
			r.sessionCursor = (r.sessionCursor - 1 + count) % count
		case "down", "right", "j", "l", "enter", " ":
			r.sessionCursor = (r.sessionCursor + 1) % count
		}
		return m, nil
	}
	if r.focus == 8 {
		if !r.largeDownloadsAllowed {
			m.status, m.severity = "Large-download override is disabled by application policy", severityWarning
			return m, nil
		}
		if key.String() == "enter" || key.String() == " " || key.String() == "left" || key.String() == "right" {
			r.allowLarge = !r.allowLarge
		}
		return m, nil
	}
	if r.focus == 9 {
		if key.String() == "enter" || key.String() == " " {
			return m, m.prepareRequest()
		}
		return m, nil
	}
	if r.focus == 10 {
		var cmd tea.Cmd
		r.response, cmd = r.response.Update(key)
		return m, cmd
	}
	if key.String() == "enter" && r.focus != 3 && r.focus != 4 {
		return m, r.focusField((r.focus + 1) % requestFieldCount)
	}
	var cmd tea.Cmd
	switch r.focus {
	case 0:
		r.method, cmd = r.method.Update(key)
	case 1:
		r.url, cmd = r.url.Update(key)
	case 3:
		r.headers, cmd = r.headers.Update(key)
	case 4:
		r.body, cmd = r.body.Update(key)
	case 5:
		r.timeout, cmd = r.timeout.Update(key)
	case 6:
		r.retries, cmd = r.retries.Update(key)
	case 7:
		r.jsonPath, cmd = r.jsonPath.Update(key)
	}
	return m, cmd
}

func (m *model) prepareRequest() tea.Cmd {
	input, err := m.request.input()
	if err != nil {
		m.status, m.severity = err.Error(), severityError
		return nil
	}
	if mutationMethod(input.Method) {
		presence := fmt.Sprintf("headers: %s · body: %s · session: %s", present(input.Headers), present(input.Payload), present(input.Session))
		m.confirm(
			fmt.Sprintf("Send mutating %s request?", safeLine(input.Method)),
			[]string{safeLine(input.Endpoint), presence, "The remote service may change state. Secret header, body, and cookie values are not shown."},
			"Send request",
			func() tea.Cmd { return m.sendRequestCmd(input) },
		)
		return nil
	}
	return m.sendRequestCmd(input)
}

func present(value any) string {
	switch value := value.(type) {
	case nil:
		return "no"
	case string:
		if value == "" {
			return "no"
		}
	}
	return "yes"
}

func (r *requestState) input() (app.CallInput, error) {
	method := strings.ToUpper(strings.TrimSpace(r.method.Value()))
	endpoint := strings.TrimSpace(r.url.Value())
	if method == "" || endpoint == "" {
		return app.CallInput{}, errors.New("method and URL are required")
	}
	headers, err := parseHeaders(r.headers.Value())
	if err != nil {
		return app.CallInput{}, fmt.Errorf("headers: %w", err)
	}
	payload, err := jsonOrPath(r.body.Value())
	if err != nil {
		return app.CallInput{}, fmt.Errorf("body: %w", err)
	}
	timeout, err := strconv.Atoi(strings.TrimSpace(r.timeout.Value()))
	if err != nil || timeout < 0 {
		return app.CallInput{}, errors.New("timeout must be a non-negative number of seconds")
	}
	var retries *int
	if raw := strings.TrimSpace(r.retries.Value()); raw != "" {
		value, parseErr := strconv.Atoi(raw)
		if parseErr != nil || value < 0 {
			return app.CallInput{}, errors.New("retries must be blank or a non-negative integer")
		}
		retries = &value
	}
	session := ""
	if r.sessionCursor > 0 && r.sessionCursor <= len(r.sessions) {
		session = r.sessions[r.sessionCursor-1].ID
	}
	return app.CallInput{
		Method: method, Endpoint: endpoint, Headers: headers, Payload: payload,
		Timeout: timeout, Retries: retries, JSONPath: strings.TrimSpace(r.jsonPath.Value()),
		Session: session, AllowLargeDownload: r.largeDownloadsAllowed && r.allowLarge,
	}, nil
}

func parseHeaders(raw string) (any, error) {
	raw = strings.TrimSpace(raw)
	drivePrefix := len(raw) >= 2 && ((raw[0] >= 'A' && raw[0] <= 'Z') || (raw[0] >= 'a' && raw[0] <= 'z')) && raw[1] == ':'
	drivePath := drivePrefix && len(raw) >= 3 && raw[2] != ' ' && raw[2] != '\t' && (raw[2] == '.' || strings.ContainsAny(raw[2:], `\\/`) || strings.HasSuffix(strings.ToLower(raw), ".json"))
	windowsPath := strings.HasPrefix(raw, `\\`) || drivePath
	if raw == "" || strings.HasPrefix(raw, "{") || strings.HasPrefix(raw, "[") || !strings.Contains(raw, ":") || windowsPath {
		return jsonOrPath(raw)
	}
	headers := make(http.Header)
	for lineNumber, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		name, value, found := strings.Cut(line, ":")
		name, value = strings.TrimSpace(name), strings.TrimSpace(value)
		if !found || name == "" {
			return nil, fmt.Errorf("line %d must be Name: value", lineNumber+1)
		}
		headers.Add(name, value)
	}
	return headers, nil
}

func mutationMethod(method string) bool {
	switch strings.ToUpper(method) {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
		return false
	default:
		return true
	}
}

func (m *model) sendRequestCmd(input app.CallInput) tea.Cmd {
	if m.runtime == nil || m.runtime.HTTP == nil || m.runtime.Sessions == nil {
		m.status, m.severity = "HTTP workspace is unavailable", severityError
		return nil
	}
	service := m.runtime.HTTP
	manager := m.runtime.Sessions
	return m.startOperation("request", "Sending "+input.Method+" request", true, func(ctx context.Context, id int64) tea.Msg {
		result, err := service.Call(ctx, input)
		items, sessionsErr := manager.List()
		return asyncMsg{id: id, kind: "request", value: requestPayload{result: result, sessions: items, sessionsErr: sessionsErr, err: err}}
	})
}

func (r *requestState) applyResult(result httpcall.Result) {
	r.applyOutcome(result, nil)
}

func (r *requestState) applyOutcome(result httpcall.Result, err error) {
	if err != nil && result.Request.ID == "" {
		r.result = nil
	} else {
		r.result = &result
	}
	r.resultErr = err
	r.renderedWidth = 0
	r.blur()
	r.focus = 10
	r.response.GotoTop()
	r.setResponseContent(max(20, r.response.Width))
}

func (r *requestState) resultStatus(err error) (string, severity) {
	if err != nil {
		return err.Error(), severityError
	}
	if r.result == nil {
		return "Request completed", severitySuccess
	}
	status := r.result.Response.Status
	message := fmt.Sprintf("HTTP %d · %s · %d decoded bytes", status, r.result.Response.State, r.result.Response.DecodedBytes)
	switch {
	case status >= 500:
		return message, severityError
	case status >= 400:
		return message, severityWarning
	default:
		return message, severitySuccess
	}
}

func (r *requestState) setResponseContent(width int) {
	if r.renderedWidth == width {
		return
	}
	r.renderedWidth = width
	if r.result == nil {
		if r.resultErr != nil {
			r.response.SetContent(styleError.Render("Request failed") + "\n" + safeMultiline(r.resultErr.Error()))
		} else {
			r.response.SetContent("No response yet. Fill the workspace and send a request.")
		}
		return
	}
	result := r.result
	lines := make([]string, 0, 24)
	if r.resultErr != nil {
		lines = append(lines, styleError.Render("Request failed"), safeMultiline(r.resultErr.Error()), "")
	}
	lines = append(lines,
		styleTitle.Render(fmt.Sprintf("HTTP %d", result.Response.Status))+"  "+safeLine(result.Response.StatusText),
		fmt.Sprintf("%s %s", safeLine(result.Request.Method), safeLine(result.Request.Endpoint)),
		"request ID: "+safeLine(result.Request.ID),
		fmt.Sprintf("state: %s  duration: %s  attempts: %d", safeLine(result.Response.State), result.Response.Duration, len(result.Attempts)),
		fmt.Sprintf("session: %s", safeLine(result.Request.SessionID)),
		fmt.Sprintf("decoded: %d bytes  wire: %d bytes  type: %s", result.Response.DecodedBytes, result.Response.WireBytes, safeLine(result.Response.ContentType)),
	)
	if result.Response.FinalURL != "" && result.Response.FinalURL != result.Request.Endpoint {
		lines = append(lines, "final URL: "+safeLine(result.Response.FinalURL))
	}
	if len(result.Request.AutomaticHeaders) > 0 {
		lines = append(lines, "", styleTitle.Render("Automatic headers"))
		names := make([]string, 0, len(result.Request.AutomaticHeaders))
		for name := range result.Request.AutomaticHeaders {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			lines = append(lines, safeLine(name)+": "+safeLine(result.Request.AutomaticHeaders[name]))
		}
	}
	if len(result.Response.Headers) > 0 {
		lines = append(lines, "", styleTitle.Render("Response headers"))
		names := make([]string, 0, len(result.Response.Headers))
		for name := range result.Response.Headers {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			values := make([]string, len(result.Response.Headers[name]))
			for index, value := range result.Response.Headers[name] {
				values[index] = safeLine(value)
			}
			lines = append(lines, safeLine(name)+": "+strings.Join(values, ", "))
		}
	}
	if len(result.Attempts) > 0 {
		lines = append(lines, "", styleTitle.Render("Attempts"))
		for _, attempt := range result.Attempts {
			line := fmt.Sprintf("#%d status=%d", attempt.Number, attempt.Status)
			if attempt.Error != "" {
				line += " error=" + safeLine(attempt.Error)
			}
			if attempt.RetryReason != "" {
				line += " retry=" + safeLine(attempt.RetryReason) + " after " + attempt.RetryDelay.String()
			}
			lines = append(lines, line)
		}
	}
	if len(result.Redirects) > 0 {
		lines = append(lines, "", styleTitle.Render("Redirects"))
		for _, redirect := range result.Redirects {
			lines = append(lines, fmt.Sprintf("HTTP %d %s → %s  %s→%s body=%t", redirect.Status, safeLine(redirect.FromURL), safeLine(redirect.ToURL), safeLine(redirect.MethodBefore), safeLine(redirect.MethodAfter), redirect.BodyRetained))
		}
	}
	cachePaths := []struct{ label, value string }{
		{"directory", result.Cache.Directory}, {"body", result.Cache.BodyPath}, {"temporary", result.Cache.TempPath},
		{"final", result.Cache.FinalPath}, {"headers", result.Cache.HeadersPath}, {"metadata", result.Cache.MetadataPath}, {"error", result.Cache.ErrorPath},
	}
	lines = append(lines, "", styleTitle.Render("Cache paths"))
	for _, item := range cachePaths {
		if item.value != "" {
			lines = append(lines, item.label+": "+safeLine(item.value))
		}
	}
	if result.Selection != nil {
		selection := fmt.Sprintf("JSONPath %s: matched=%t", safeLine(result.Selection.JSONPath), result.Selection.Matched)
		if result.Selection.Error != "" {
			selection += " (" + safeLine(result.Selection.Error) + ")"
		}
		lines = append(lines, "", selection)
	}
	if result.Preview != nil {
		lines = append(lines, "", styleTitle.Render("Preview")+"  "+styleDim.Render(safeLine(result.Preview.Kind)))
		lines = append(lines, safeMultiline(result.Preview.Content))
		if result.Preview.Truncated {
			lines = append(lines, styleWarning.Render("Preview truncated; use the cached body path for the complete response."))
		}
	} else {
		lines = append(lines, "", styleDim.Render("No textual preview is available; the cached body path is authoritative."))
	}
	var wrapped []string
	for _, line := range lines {
		wrapped = append(wrapped, wrapLines(line, max(10, width))...)
	}
	r.response.SetContent(strings.Join(wrapped, "\n"))
}

func (m *model) viewRequest(width, height int) []string {
	r := m.request
	if m.width >= 120 {
		basicWidth := 37
		payloadWidth := 43
		responseWidth := width - basicWidth - payloadWidth - 2
		r.resizeRequestEditors(payloadWidth - 4)
		r.response.Width, r.response.Height = responseWidth, max(3, height-1)
		r.setResponseContent(max(10, responseWidth-2))
		return joinPanes(height,
			pane("Request", r.focus >= 0 && r.focus <= 2 || r.focus >= 5 && r.focus <= 9, basicWidth, height, r.basicRows(basicWidth)),
			pane("Headers and body", r.focus == 3 || r.focus == 4, payloadWidth, height, r.payloadRows(payloadWidth)),
			pane("Response", r.focus == 10, responseWidth, height, strings.Split(r.response.View(), "\n")),
		)
	}
	if m.width >= 80 {
		formWidth := width/2 + 5
		responseWidth := width - formWidth - 1
		r.resizeRequestEditors(formWidth - 20)
		r.response.Width, r.response.Height = responseWidth, max(3, height-1)
		r.setResponseContent(max(10, responseWidth-2))
		return joinPanes(height,
			pane("Request workspace", r.focus != 10, formWidth, height, r.scrollingFormRows(formWidth, height-1)),
			pane("Response", r.focus == 10, responseWidth, height, strings.Split(r.response.View(), "\n")),
		)
	}
	if r.focus == 10 {
		r.response.Width, r.response.Height = width, max(3, height-1)
		r.setResponseContent(max(10, width-2))
		return pane("Response", true, width, height, strings.Split(r.response.View(), "\n"))
	}
	r.resizeRequestEditors(width - 20)
	return pane("Request workspace", true, width, height, r.scrollingFormRows(width, height-1))
}

func (r *requestState) resizeRequestEditors(width int) {
	width = max(12, width)
	r.url.Width, r.jsonPath.Width = width, width
	r.headers.SetWidth(width)
	r.headers.SetHeight(3)
	r.body.SetWidth(width)
	r.body.SetHeight(4)
}

func focusLabel(label string, focused bool) string {
	if focused {
		return styleTitle.Render("› " + label)
	}
	return styleDim.Render("  " + label)
}

func (r *requestState) basicRows(width int) []string {
	return []string{
		focusLabel("Method", r.focus == 0) + "  " + r.method.View(),
		focusLabel("URL", r.focus == 1), "  " + r.url.View(),
		focusLabel("Cookie session", r.focus == 2), "  " + r.sessionName(),
		"", styleDim.Render("  Advanced"),
		focusLabel("Header timeout", r.focus == 5) + "  " + r.timeout.View() + " sec",
		focusLabel("Retries", r.focus == 6) + "  " + r.retries.View(),
		focusLabel("JSONPath", r.focus == 7), "  " + r.jsonPath.View(),
		r.largeDownloadRow(),
		"", selectedLine("  Send request", max(1, width), r.focus == 9),
	}
}

func (r *requestState) largeDownloadRow() string {
	if !r.largeDownloadsAllowed {
		return styleDim.Render("  Large download override  disabled by policy")
	}
	return focusLabel("Large download", r.focus == 8) + fmt.Sprintf("  [%s]", checkMark(r.allowLarge))
}

func (r *requestState) payloadRows(_ int) []string {
	rows := []string{focusLabel("Headers", r.focus == 3)}
	rows = append(rows, strings.Split(r.headers.View(), "\n")...)
	rows = append(rows, focusLabel("Body", r.focus == 4))
	rows = append(rows, strings.Split(r.body.View(), "\n")...)
	return rows
}

func (r *requestState) scrollingFormRows(width, height int) []string {
	rows := append(r.basicRows(width), "", focusLabel("Headers", r.focus == 3))
	headerStart := len(rows) - 1
	rows = append(rows, strings.Split(r.headers.View(), "\n")...)
	rows = append(rows, focusLabel("Body", r.focus == 4))
	bodyStart := len(rows) - 1
	rows = append(rows, strings.Split(r.body.View(), "\n")...)
	focusRow := map[int]int{0: 0, 1: 1, 2: 3, 3: headerStart, 4: bodyStart, 5: 7, 6: 8, 7: 9, 8: 11, 9: 13}[r.focus]
	if focusRow < r.formOffset {
		r.formOffset = focusRow
	}
	if focusRow >= r.formOffset+height {
		r.formOffset = focusRow - height + 1
	}
	r.formOffset = max(0, min(r.formOffset, max(0, len(rows)-height)))
	end := min(len(rows), r.formOffset+height)
	return rows[r.formOffset:end]
}

func (r *requestState) sessionName() string {
	if r.sessionCursor <= 0 || r.sessionCursor > len(r.sessions) {
		return "New session (cookies persist after response)"
	}
	session := r.sessions[r.sessionCursor-1]
	domains := safeLine(strings.Join(session.Domains, ", "))
	if domains == "" {
		domains = "no cookies"
	}
	return fmt.Sprintf("%s · %s", shortID(session.ID), domains)
}

func checkMark(value bool) string {
	if value {
		return "x"
	}
	return " "
}

func (m *model) requestFooter() string {
	if m.request.focus == 10 {
		return "j/k or arrows scroll response  shift+tab return  esc unfocus; then 1-4/?/q"
	}
	return "tab fields  enter advance  ctrl+enter send  esc unfocus; then 1-4/?/q  ctrl+c cancel"
}

func shortID(id string) string {
	id = safeLine(id)
	runes := []rune(id)
	if len(runes) <= 12 {
		return id
	}
	return string(runes[:8]) + "…"
}
