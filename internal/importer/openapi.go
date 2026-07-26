package importer

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type apiDescription struct {
	OpenAPI    string                          `yaml:"openapi"`
	Swagger    string                          `yaml:"swagger"`
	Info       apiInfo                         `yaml:"info"`
	Servers    []apiServer                     `yaml:"servers"`
	Host       string                          `yaml:"host"`
	BasePath   string                          `yaml:"basePath"`
	Schemes    []string                        `yaml:"schemes"`
	Paths      map[string]map[string]yaml.Node `yaml:"paths"`
	Components struct {
		Schemas map[string]yaml.Node `yaml:"schemas"`
	} `yaml:"components"`
	Definitions map[string]yaml.Node `yaml:"definitions"`
}

type apiInfo struct {
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

type apiServer struct {
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

type apiOperation struct {
	Summary     string   `yaml:"summary"`
	Description string   `yaml:"description"`
	OperationID string   `yaml:"operationId"`
	Tags        []string `yaml:"tags"`
}

var operationMethods = map[string]bool{
	"get": true, "put": true, "post": true, "delete": true,
	"options": true, "head": true, "patch": true, "trace": true,
}

var (
	scriptConfigURL = regexp.MustCompile(`(?i)["']?(?:url|specUrl|spec-url)["']?\s*:\s*["']([^"']+)["']`)
	redocInitURL    = regexp.MustCompile(`(?i)\bRedoc\.init\s*\(\s*["']([^"']+)["']`)
)

// ImportOpenAPI imports an OpenAPI 3.x or Swagger 2.x JSON/YAML document.
func ImportOpenAPI(ctx context.Context, name, version, source string, options Options) (Result, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return Result{}, err
	}
	if strings.TrimSpace(name) == "" || strings.TrimSpace(version) == "" {
		return Result{}, errors.New("OpenAPI import requires API name and version")
	}
	reader := newSourceReader(options)
	raw, provenance, err := reader.read(ctx, source, nil)
	if err != nil {
		return Result{}, err
	}
	document, kind, parseErr := parseAPIDescription(raw)
	if parseErr != nil {
		landing, htmlErr := parseHTML(raw)
		if htmlErr != nil {
			return Result{}, fmt.Errorf("parse OpenAPI HTML landing page: %w", htmlErr)
		}
		if !isHTMLDocument(landing) {
			return Result{}, parseErr
		}
		base, urlErr := url.Parse(provenance)
		if urlErr != nil || base.Scheme != "http" && base.Scheme != "https" {
			return Result{}, errors.New("local OpenAPI HTML landing pages are not supported")
		}
		candidates := openAPISpecCandidates(landing, base)
		if len(candidates) == 0 {
			return Result{}, errors.New("OpenAPI HTML page contains no discoverable specification URL")
		}
		var candidateErrors []error
		for _, candidate := range candidates {
			candidateRaw, candidateSource, readErr := reader.read(ctx, candidate, base)
			if readErr != nil {
				candidateErrors = append(candidateErrors, readErr)
				continue
			}
			candidateDocument, candidateKind, candidateErr := parseAPIDescription(candidateRaw)
			if candidateErr != nil {
				candidateErrors = append(candidateErrors, fmt.Errorf("parse discovered specification %s: %w", candidateSource, candidateErr))
				continue
			}
			document, kind, provenance = candidateDocument, candidateKind, candidateSource
			parseErr = nil
			break
		}
		if parseErr != nil {
			return Result{}, fmt.Errorf("no discovered OpenAPI specification could be imported: %w", errors.Join(candidateErrors...))
		}
	}
	description := strings.TrimSpace(document.Info.Description)
	pages := 1
	result, err := publish(ctx, options, name, version, func(stage string) error {
		metadata := manifest{
			Name: name, Version: version, Description: description,
			SourceRoot: provenance, SourceType: kind, ImportedFrom: provenance,
		}
		if err := writeCanonicalFile(stage, "_index.md", metadata, "This document set was generated from an API description."); err != nil {
			return err
		}
		if err := writeCanonicalFile(stage, "overview.md", pageFront{
			Title: "Overview", PageID: "overview", Description: description,
			Source: provenance, SourceType: kind, ImportedFrom: provenance,
		}, openAPIOverview(name, version, provenance, document)); err != nil {
			return err
		}

		for _, endpoint := range sortedKeys(document.Paths) {
			pathItem := document.Paths[endpoint]
			for _, method := range sortedKeys(pathItem) {
				method = strings.ToLower(method)
				if !operationMethods[method] {
					continue
				}
				node := pathItem[method]
				var operation apiOperation
				if err := node.Decode(&operation); err != nil {
					return fmt.Errorf("parse operation %s %s: %w", strings.ToUpper(method), endpoint, err)
				}
				title := strings.TrimSpace(operation.Summary)
				if title == "" {
					title = strings.ToUpper(method) + " " + endpoint
				}
				description := strings.TrimSpace(operation.Description)
				if description == "" {
					description = strings.TrimSpace(operation.Summary)
				}
				tag := "untagged"
				if len(operation.Tags) > 0 && SafeSlug(operation.Tags[0]) != "" {
					tag = SafeSlug(operation.Tags[0])
				}
				identity := strings.ToUpper(method) + " " + endpoint
				front := pageFront{
					Title: title, PageID: stableID("operation-", identity), Path: "operations/" + tag,
					Description: description, Source: provenance, HTTPMethods: []string{strings.ToUpper(method)},
					APIEndpoints: []string{endpoint}, SourceType: kind, ImportedFrom: provenance,
				}
				if operation.OperationID != "" {
					front.OperationIDs = []string{operation.OperationID}
				}
				filename := stableID(method+"-", endpoint) + ".md"
				body, err := operationMarkdown(method, endpoint, operation, node, pathItem["parameters"])
				if err != nil {
					return err
				}
				if err := writeCanonicalFile(stage, filepath.Join("operations", tag, filename), front, body); err != nil {
					return err
				}
				pages++
			}
		}

		schemas := document.Components.Schemas
		if kind == "swagger" {
			schemas = document.Definitions
		}
		for _, schemaName := range sortedKeys(schemas) {
			node := schemas[schemaName]
			description := nodeString(node, "description")
			front := pageFront{
				Title: schemaName, PageID: stableID("schema-", schemaName), Path: "schemas",
				Description: description, Source: provenance, SourceType: kind, ImportedFrom: provenance,
			}
			rendered, err := yaml.Marshal(node)
			if err != nil {
				return err
			}
			body := "# " + schemaName + "\n\n"
			if description != "" {
				body += description + "\n\n"
			}
			body += "```yaml\n" + string(rendered) + "```"
			if err := writeCanonicalFile(stage, filepath.Join("schemas", stableID("", schemaName)+".md"), front, body); err != nil {
				return err
			}
			pages++
		}
		return nil
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind, result.Source, result.Pages = kind, provenance, pages
	return result, nil
}

func parseAPIDescription(raw []byte) (apiDescription, string, error) {
	var document apiDescription
	if err := yaml.Unmarshal(raw, &document); err != nil {
		return apiDescription{}, "", fmt.Errorf("parse OpenAPI document: %w", err)
	}
	kind := ""
	switch {
	case strings.HasPrefix(strings.TrimSpace(document.OpenAPI), "3."):
		kind = "openapi"
	case strings.HasPrefix(strings.TrimSpace(document.Swagger), "2."):
		kind = "swagger"
	default:
		return apiDescription{}, "", errors.New("source is not an OpenAPI 3.x or Swagger 2.x document")
	}
	if len(document.Paths) == 0 {
		return apiDescription{}, "", errors.New("OpenAPI document has no paths")
	}
	return document, kind, nil
}

func openAPISpecCandidates(root *htmlNode, base *url.URL) []string {
	var candidates []string
	add := func(value string) {
		value = strings.TrimSpace(strings.ReplaceAll(value, `\/`, `/`))
		if value == "" || strings.ContainsAny(value, "{}\n\r") {
			return
		}
		resolved, err := base.Parse(value)
		if err != nil || resolved.Host == "" || resolved.Scheme != "http" && resolved.Scheme != "https" {
			return
		}
		resolved.Fragment = ""
		candidates = append(candidates, resolved.String())
	}
	walkHTML(root, func(node *htmlNode) {
		identity := strings.ToLower(node.tag + " " + node.attrs["id"] + " " + node.attrs["class"])
		for _, attribute := range []string{"spec-url", "data-spec-url", "data-swagger-url"} {
			add(node.attrs[attribute])
		}
		if node.tag == "script" || strings.Contains(identity, "swagger") || strings.Contains(identity, "redoc") || strings.Contains(identity, "openapi") || strings.Contains(identity, "rapi-doc") {
			add(node.attrs["data-url"])
		}
		if node.tag == "link" {
			relationship := strings.ToLower(node.attrs["rel"] + " " + node.attrs["type"])
			if strings.Contains(relationship, "service-desc") || strings.Contains(relationship, "openapi") || strings.Contains(relationship, "swagger") {
				add(node.attrs["href"])
			}
		}
		if node.tag == "meta" {
			name := strings.ToLower(node.attrs["name"] + " " + node.attrs["property"])
			if strings.Contains(name, "openapi") || strings.Contains(name, "swagger") || strings.Contains(name, "spec-url") {
				add(node.attrs["content"])
			}
		}
		if node.tag != "script" {
			return
		}
		script := htmlNodeText(node)
		lower := strings.ToLower(script + " " + node.attrs["id"] + " " + node.attrs["class"])
		if strings.Contains(lower, "swaggeruibundle") || strings.Contains(lower, "swaggerui(") || strings.Contains(lower, "swagger-config") {
			for _, match := range scriptConfigURL.FindAllStringSubmatch(script, -1) {
				add(match[1])
			}
		}
		if strings.Contains(lower, "redoc") {
			for _, match := range redocInitURL.FindAllStringSubmatch(script, -1) {
				add(match[1])
			}
			for _, match := range scriptConfigURL.FindAllStringSubmatch(script, -1) {
				add(match[1])
			}
		}
	})
	unique := make(map[string]bool)
	output := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		if !unique[candidate] {
			unique[candidate] = true
			output = append(output, candidate)
		}
	}
	return output
}

func openAPIOverview(name, version, provenance string, document apiDescription) string {
	var body strings.Builder
	fmt.Fprintf(&body, "# %s %s\n\n", name, version)
	if description := strings.TrimSpace(document.Info.Description); description != "" {
		body.WriteString(description + "\n\n")
	}
	fmt.Fprintf(&body, "Source: `%s`\n", provenance)
	var servers []string
	for _, server := range document.Servers {
		if strings.TrimSpace(server.URL) != "" {
			servers = append(servers, server.URL)
		}
	}
	if len(servers) == 0 && document.Host != "" {
		schemes := document.Schemes
		if len(schemes) == 0 {
			schemes = []string{"https"}
		}
		for _, scheme := range schemes {
			servers = append(servers, scheme+"://"+document.Host+document.BasePath)
		}
	}
	if len(servers) > 0 {
		sort.Strings(servers)
		body.WriteString("\n## Servers\n\n")
		for _, server := range servers {
			fmt.Fprintf(&body, "- `%s`\n", server)
		}
	}
	return body.String()
}

func operationMarkdown(method, endpoint string, operation apiOperation, node, pathParameters yaml.Node) (string, error) {
	var body strings.Builder
	title := strings.TrimSpace(operation.Summary)
	if title == "" {
		title = strings.ToUpper(method) + " " + endpoint
	}
	fmt.Fprintf(&body, "# %s\n\n`%s %s`\n", title, strings.ToUpper(method), endpoint)
	if operation.OperationID != "" {
		fmt.Fprintf(&body, "\nOperation ID: `%s`\n", operation.OperationID)
	}
	if operation.Description != "" {
		body.WriteString("\n" + strings.TrimSpace(operation.Description) + "\n")
	}
	if !pathParameters.IsZero() {
		rendered, err := yaml.Marshal(pathParameters)
		if err != nil {
			return "", err
		}
		body.WriteString("\n## Path Parameters\n\n```yaml\n" + string(rendered) + "```\n")
	}
	rendered, err := yaml.Marshal(node)
	if err != nil {
		return "", err
	}
	body.WriteString("\n## Definition\n\n```yaml\n" + string(rendered) + "```\n")
	return body.String(), nil
}

func nodeString(node yaml.Node, key string) string {
	var values map[string]any
	if err := node.Decode(&values); err != nil {
		return ""
	}
	value, _ := values[key].(string)
	return strings.TrimSpace(value)
}
