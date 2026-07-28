package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sairaph/apis-mcp/internal/sessions"
)

type sessionState struct {
	items        []sessions.SessionInfo
	cursor       int
	inspection   sessions.Inspection
	cookieCursor int
}

type sessionActionPayload struct {
	message     string
	resetDetail bool
	items       []sessions.SessionInfo
	warning     error
}

type sessionInspectionPayload struct {
	inspection sessions.Inspection
	items      []sessions.SessionInfo
	warning    error
}

func newSessionState() *sessionState { return &sessionState{} }

func (s *sessionState) clamp() {
	if len(s.items) == 0 {
		s.cursor = 0
		return
	}
	s.cursor = max(0, min(s.cursor, len(s.items)-1))
}

func (s *sessionState) setItems(items []sessions.SessionInfo) {
	selectedID := s.selected().ID
	s.items = items
	s.inspection = sessions.Inspection{}
	s.cookieCursor = 0
	s.cursor = 0
	for index, item := range items {
		if item.ID == selectedID {
			s.cursor = index
			break
		}
	}
	s.clamp()
}

func (s *sessionState) setInspection(inspection sessions.Inspection) {
	s.inspection = inspection
	if len(inspection.Cookies) == 0 {
		s.cookieCursor = 0
	} else {
		s.cookieCursor = max(0, min(s.cookieCursor, len(inspection.Cookies)-1))
	}
}

func (s *sessionState) selected() sessions.SessionInfo {
	if len(s.items) == 0 || s.cursor >= len(s.items) {
		return sessions.SessionInfo{}
	}
	return s.items[s.cursor]
}

func (m *model) updateSessions(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	s := m.sessions
	if m.currentScreen() == screenSessionDetail {
		switch key.String() {
		case "esc", "backspace":
			m.pop()
		case "up", "k":
			if count := len(s.inspection.Cookies); count > 0 {
				s.cookieCursor = (s.cookieCursor - 1 + count) % count
			}
		case "down", "j":
			if count := len(s.inspection.Cookies); count > 0 {
				s.cookieCursor = (s.cookieCursor + 1) % count
			}
		case "x", "delete":
			m.confirmCookieDelete()
		case "d":
			m.confirmSessionDelete(s.inspection.Session.ID)
		case "s":
			m.selectRequestSession(s.inspection.Session.ID)
		case "r":
			return m, m.inspectSessionCmd(s.inspection.Session.ID)
		}
		return m, nil
	}
	switch key.String() {
	case "up", "k":
		if len(s.items) > 0 {
			s.cursor = (s.cursor - 1 + len(s.items)) % len(s.items)
		}
	case "down", "j":
		if len(s.items) > 0 {
			s.cursor = (s.cursor + 1) % len(s.items)
		}
	case "enter":
		if selected := s.selected(); selected.ID != "" {
			return m, m.inspectSessionCmd(selected.ID)
		}
	case "n":
		return m, m.createSessionCmd()
	case "s", " ":
		m.selectRequestSession(s.selected().ID)
	case "d", "delete", "backspace":
		m.confirmSessionDelete(s.selected().ID)
	case "c":
		m.confirm("Clean up expired sessions?", []string{"Only sessions older than the configured retention period are removed."}, "Clean up", func() tea.Cmd { return m.cleanupSessionsCmd() })
	case "r":
		return m, m.loadSessionsCmd()
	}
	return m, nil
}

func (m *model) inspectSessionCmd(id string) tea.Cmd {
	if m.runtime == nil || m.runtime.Sessions == nil || id == "" {
		return nil
	}
	manager := m.runtime.Sessions
	return m.startOperation("session-inspect", "Loading session cookies", false, func(_ context.Context, operationID int64) tea.Msg {
		inspection, err := manager.Inspect(id)
		if err != nil {
			return asyncMsg{id: operationID, kind: "session-inspect", err: err}
		}
		items, listErr := manager.List()
		return asyncMsg{id: operationID, kind: "session-inspect", value: sessionInspectionPayload{inspection: inspection, items: items, warning: listErr}}
	})
}

