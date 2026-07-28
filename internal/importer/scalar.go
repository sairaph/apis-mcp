package importer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const maxScalarNextRouteScripts = 3

var (
	scalarStaticJSURL  = regexp.MustCompile("(?is)(?:^|[,{])\\s*[\\\"']?(?:url|specUrl|spec-url)[\\\"']?\\s*:\\s*(?:\\\"([^\\\"\\\\\\r\\n]*)\\\"|'([^'\\\\\\r\\n]*)'|`([^`$\\\\\\r\\n]*)`)")
	scalarConfigObject = regexp.MustCompile(`(?is)(?:^|[,{])\s*["']?configuration["']?\s*:\s*\{([^{}]{0,1000})\}`)
	scalarConfigMarker = regexp.MustCompile(`(?is)(?:^|[,{])\s*["']?configuration["']?\s*:\s*\{\s*["']?url["']?\s*:`)
	scalarBaseURLAlias = regexp.MustCompile(`(?:\{|,)\s*baseURL\s*:\s*([A-Za-z_$][A-Za-z0-9_$]*)\s*(?:[,}])`)
)

type scalarSchemaResolution struct {
	Proven bool
	URL    string
}

type scalarNextData struct {
	Page    string
	BaseURL string
}

// resolveScalarSchema only resolves literal Scalar configuration and one
// constrained Next.js composition: a static pageProps.baseURL followed by a
// literal suffix in the route chunk. It never evaluates JavaScript.
func resolveScalarSchema(ctx context.Context, root *htmlNode, base *url.URL, reader *sourceReader) (scalarSchemaResolution, error) {
	candidates, proven := scalarHTMLSchemaCandidates(root, base)
	if len(candidates) > 0 {
		return stableScalarSchema(candidates, proven), nil
	}

	next, ok := parseScalarNextData(root)
	if !ok {
		return scalarSchemaResolution{Proven: proven}, nil
	}
	scripts := scalarNextRouteScripts(root, base, next.Page)
	for _, scriptURL := range scripts {
		if err := ctx.Err(); err != nil {
			return scalarSchemaResolution{}, err
		}
		raw, _, err := reader.readFromOrigin(ctx, scriptURL, nil, httpOrigin(base))
		if err != nil {
			return scalarSchemaResolution{}, fmt.Errorf("inspect Scalar Next.js route script %s: %w", scriptURL, err)
		}
		routeCandidates, routeProven := scalarRouteScriptCandidates(string(raw), next, base)
		proven = proven || routeProven
		candidates = append(candidates, routeCandidates...)
	}
	return stableScalarSchema(candidates, proven), nil
}

func stableScalarSchema(candidates []string, proven bool) scalarSchemaResolution {
	candidates = uniqueStrings(candidates)
	if len(candidates) == 1 {
		return scalarSchemaResolution{Proven: true, URL: candidates[0]}
	}
	return scalarSchemaResolution{Proven: proven || len(candidates) > 0}
}

func scalarHTMLSchemaCandidates(root *htmlNode, base *url.URL) ([]string, bool) {
	var candidates []string
	proven := false
	walkHTML(root, func(node *htmlNode) {
		switch node.tag {
		case "scalar-api-reference":
			proven = true
			for _, attribute := range []string{"url", "spec-url", "data-url"} {
				appendScalarSchemaURL(&candidates, base, node.attrs[attribute])
			}
			for _, attribute := range []string{"configuration", "data-configuration"} {
				appendScalarJSONConfigURLs(&candidates, base, node.attrs[attribute])
			}
		case "script":
			text := htmlNodeText(node)
			lower := strings.ToLower(text)
			runtimeMarker := scalarRuntimeMarker(node, lower)
			configurationScript := scalarConfigurationScript(node, runtimeMarker)
			if runtimeMarker || configurationScript {
				proven = true
			}
			if configurationScript {
				for _, attribute := range []string{"data-url", "url", "spec-url"} {
					appendScalarSchemaURL(&candidates, base, node.attrs[attribute])
				}
				for _, attribute := range []string{"configuration", "data-configuration"} {
					appendScalarJSONConfigURLs(&candidates, base, node.attrs[attribute])
				}
				appendScalarJSONConfigURLs(&candidates, base, text)
			}
			if runtimeMarker {
				for _, call := range scalarCreateReferenceCalls(text) {
					appendScalarStaticJSURLs(&candidates, base, call)
				}
			}
		}
	})
	return uniqueStrings(candidates), proven
}

