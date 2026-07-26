package app

type CollectionsInput struct {
	Page int `json:"page,omitempty" jsonschema:"One-based result page; defaults to 1"`
}

type ListInput struct {
	Name       string `json:"name,omitempty" jsonschema:"Case-insensitive substring matched against API names"`
	Version    string `json:"version,omitempty" jsonschema:"Case-insensitive exact version match"`
	Collection string `json:"collection,omitempty" jsonschema:"Case-insensitive exact collection identifier"`
	Page       int    `json:"page,omitempty" jsonschema:"One-based result page; defaults to 1"`
}

type PagesInput struct {
	DocID string `json:"doc_id" jsonschema:"Documentation version returned by apis_list"`
	Path  string `json:"path,omitempty" jsonschema:"Exact documentation category path"`
	Page  int    `json:"page,omitempty" jsonschema:"One-based result page; defaults to 1"`
}

type SearchInput struct {
	DocID string `json:"doc_id" jsonschema:"Documentation version to search"`
	Query string `json:"query" jsonschema:"Terms and quoted phrases to find"`
	Path  string `json:"path,omitempty" jsonschema:"Restrict results to this path and descendants"`
	Page  int    `json:"page,omitempty" jsonschema:"One-based result page; defaults to 1"`
}

type ReadInput struct {
	DocID  string `json:"doc_id" jsonschema:"Documentation version containing the page"`
	PageID string `json:"page_id" jsonschema:"Page identifier returned by apis_pages or apis_search"`
	Lines  []int  `json:"lines,omitempty" jsonschema:"Inclusive one-based start and end line"`
}

type CallInput struct {
	Method             string `json:"method" jsonschema:"Case-sensitive HTTP method token other than CONNECT"`
	Endpoint           string `json:"endpoint" jsonschema:"Complete absolute HTTP or HTTPS URL"`
	Headers            any    `json:"headers,omitempty" jsonschema:"Inline header object or local JSON file path"`
	Payload            any    `json:"payload,omitempty" jsonschema:"Inline JSON object or array, or local JSON file path"`
	Timeout            int    `json:"timeout,omitempty" jsonschema:"Response-header timeout in seconds; defaults to 30"`
	Retries            *int   `json:"retries,omitempty" jsonschema:"Retries after the initial attempt"`
	JSONPath           string `json:"json_path,omitempty" jsonschema:"RFC 9535 JSONPath for the embedded preview"`
	Session            string `json:"session,omitempty" jsonschema:"Existing server-generated cookie session ID"`
	AllowLargeDownload bool   `json:"allow_large_download,omitempty" jsonschema:"Bypass the configured response-size cap"`
}

type SessionsInput struct {
	ID     string `json:"id,omitempty" jsonschema:"Server-generated cookie session ID"`
	Delete bool   `json:"delete,omitempty" jsonschema:"Delete the identified session"`
	Page   int    `json:"page,omitempty" jsonschema:"One-based result page; defaults to 1"`
}

type Error struct {
	Code    string
	Message string
	Hint    string
	Fields  map[string]any
}

func (e *Error) Error() string { return e.Message }