func (m *model) createSessionCmd() tea.Cmd {
	if m.runtime == nil || m.runtime.Sessions == nil {
		return nil
	}
	manager := m.runtime.Sessions
	return m.startOperation("session-action", "Creating cookie session", false, func(ctx context.Context, id int64) tea.Msg {
		handle, err := manager.Create(ctx)
		if err != nil {
			return asyncMsg{id: id, kind: "session-action", err: err}
		}
		sessionID := handle.ID()
		err = handle.Close()
		if err != nil {
			return asyncMsg{id: id, kind: "session-action", err: err}
		}
		items, listErr := manager.List()
		return asyncMsg{id: id, kind: "session-action", value: sessionActionPayload{message: "Created session " + sessionID, items: items, warning: listErr}}
	})
}

func (m *model) confirmSessionDelete(id string) {
	if id == "" || m.runtime == nil || m.runtime.Sessions == nil {
		return
	}
	m.confirm("Delete cookie session?", []string{id, "All persisted cookies in this session are removed. This cannot be undone."}, "Delete session", func() tea.Cmd {
		manager := m.runtime.Sessions
		return m.startOperation("session-action", "Deleting cookie session", false, func(_ context.Context, operationID int64) tea.Msg {
			err := manager.Delete(id)
			if err != nil {
				return asyncMsg{id: operationID, kind: "session-action", err: err}
			}
			items, listErr := manager.List()
			return asyncMsg{id: operationID, kind: "session-action", value: sessionActionPayload{message: "Deleted session " + id, resetDetail: true, items: items, warning: listErr}}
		})
	})
}

func (m *model) confirmCookieDelete() {
	s := m.sessions
	if m.runtime == nil || m.runtime.Sessions == nil || len(s.inspection.Cookies) == 0 || s.cookieCursor >= len(s.inspection.Cookies) {
		return
	}
	cookie := s.inspection.Cookies[s.cookieCursor]
	sessionID := s.inspection.Session.ID
	m.confirm("Delete cookie?", []string{
		fmt.Sprintf("%s for %s%s", cookie.Name, cookie.Domain, cookie.Path),
		"The cookie is removed from this persisted session only.",
	}, "Delete cookie", func() tea.Cmd {
		manager := m.runtime.Sessions
		return m.startOperation("session-inspect", "Deleting cookie", false, func(ctx context.Context, operationID int64) tea.Msg {
			if err := manager.DeleteCookie(ctx, sessionID, cookie.Domain, cookie.Path, cookie.Name); err != nil {
				return asyncMsg{id: operationID, kind: "session-inspect", err: err}
			}
			inspection, err := manager.Inspect(sessionID)
			if err != nil {
				return asyncMsg{id: operationID, kind: "session-inspect", err: err}
			}
			items, listErr := manager.List()
			return asyncMsg{id: operationID, kind: "session-inspect", value: sessionInspectionPayload{inspection: inspection, items: items, warning: listErr}}
		})
	})
}

func (m *model) cleanupSessionsCmd() tea.Cmd {
	if m.runtime == nil || m.runtime.Sessions == nil {
		return nil
	}
	manager := m.runtime.Sessions
	return m.startOperation("session-action", "Cleaning expired sessions", false, func(_ context.Context, id int64) tea.Msg {
		result, err := manager.Cleanup()
		if err != nil {
			return asyncMsg{id: id, kind: "session-action", err: err}
		}
		items, listErr := manager.List()
		return asyncMsg{id: id, kind: "session-action", value: sessionActionPayload{message: fmt.Sprintf("Removed %d expired sessions", result.Removed), items: items, warning: listErr}}
	})
}

func (m *model) selectRequestSession(id string) {
	if id == "" {
		return
	}
	for index, session := range m.request.sessions {
		if session.ID == id {
			m.request.sessionCursor = index + 1
			m.status, m.severity = "Selected "+id+" for the Request workspace", severitySuccess
			return
		}
	}
	m.status, m.severity = "Session is no longer available", severityWarning
}