func scalarScriptEvidence(node *htmlNode, text string) bool {
	lower := strings.ToLower(text)
	runtimeMarker := scalarRuntimeMarker(node, lower)
	return runtimeMarker || scalarConfigurationScript(node, runtimeMarker)
}

func scalarRuntimeMarker(node *htmlNode, lowerText string) bool {
	identity := strings.ToLower(node.attrs["src"] + " " + node.attrs["class"])
	return strings.Contains(identity, "@scalar/api-reference") || strings.Contains(lowerText, "scalar.createapireference(")
}

func scalarConfigurationScript(node *htmlNode, runtimeMarker bool) bool {
	if !strings.EqualFold(strings.TrimSpace(node.attrs["id"]), "api-reference") {
		return false
	}
	for _, attribute := range []string{"data-configuration", "data-url", "data-proxy-url"} {
		if _, present := node.attrs[attribute]; present {
			return true
		}
	}
	contentType := strings.ToLower(strings.TrimSpace(node.attrs["type"]))
	return runtimeMarker || contentType == "application/json" || contentType == "application/yaml" || contentType == "text/yaml"
}

func appendScalarJSONConfigURLs(candidates *[]string, base *url.URL, raw string) {
	if strings.TrimSpace(raw) == "" {
		return
	}
	var configuration map[string]any
	if json.Unmarshal([]byte(raw), &configuration) != nil {
		return
	}
	for _, key := range []string{"url", "specUrl", "spec-url"} {
		if value, ok := configuration[key].(string); ok {
			appendScalarSchemaURL(candidates, base, value)
		}
	}
	if sources, ok := configuration["sources"].([]any); ok {
		for _, source := range sources {
			entry, ok := source.(map[string]any)
			if !ok {
				continue
			}
			if value, ok := entry["url"].(string); ok {
				appendScalarSchemaURL(candidates, base, value)
			}
		}
	}
}

func scalarCreateReferenceCalls(script string) []string {
	lower := strings.ToLower(script)
	var calls []string
	for offset := 0; offset < len(lower); {
		index := strings.Index(lower[offset:], "createapireference")
		if index < 0 {
			break
		}
		index += offset + len("createapireference")
		open := index
		for open < len(script) && (script[open] == ' ' || script[open] == '\t' || script[open] == '\r' || script[open] == '\n') {
			open++
		}
		if open >= len(script) || script[open] != '(' {
			offset = index
			continue
		}
		if end := balancedJSCallEnd(script, open); end > open {
			calls = append(calls, script[open+1:end])
			offset = end + 1
		} else {
			break
		}
	}
	return calls
}

