package httpcall

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const maxHeaderFile = 1 << 20
const maxPayloadFile = 50 << 20

type preparedInput struct {
	method          string
	endpoint        *url.URL
	headers         http.Header
	payload         []byte
	timeout         int
	retries         int
	session         string
	allowLarge      bool
	jsonPath        string
	automatic       map[string]string
	explicitCookies map[string]struct{}
}

func (s *Service) prepare(input Input) (preparedInput, error) {
	if !validToken(input.Method) || input.Method == "CONNECT" {
		return preparedInput{}, validation("method", "method must be a valid case-sensitive HTTP token other than CONNECT")
	}
	endpoint, err := parseEndpoint(input.Endpoint)
	if err != nil {
		return preparedInput{}, validation("endpoint", err.Error())
	}
	headers, supplied, err := parseHeaders(input.Headers)
	if err != nil {
		return preparedInput{}, validation("headers", err.Error())
	}
	payload, present, err := parsePayload(input.Payload)
	if err != nil {
		return preparedInput{}, validation("payload", err.Error())
	}
	automatic := make(map[string]string)
	if present && !supplied {
		headers.Set("Content-Type", "application/json")
		automatic["content-type"] = "application/json"
	}
	timeout := input.Timeout
	if timeout == 0 {
		timeout = 30
	}
	if timeout < 1 || timeout > s.maxTimeout {
		return preparedInput{}, validation("timeout", fmt.Sprintf("timeout must be between 1 and %d seconds", s.maxTimeout))
	}
	retries := inferredRetries(input.Method, headers)
	if retries > s.maxRetries {
		retries = s.maxRetries
	}
	if input.Retries != nil {
		retries = *input.Retries
	}
	if retries < 0 || retries > s.maxRetries {
		return preparedInput{}, validation("retries", fmt.Sprintf("retries must be between 0 and %d", s.maxRetries))
	}
	return preparedInput{
		method: input.Method, endpoint: endpoint, headers: headers, payload: payload,
		timeout: timeout, retries: retries, session: input.Session, allowLarge: input.AllowLargeDownload,
		jsonPath: input.JSONPath, automatic: automatic, explicitCookies: cookieNames(headers),
	}, nil
}

func parseEndpoint(raw string) (*url.URL, error) {
	if raw == "" {
		return nil, errors.New("endpoint is required")
	}
	u, err := url.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}
	if (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return nil, errors.New("endpoint must be an absolute HTTP or HTTPS URL with a host")
	}
	if _, err := http.NewRequest(http.MethodGet, u.String(), nil); err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}
	return u, nil
}

func parseHeaders(value any) (http.Header, bool, error) {
	headers := make(http.Header)
	if isNil(value) {
		return headers, false, nil
	}
	var object map[string]any
	switch value := value.(type) {
	case string:
		raw, err := readLimitedFile(value, maxHeaderFile)
		if err != nil {
			return nil, true, err
		}
		if err := decodeJSON(raw, &object); err != nil {
			return nil, true, err
		}
	default:
		raw, err := json.Marshal(value)
		if err != nil {
			return nil, true, fmt.Errorf("encode inline headers: %w", err)
		}
		if err := decodeJSON(raw, &object); err != nil {
			return nil, true, errors.New("headers must be a JSON object")
		}
	}
	for name, value := range object {
		if !validToken(name) {
			return nil, true, fmt.Errorf("invalid header name %q", name)
		}
		values, err := headerValues(value)
		if err != nil {
			return nil, true, fmt.Errorf("header %q: %w", name, err)
		}
		for _, item := range values {
			if !validHeaderValue(item) {
				return nil, true, fmt.Errorf("header %q contains an invalid value", name)
			}
			headers.Add(name, item)
		}
	}
	return headers, true, nil
}

func parsePayload(value any) ([]byte, bool, error) {
	if isNil(value) {
		return nil, false, nil
	}
	var raw []byte
	var err error
	if path, ok := value.(string); ok {
		raw, err = readLimitedFile(path, maxPayloadFile)
		if err != nil {
			return nil, true, err
		}
	} else {
		raw, err = json.Marshal(value)
		if err != nil {
			return nil, true, fmt.Errorf("encode inline payload: %w", err)
		}
	}
	var decoded any
	if err := decodeJSON(raw, &decoded); err != nil {
		return nil, true, err
	}
	switch decoded.(type) {
	case map[string]any, []any:
	default:
		return nil, true, errors.New("payload must be a JSON object or array")
	}
	return raw, true, nil
}

func decodeJSON(raw []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("parse JSON: %w", err)
	}
	if err := decoder.Decode(new(any)); !errors.Is(err, io.EOF) {
		return errors.New("JSON input contains trailing data")
	}
	return nil
}

func readLimitedFile(path string, limit int64) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read local JSON file %q: %w", path, err)
	}
	defer f.Close()
	raw, err := io.ReadAll(io.LimitReader(f, limit+1))
	if err != nil {
		return nil, fmt.Errorf("read local JSON file %q: %w", path, err)
	}
	if int64(len(raw)) > limit {
		return nil, fmt.Errorf("local JSON file %q exceeds %d bytes", path, limit)
	}
	return raw, nil
}

func headerValues(value any) ([]string, error) {
	switch value := value.(type) {
	case string:
		return []string{value}, nil
	case []any:
		result := make([]string, len(value))
		for i, item := range value {
			text, ok := item.(string)
			if !ok {
				return nil, errors.New("values must be strings or arrays of strings")
			}
			result[i] = text
		}
		return result, nil
	default:
		return nil, errors.New("value must be a string or array of strings")
	}
}

func validToken(value string) bool {
	if value == "" {
		return false
	}
	for i := 0; i < len(value); i++ {
		c := value[i]
		if c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			continue
		}
		switch c {
		case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '.', '^', '_', '`', '|', '~':
			continue
		}
		return false
	}
	return true
}

func validHeaderValue(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] == '\r' || value[i] == '\n' || value[i] == 0 {
			return false
		}
	}
	return true
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func inferredRetries(method string, headers http.Header) int {
	if headers.Get("Idempotency-Key") != "" {
		return 3
	}
	switch method {
	case "GET", "HEAD", "OPTIONS", "TRACE", "PUT", "DELETE":
		return 3
	default:
		return 0
	}
}

func cookieNames(headers http.Header) map[string]struct{} {
	result := make(map[string]struct{})
	for _, line := range headers.Values("Cookie") {
		for _, part := range strings.Split(line, ";") {
			name, _, ok := strings.Cut(strings.TrimSpace(part), "=")
			if ok && name != "" {
				result[name] = struct{}{}
			}
		}
	}
	return result
}

func validation(field, message string) *Error {
	return &Error{Code: "invalid_input", Message: message, Fields: map[string]any{"field": field}}
}