func (m *model) viewSessions(width, height int) []string {
	s := m.sessions
	list := s.sessionRows(height - 1)
	summary := s.summaryRows()
	cookies := s.cookieRows(height - 1)
	if m.width >= 120 {
		left, middle := 42, 32
		return joinPanes(height,
			pane("Cookie sessions", m.currentScreen() == screenSessionList, left, height, list),
			pane("Selected session", false, middle, height, summary),
			pane("Cookies by domain", m.currentScreen() == screenSessionDetail, width-left-middle-2, height, cookies),
		)
	}
	if m.width >= 80 {
		left := width / 2
		cookies = s.cookieRows(max(1, height-len(summary)-1))
		return joinPanes(height,
			pane("Cookie sessions", m.currentScreen() == screenSessionList, left, height, list),
			pane("Session detail", m.currentScreen() == screenSessionDetail, width-left-1, height, append(summary, cookies...)),
		)
	}
	if m.currentScreen() == screenSessionDetail {
		cookies = s.cookieRows(max(1, height-len(summary)-1))
		return pane("Session detail", true, width, height, append(summary, cookies...))
	}
	return pane("Cookie sessions", true, width, height, list)
}

func (s *sessionState) sessionRows(height int) []string {
	if len(s.items) == 0 {
		return []string{"  No sessions yet.", "  Press n to create one, or send a request."}
	}
	start, end := visibleWindow(s.cursor, len(s.items), height)
	rows := make([]string, 0, end-start)
	for index := start; index < end; index++ {
		item := s.items[index]
		domains := strings.Join(item.Domains, ", ")
		if domains == "" {
			domains = "no domains"
		}
		line := fmt.Sprintf("  %-10s %3d cookies  %-16s %s", shortID(item.ID), item.CookieCount, relativeSessionTime(item.LastUsedAt), safeLine(domains))
		rows = append(rows, selectedLine(line, 120, index == s.cursor))
	}
	return rows
}

func (s *sessionState) summaryRows() []string {
	item := s.selected()
	if s.inspection.Session.ID != "" && (item.ID == "" || s.inspection.Session.ID == item.ID) {
		item = s.inspection.Session
	}
	if item.ID == "" {
		return []string{"  Select a session to inspect its cookies."}
	}
	return []string{
		"  " + styleTitle.Render(shortID(item.ID)),
		"  ID " + safeLine(item.ID),
		fmt.Sprintf("  %d active cookies", item.CookieCount),
		"  created " + item.CreatedAt.UTC().Format(time.RFC3339),
		"  used " + item.LastUsedAt.UTC().Format(time.RFC3339),
		"", "  " + styleDim.Render(safeLine(strings.Join(item.Domains, ", "))),
	}
}

func (s *sessionState) cookieRows(height int) []string {
	if s.inspection.Session.ID == "" || s.selected().ID != s.inspection.Session.ID {
		return []string{"  Press enter to load typed cookie details."}
	}
	if len(s.inspection.Cookies) == 0 {
		return []string{"  This session has no active cookies."}
	}
	start, end := visibleWindow(s.cookieCursor, len(s.inspection.Cookies), height)
	rows := make([]string, 0, end-start+3)
	lastDomain := ""
	for index := start; index < end; index++ {
		cookie := s.inspection.Cookies[index]
		if cookie.Domain != lastDomain {
			rows = append(rows, "  "+styleTitle.Render(safeLine(cookie.Domain)))
			lastDomain = cookie.Domain
		}
		flags := ""
		if cookie.Secure {
			flags += " secure"
		}
		if cookie.HTTPOnly {
			flags += " http-only"
		}
		line := fmt.Sprintf("    %-22s %-12s%s", safeLine(cookie.Name), safeLine(cookie.Path), flags)
		rows = append(rows, selectedLine(line, 120, index == s.cookieCursor))
	}
	return rows
}

func (m *model) sessionsFooter() string {
	if m.currentScreen() == screenSessionDetail {
		return "j/k cookies  x delete cookie  d delete session  s use in Request  r refresh  esc back  ? help"
	}
	return "j/k navigate  enter inspect  n create  s use in Request  d delete  c cleanup  r refresh  1-4 contexts  q quit"
}

func relativeSessionTime(value time.Time) string {
	if value.IsZero() {
		return "never"
	}
	elapsed := time.Since(value)
	switch {
	case elapsed < time.Minute:
		return "just now"
	case elapsed < time.Hour:
		return fmt.Sprintf("%dm ago", int(elapsed.Minutes()))
	case elapsed < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(elapsed.Hours()))
	default:
		return fmt.Sprintf("%dd ago", int(elapsed.Hours()/24))
	}
}
