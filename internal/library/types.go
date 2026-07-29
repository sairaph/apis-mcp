package library

import (
	"errors"
	"fmt"
)

const (
	defaultListTokenBudget = 2_000
	defaultReadTokenBudget = 4_000
	DefaultBrowseLimit     = 75
	MaxBrowseLimit         = 100
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrNotFound        = errors.New("documentation not found")
)

// Options identifies official pack archives, the canonical user tree, and the
// published SQLite index. Token budgets apply to retrieved records, not
// transport rendering overhead.
type Options struct {
	UserRoot        string
	IndexPath       string
	ListTokenBudget int
	ReadTokenBudget int
	PackArchives    []string
}

type CollectionsRequest struct {
	Page int `json:"page,omitempty" yaml:"page,omitempty"`
}

type ListRequest struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Version    string `json:"version,omitempty" yaml:"version,omitempty"`
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`
	Page       int    `json:"page,omitempty" yaml:"page,omitempty"`
}

type PagesRequest struct {
	DocID string `json:"doc_id" yaml:"doc_id"`
	Path  string `json:"path,omitempty" yaml:"path,omitempty"`
	Page  int    `json:"page,omitempty" yaml:"page,omitempty"`
}

// BrowseRequest selects a fixed navigation window without applying MCP token
// pagination. Limit defaults to DefaultBrowseLimit and is capped at
// MaxBrowseLimit.
type BrowseRequest struct {
	DocID  string `json:"doc_id" yaml:"doc_id"`
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Offset int    `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty" yaml:"limit,omitempty"`
}

type SearchRequest struct {
	DocID string `json:"doc_id" yaml:"doc_id"`
	Query string `json:"query" yaml:"query"`
	Path  string `json:"path,omitempty" yaml:"path,omitempty"`
	Page  int    `json:"page,omitempty" yaml:"page,omitempty"`
}

type ReadRequest struct {
	DocID  string `json:"doc_id" yaml:"doc_id"`
	PageID string `json:"page_id" yaml:"page_id"`
	Lines  []int  `json:"lines,omitempty" yaml:"lines,omitempty"`
}

type Pagination struct {
	Page       int `yaml:"page" json:"page"`
	Total      int `yaml:"total" json:"total"`
	TotalPages int `yaml:"total_pages" json:"total_pages"`
}

type Collection struct {
	Collection  string `yaml:"collection" json:"collection"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	APICount    int    `yaml:"api_count" json:"api_count"`
}

type CollectionsResult struct {
	Pagination  `yaml:",inline"`
	Collections []Collection `yaml:"collections,omitempty" json:"collections,omitempty"`
}

type APIVersion struct {
	Version string `yaml:"version" json:"version"`
	DocID   string `yaml:"doc_id" json:"doc_id"`
	Pages   int    `yaml:"pages" json:"pages"`
}

type API struct {
	Name        string       `yaml:"name" json:"name"`
	Description string       `yaml:"description,omitempty" json:"description,omitempty"`
	Collections []string     `yaml:"collections,omitempty" json:"collections,omitempty"`
	Versions    []APIVersion `yaml:"versions" json:"versions"`
}

type ListResult struct {
	Pagination `yaml:",inline"`
	APIs       []API `yaml:"apis,omitempty" json:"apis,omitempty"`
}

type Path struct {
	Path        string `yaml:"path" json:"path"`
	NestedPages int    `yaml:"nested_pages" json:"nested_pages"`
}

type Page struct {
	PageID      string `yaml:"page_id" json:"page_id"`
	Title       string `yaml:"title" json:"title"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
}

type PagesResult struct {
	DocID      string `yaml:"doc_id" json:"doc_id"`
	Path       string `yaml:"path,omitempty" json:"path,omitempty"`
	Pagination `yaml:",inline"`
	Paths      []Path `yaml:"paths,omitempty" json:"paths,omitempty"`
	Pages      []Page `yaml:"pages,omitempty" json:"pages,omitempty"`
}

type BrowseResult struct {
	DocID  string `json:"doc_id" yaml:"doc_id"`
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Offset int    `json:"offset" yaml:"offset"`
	Total  int    `json:"total" yaml:"total"`
	Paths  []Path `json:"paths,omitempty" yaml:"paths,omitempty"`
	Pages  []Page `json:"pages,omitempty" yaml:"pages,omitempty"`
}

type SearchHit struct {
	PageID  string `yaml:"page_id" json:"page_id"`
	Line    int    `yaml:"line,omitempty" json:"line,omitempty"`
	Title   string `yaml:"title" json:"title"`
	Path    string `yaml:"path,omitempty" json:"path,omitempty"`
	Match   string `yaml:"match" json:"match"`
	Snippet string `yaml:"snippet" json:"snippet"`
}

type SearchResult struct {
	DocID      string `yaml:"doc_id" json:"doc_id"`
	Query      string `yaml:"query" json:"query"`
	Path       string `yaml:"path,omitempty" json:"path,omitempty"`
	Pagination `yaml:",inline"`
	Hits       []SearchHit `yaml:"hits,omitempty" json:"hits,omitempty"`
}

type ReadResult struct {
	DocID      string `yaml:"doc_id" json:"doc_id"`
	Title      string `yaml:"title" json:"title"`
	PageID     string `yaml:"page_id" json:"page_id"`
	Path       string `yaml:"path,omitempty" json:"path,omitempty"`
	Lines      [2]int `yaml:"lines" json:"lines"`
	TotalLines int    `yaml:"total_lines" json:"total_lines"`
	Truncated  bool   `yaml:"truncated" json:"truncated"`
	Markdown   string `yaml:"-" json:"markdown"`
}

// ValidationError reports a canonical source error and, where applicable, all
// source locations participating in a collision.
type ValidationError struct {
	Message   string
	Locations []string
}

func (e *ValidationError) Error() string {
	if len(e.Locations) == 0 {
		return e.Message
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Locations)
}

func normalizeOptions(options Options) (Options, error) {
	if options.IndexPath == "" {
		return Options{}, fmt.Errorf("%w: index path is required", ErrInvalidArgument)
	}
	if options.ListTokenBudget == 0 {
		options.ListTokenBudget = defaultListTokenBudget
	}
	if options.ReadTokenBudget == 0 {
		options.ReadTokenBudget = defaultReadTokenBudget
	}
	if options.ListTokenBudget < 1 || options.ReadTokenBudget < 1 {
		return Options{}, fmt.Errorf("%w: token budgets must be positive", ErrInvalidArgument)
	}
	return options, nil
}