func balancedJSCallEnd(script string, open int) int {
	depth := 0
	quote := byte(0)
	escaped := false
	for index := open; index < len(script); index++ {
		character := script[index]
		if quote != 0 {
			if escaped {
				escaped = false
				continue
			}
			if character == '\\' {
				escaped = true
				continue
			}
			if character == quote {
				quote = 0
			}
			continue
		}
		switch character {
		case '\'', '"', '`':
			quote = character
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return -1
}

func appendScalarStaticJSURLs(candidates *[]string, base *url.URL, script string) {
	for _, match := range scalarStaticJSURL.FindAllStringSubmatchIndex(script, -1) {
		end := match[1]
		for end < len(script) && (script[end] == ' ' || script[end] == '\t' || script[end] == '\r' || script[end] == '\n') {
			end++
		}
		if end < len(script) && script[end] != ',' && script[end] != '}' && script[end] != ']' && script[end] != ')' {
			continue
		}
		for index := 2; index+1 < len(match); index += 2 {
			if match[index] >= 0 {
				appendScalarSchemaURL(candidates, base, script[match[index]:match[index+1]])
				break
			}
		}
	}
}

func appendScalarSchemaURL(candidates *[]string, base *url.URL, value string) {
	value = strings.TrimSpace(strings.ReplaceAll(value, `\/`, `/`))
	if value == "" || strings.ContainsAny(value, "{}\n\r") {
		return
	}
	resolved, err := base.Parse(value)
	if err != nil || resolved.User != nil || resolved.Host == "" || resolved.Scheme != "http" && resolved.Scheme != "https" {
		return
	}
	resolved.Fragment = ""
	*candidates = append(*candidates, resolved.String())
}

func parseScalarNextData(root *htmlNode) (scalarNextData, bool) {
	var result scalarNextData
	found := false
	walkHTML(root, func(node *htmlNode) {
		if found || node.tag != "script" || node.attrs["id"] != "__NEXT_DATA__" {
			return
		}
		var envelope struct {
			Page  string `json:"page"`
			Props struct {
				PageProps map[string]json.RawMessage `json:"pageProps"`
			} `json:"props"`
		}
		if json.Unmarshal([]byte(htmlNodeText(node)), &envelope) != nil || !validNextRoutePage(envelope.Page) {
			return
		}
		result.Page = envelope.Page
		if raw := envelope.Props.PageProps["baseURL"]; raw != nil {
			_ = json.Unmarshal(raw, &result.BaseURL)
		}
		found = true
	})
	return result, found
}

func validNextRoutePage(page string) bool {
	return strings.HasPrefix(page, "/") && !strings.Contains(page, "//") && !strings.ContainsAny(page, "\\?#") && !hasDotPathSegment(page)
}

func scalarNextRouteScripts(root *htmlNode, base *url.URL, page string) []string {
	var scripts []string
	walkHTML(root, func(node *htmlNode) {
		if len(scripts) >= maxScalarNextRouteScripts || node.tag != "script" || strings.TrimSpace(node.attrs["src"]) == "" {
			return
		}
		candidate, err := base.Parse(strings.TrimSpace(node.attrs["src"]))
		if err != nil || candidate.User != nil || candidate.RawPath != "" || candidate.RawQuery != "" || candidate.ForceQuery || candidate.Fragment != "" || candidate.RawFragment != "" || !sameHTTPOrigin(base, candidate) || !isNextRouteScript(candidate.Path, page) {
			return
		}
		scripts = append(scripts, candidate.String())
	})
	return uniqueStrings(scripts)
}

func isNextRouteScript(assetPath, page string) bool {
	route := strings.TrimPrefix(page, "/")
	pagesRoute := route
	if pagesRoute == "" {
		pagesRoute = "index"
	}
	var hash string
	pagesPrefix := "/_next/static/chunks/pages/" + pagesRoute + "-"
	appPrefix := "/_next/static/chunks/app/"
	if route != "" {
		appPrefix += route + "/"
	}
	appPrefix += "page-"
	switch {
	case strings.HasPrefix(assetPath, pagesPrefix) && strings.HasSuffix(assetPath, ".js"):
		hash = strings.TrimSuffix(strings.TrimPrefix(assetPath, pagesPrefix), ".js")
	case strings.HasPrefix(assetPath, appPrefix) && strings.HasSuffix(assetPath, ".js"):
		hash = strings.TrimSuffix(strings.TrimPrefix(assetPath, appPrefix), ".js")
	default:
		return false
	}
	if len(hash) < 6 {
		return false
	}
	for _, character := range hash {
		if character < 'a' || character > 'z' {
			if character < 'A' || character > 'Z' {
				if character < '0' || character > '9' {
					if character != '-' && character != '_' {
						return false
					}
				}
			}
		}
	}
	return true
}

func scalarRouteScriptCandidates(script string, next scalarNextData, landing *url.URL) ([]string, bool) {
	lower := strings.ToLower(script)
	proven := strings.Contains(lower, "@scalar/api-reference") || strings.Contains(lower, "scalar.createapireference(")
	configurations := scalarConfigObject.FindAllStringSubmatch(script, -1)
	if scalarConfigMarker.MatchString(script) && strings.Contains(lower, "forcedarkmodestate") && strings.Contains(lower, "hidetestrequestbutton") {
		proven = true
	}
	if !proven {
		return nil, false
	}

	var candidates []string
	for _, configuration := range configurations {
		appendScalarStaticJSURLs(&candidates, landing, configuration[1])
	}
	if strings.TrimSpace(next.BaseURL) == "" {
		return uniqueStrings(candidates), true
	}
	baseURL, err := url.Parse(strings.TrimSpace(next.BaseURL))
	if err != nil || baseURL.User != nil || baseURL.Host == "" || baseURL.Scheme != "http" && baseURL.Scheme != "https" || baseURL.Fragment != "" {
		return uniqueStrings(candidates), true
	}
	for _, alias := range scalarBaseURLAlias.FindAllStringSubmatch(script, -1) {
		pattern := regexp.MustCompile("(?is)(?:^|[,{])\\s*[\\\"']?url[\\\"']?\\s*:\\s*`\\$\\{" + regexp.QuoteMeta(alias[1]) + "\\}([^`$\\\\\\r\\n]*)`")
		for _, match := range pattern.FindAllStringSubmatch(script, -1) {
			appendScalarSchemaURL(&candidates, landing, baseURL.String()+match[1])
		}
	}
	return uniqueStrings(candidates), true
}

func scalarUnsupportedError() error {
	return errors.New("Scalar API Reference contains no single statically resolvable specification URL")
}
