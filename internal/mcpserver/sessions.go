package mcpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sairaph/apis-mcp/internal/app"
	"github.com/sairaph/apis-mcp/internal/bootstrap"
	"github.com/sairaph/apis-mcp/internal/budget"
	"github.com/sairaph/apis-mcp/internal/sessions"
	"gopkg.in/yaml.v3"
)

type pagination struct {
	Page       int `yaml:"page"`
	Total      int `yaml:"total"`
	TotalPages int `yaml:"total_pages"`
}

type sessionListFront struct {
	pagination `yaml:",inline"`
	Sessions   []sessionInfo `yaml:"sessions,omitempty"`
}

type sessionInspectFront struct {
	Session    sessionInfo `yaml:"session"`
	pagination `yaml:",inline"`
	Cookies    []cookieInfo `yaml:"cookies,omitempty"`
}

type sessionDeleteFront struct {
	Deleted string `yaml:"deleted"`
}

type sessionInfo struct {
	ID          string    `yaml:"id"`
	CreatedAt   time.Time `yaml:"created_at"`
	LastUsedAt  time.Time `yaml:"last_used_at"`
	CookieCount int       `yaml:"cookie_count"`
	Domains     []string  `yaml:"domains,omitempty"`
}

type cookieInfo struct {
	Name      string     `yaml:"name"`
	Value     string     `yaml:"value"`
	Domain    string     `yaml:"domain"`
	Path      string     `yaml:"path"`
	Expires   *time.Time `yaml:"expires,omitempty"`
	Secure    bool       `yaml:"secure"`
	HTTPOnly  bool       `yaml:"http_only"`
	SameSite  string     `yaml:"same_site"`
	HostOnly  bool       `yaml:"host_only"`
	CreatedAt time.Time  `yaml:"created_at"`
	UpdatedAt time.Time  `yaml:"updated_at"`
}

func sessionResult(_ context.Context, runtime *bootstrap.Runtime, input app.SessionsInput) (any, string, error) {
	if input.Delete {
		if input.ID == "" {
			return nil, "", invalidInput("id is required when delete is true", "Provide a server-generated session id from apis_sessions.", map[string]any{"field": "id"})
		}
		if input.Page != 0 {
			return nil, "", invalidInput("page is invalid when delete is true", "Omit page when deleting a session.", map[string]any{"field": "page"})
		}
		if err := runtime.Sessions.Delete(input.ID); err != nil {
			return nil, "", err
		}
		return sessionDeleteFront{Deleted: input.ID}, "Deleted cookie session `" + input.ID + "`. Cached HTTP responses were not deleted.", nil
	}

	page := input.Page
	if page == 0 {
		page = 1
	}
	if page < 1 {
		return nil, "", invalidInput("page must be one-based", "Use page 1 or greater.", map[string]any{"field": "page"})
	}
	budget := runtime.Config.ListTokenBudget
	if budget <= 0 {
		budget = 2_000
	}
	if input.ID == "" {
		items, err := runtime.Sessions.List()
		if err != nil {
			return nil, "", err
		}
		converted := make([]sessionInfo, len(items))
		for index, item := range items {
			converted[index] = makeSessionInfo(item)
		}
		window, pagination, err := paginate(converted, page, budget, func(records []sessionInfo) (string, error) {
			raw, err := yaml.Marshal(struct {
				Sessions []sessionInfo `yaml:"sessions"`
			}{Sessions: records})
			return string(raw), err
		})
		if err != nil {
			return nil, "", err
		}
		front := sessionListFront{pagination: pagination, Sessions: window}
		body := "No cookie sessions found."
		if len(window) > 0 {
			body = "Inspect a session with `" + toolCall("apis_sessions", app.SessionsInput{ID: window[0].ID}) + "`."
		}
		if pagination.Page < pagination.TotalPages {
			body += "\n\n" + continuation("apis_sessions", app.SessionsInput{Page: pagination.Page + 1})
		}
		return front, body, nil
	}

	inspection, err := runtime.Sessions.Inspect(input.ID)
	if err != nil {
		return nil, "", err
	}
	cookies := make([]cookieInfo, 0, len(inspection.Cookies))
	for _, cookie := range inspection.Cookies {
		item := cookieInfo{
			Name: cookie.Name, Value: cookie.Value, Domain: cookie.Domain, Path: cookie.Path,
			Secure: cookie.Secure, HTTPOnly: cookie.HTTPOnly, SameSite: sameSite(cookie.SameSite),
			HostOnly: cookie.HostOnly, CreatedAt: cookie.CreatedAt, UpdatedAt: cookie.UpdatedAt,
		}
		if !cookie.Expires.IsZero() {
			expires := cookie.Expires
			item.Expires = &expires
		}
		cookies = append(cookies, item)
	}
	window, pagination, err := paginate(cookies, page, budget, func(records []cookieInfo) (string, error) {
		raw, err := yaml.Marshal(struct {
			Cookies []cookieInfo `yaml:"cookies"`
		}{Cookies: records})
		return string(raw), err
	})
	if err != nil {
		return nil, "", err
	}
	front := sessionInspectFront{Session: makeSessionInfo(inspection.Session), pagination: pagination, Cookies: window}
	body := fmt.Sprintf("Showing %d of %d cookies for session `%s`.", len(window), pagination.Total, input.ID)
	if pagination.Total == 0 {
		body = "This session has no active cookies."
	}
	if pagination.Page < pagination.TotalPages {
		body += "\n\n" + continuation("apis_sessions", app.SessionsInput{ID: input.ID, Page: pagination.Page + 1})
	}
	return front, body, nil
}

func makeSessionInfo(item sessions.SessionInfo) sessionInfo {
	return sessionInfo{
		ID: item.ID, CreatedAt: item.CreatedAt, LastUsedAt: item.LastUsedAt,
		CookieCount: item.CookieCount, Domains: item.Domains,
	}
}

func paginate[T any](items []T, page, tokenBudget int, render func([]T) (string, error)) ([]T, pagination, error) {
	window, totalPages, err := budget.Paginate(items, page, tokenBudget, render)
	result := pagination{Page: page, Total: len(items), TotalPages: totalPages}
	if err != nil {
		return nil, result, fmt.Errorf("paginate sessions: %w", err)
	}
	return window, result, nil
}

func sameSite(value http.SameSite) string {
	switch value {
	case http.SameSiteDefaultMode:
		return "default"
	case http.SameSiteLaxMode:
		return "lax"
	case http.SameSiteStrictMode:
		return "strict"
	case http.SameSiteNoneMode:
		return "none"
	default:
		return fmt.Sprintf("unknown(%d)", value)
	}
}
