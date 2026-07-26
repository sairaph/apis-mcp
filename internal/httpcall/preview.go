package httpcall

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/sairaph/apis-mcp/internal/budget"
	"github.com/theory/jsonpath"
)

const maxPreviewBytes = 1 << 20

func makePreview(path, contentType string, decoded bool, tokenLimit int, jsonPath string, maxJSONInput int64) (*Preview, *Selection, error) {
	var selection *Selection
	if jsonPath != "" {
		selection = &Selection{JSONPath: jsonPath}
	}
	if !decoded {
		if selection != nil {
			selection.Error = "JSONPath is unavailable for an unsupported content encoding"
		}
		return nil, selection, nil
	}
	file, err := os.Open(path)
	if err != nil {
		if selection != nil {
			selection.Error = err.Error()
		}
		return nil, selection, nil
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, selection, nil
	}
	readLimit := int64(maxPreviewBytes + utf8.UTFMax)
	if stat.Size() <= maxJSONInput {
		readLimit = stat.Size() + 1
	}
	raw, err := io.ReadAll(io.LimitReader(file, readLimit))
	if err != nil {
		return nil, selection, nil
	}
	complete := int64(len(raw)) == stat.Size()

	var value any
	if complete && json.Unmarshal(raw, &value) == nil {
		selected := value
		if selection != nil {
			selected, err = selectJSON(value, jsonPath)
			if err != nil {
				selection.Error = err.Error()
				return nil, selection, nil
			}
			selection.Matched = true
		}
		pretty, err := json.MarshalIndent(selected, "", "  ")
		if err == nil {
			content, tokens, truncated, err := budget.Truncate(string(pretty), tokenLimit, maxPreviewBytes)
			if err != nil {
				return nil, selection, err
			}
			return &Preview{Kind: "json", Language: "json", Content: content, Truncated: truncated, ApproximateTokens: tokens}, selection, nil
		}
	}
	if selection != nil {
		selection.Error = "response body is not complete valid JSON"
		return nil, selection, nil
	}
	sample, valid := textPrefix(raw, maxPreviewBytes)
	if !valid || !textual(sample, contentType) {
		return nil, nil, nil
	}
	content, tokens, truncated, err := budget.Truncate(string(sample), tokenLimit, maxPreviewBytes)
	if err != nil {
		return nil, nil, err
	}
	if !complete || len(sample) != len(raw) {
		truncated = true
	}
	return &Preview{Kind: "text", Language: language(contentType), Content: content, Truncated: truncated, ApproximateTokens: tokens}, nil, nil
}

func textual(raw []byte, contentType string) bool {
	if !utf8.Valid(raw) || bytes.IndexByte(raw, 0) >= 0 {
		return false
	}
	mediaType, _, _ := mime.ParseMediaType(contentType)
	if strings.HasPrefix(strings.ToLower(mediaType), "text/") {
		return true
	}
	for _, b := range raw {
		if b < 0x09 || b > 0x0d && b < 0x20 {
			return false
		}
	}
	return true
}

func language(contentType string) string {
	mediaType, _, _ := mime.ParseMediaType(contentType)
	switch strings.ToLower(mediaType) {
	case "text/html":
		return "html"
	case "text/css":
		return "css"
	case "text/csv":
		return "csv"
	case "application/xml", "text/xml":
		return "xml"
	case "application/javascript", "text/javascript":
		return "javascript"
	default:
		return "text"
	}
}

func textPrefix(raw []byte, limit int) ([]byte, bool) {
	if len(raw) > limit {
		raw = raw[:limit]
	}
	for removed := 0; removed < utf8.UTFMax && len(raw) > 0; removed++ {
		if utf8.Valid(raw) {
			return raw, true
		}
		raw = raw[:len(raw)-1]
	}
	return nil, len(raw) == 0
}

func selectJSON(root any, expression string) (any, error) {
	path, err := jsonpath.Parse(expression)
	if err != nil {
		return nil, err
	}
	values := path.Select(root)
	if len(values) == 0 {
		return nil, errors.New("JSONPath did not match a value")
	}
	if path.Query().Singular() != nil {
		return values[0], nil
	}
	return []any(values), nil
}
