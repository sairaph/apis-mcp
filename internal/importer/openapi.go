package importer

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path"
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
	Webhooks   map[string]map[string]yaml.Node `yaml:"webhooks"`
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
	scriptConfigURL      = regexp.MustCompile(`(?i)(?:^|[,{]\s*)["']?(?:url|specUrl|spec-url)["']?\s*:\s*["']([^"']+)["']`)
	scriptConfigVariable = regexp.MustCompile(`(?i)(?:^|[,{]\s*)["']?(?:url|specUrl|spec-url)["']?\s*:\s*([A-Za-z_$][A-Za-z0-9_$]*)`)
	scriptStringVariable = regexp.MustCompile(`\b(?:const|let|var)\s+([A-Za-z_$][A-Za-z0-9_$]*)\s*=\s*["']([^"']+)["']`)
	scriptVariableAlias  = regexp.MustCompile(`\b(?:const|let|var)\s+([A-Za-z_$][A-Za-z0-9_$]*)\s*=\s*([A-Za-z_$][A-Za-z0-9_$]*)\s*;`)
	hostConfigURL        = regexp.MustCompile("(?i)([a-z0-9.-]+)=(https?://[^\\s\"'`,;]+)")
	redocInitURL         = regexp.MustCompile(`(?i)\bRedoc\.init\s*\(\s*["']([^"']+)["']`)
)

type apiSpecSource struct {
	URL       string
	Namespace string
	Label     string
	Origin    string
}

type apiCatalog struct {
	APIs []struct {
		Path    string   `yaml:"path"`
		Schema  string   `yaml:"schema"`
		Formats []string `yaml:"format"`
	} `yaml:"apis"`
	BasePath string `yaml:"basePath"`
}

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
		catalogSources, catalog, catalogErr := rapidocSources(ctx, landing, base, reader)
		if catalogErr != nil {
			return Result{}, catalogErr
		}
		if catalog {
			return importOpenAPICatalog(ctx, name, version, provenance, catalogSources, reader, options)
		}
		candidates := openAPISpecCandidates(landing, base)
		configScripts := openAPIConfigScripts(landing, base)
		if len(candidates) == 0 && len(configScripts) == 0 {
			return Result{}, errors.New("OpenAPI HTML page contains no discoverable specification URL")
		}
		var candidateErrors []error
		found := false
		tryCandidates := func(discovered []string) {
			for _, candidate := range discovered {
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
				raw, document, kind, provenance = candidateRaw, candidateDocument, candidateKind, candidateSource
				found = true
				return
			}
		}
		tryCandidates(candidates)
		for _, scriptSource := range configScripts {
			if found {
				break
			}
			script, _, readErr := reader.read(ctx, scriptSource, base)
			if readErr != nil {
				candidateErrors = append(candidateErrors, readErr)
				continue
			}
			tryCandidates(openAPIConfigCandidates(string(script), base))
		}
		if !found {
			if len(candidateErrors) == 0 {
				return Result{}, errors.New("OpenAPI HTML page contains no discoverable specification URL")
			}
			return Result{}, fmt.Errorf("no discovered OpenAPI specification could be imported: %w", errors.Join(candidateErrors...))
		}
	}
	document, sourceCount, err := resolveExternalOpenAPIRefs(ctx, raw, provenance, kind, document, reader)
	if err != nil {
		return Result{}, err
	}
	if !hasAPIDescriptionOperations(document) {
		return Result{}, errors.New("OpenAPI document has no path or webhook operations")
	}
	return importAPIDescription(ctx, name, version, document, kind, provenance, sourceCount, options)
}

func importAPIDescription(ctx context.Context, name, version string, document apiDescription, kind, provenance string, sources int, options Options) (Result, error) {
	description := strings.TrimSpace(document.Info.Description)
	pages := 0
	reportProgress(options, Progress{Stage: "generating", Framework: kind, URL: provenance, Pages: pages})
	result, err := publish(ctx, options, name, version, func(stage string) error {
		metadata := manifest{
			Name: name, Version: version, Description: description, Collections: options.Collections,
			SourceRoot: provenance, SourceType: kind, ImportedFrom: provenance, Sources: sources,
		}
		if err := writeCanonicalFile(stage, "_index.md", metadata, "This document set was generated from an API description."); err != nil {
			return err
		}
		return writeAPIDescription(stage, "", "", name, version, provenance, kind, document, &pages, 0, options)
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind, result.Source, result.Pages, result.Sources = kind, provenance, pages, sources
	reportProgress(options, Progress{Stage: "completed", Framework: kind, URL: provenance, Pages: pages})
	return result, nil
}

func importOpenAPICatalog(ctx context.Context, name, version, provenance string, sources []apiSpecSource, reader *sourceReader, options Options) (Result, error) {
	if len(sources) == 0 {
		return Result{}, errors.New("RapiDoc catalog contains no specifications")
	}
	pages := 0
	sourceCount := 0
	reportProgress(options, Progress{Stage: "generating", Framework: "rapidoc", URL: provenance, Queued: len(sources)})
	result, err := publish(ctx, options, name, version, func(stage string) error {
		description := fmt.Sprintf("RapiDoc catalog containing %d API specifications.", len(sources))
		if err := writeCanonicalFile(stage, "overview.md", pageFront{
			Title: "Overview", PageID: "overview", Description: description,
			Source: provenance, SourceType: "openapi-catalog", ImportedFrom: provenance,
		}, openAPICatalogOverview(name, version, sources)); err != nil {
			return err
		}
		pages = 1
		resolvedSources := make(map[string]string, len(sources))
		for index, spec := range sources {
			if err := ctx.Err(); err != nil {
				return err
			}
			raw, resolved, err := reader.readFromOrigin(ctx, spec.URL, nil, spec.Origin)
			if err != nil {
				return fmt.Errorf("fetch RapiDoc catalog specification %s: %w", spec.Namespace, err)
			}
			resolvedURL, err := url.Parse(resolved)
			if err != nil || spec.Origin != "" && httpOrigin(resolvedURL) != spec.Origin {
				return fmt.Errorf("RapiDoc catalog specification %s resolved outside its origin", spec.Namespace)
			}
			canonicalResolved, err := canonicalHTTPURL(resolved)
			if err != nil {
				return fmt.Errorf("canonicalize RapiDoc catalog specification %s: %w", spec.Namespace, err)
			}
			if previous, exists := resolvedSources[canonicalResolved]; exists {
				return fmt.Errorf("RapiDoc catalog specifications %s and %s resolve to the same source %s", previous, spec.Namespace, resolved)
			}
			resolvedSources[canonicalResolved] = spec.Namespace
			document, kind, err := parseCatalogAPIDescription(raw)
			if err != nil {
				return fmt.Errorf("parse RapiDoc catalog specification %s (%s): %w", spec.Namespace, resolved, err)
			}
			document, graphSources, err := resolveExternalOpenAPIRefs(ctx, raw, resolved, kind, document, reader)
			if err != nil {
				return fmt.Errorf("resolve RapiDoc catalog specification %s (%s): %w", spec.Namespace, resolved, err)
			}
			sourceCount += graphSources
			label := strings.TrimSpace(spec.Label)
			if label == "" {
				label = strings.TrimSpace(document.Info.Title)
			}
			if label == "" {
				label = spec.Namespace
			}
			release := strings.TrimSpace(document.Info.Version)
			if release == "" {
				release = version
			}
			prefix := filepath.Join("apis", filepath.FromSlash(spec.Namespace))
			if err := writeAPIDescription(stage, prefix, spec.Namespace, label, release, resolved, kind, document, &pages, len(sources)-index-1, options); err != nil {
				return fmt.Errorf("generate RapiDoc catalog specification %s: %w", spec.Namespace, err)
			}
		}
		metadata := manifest{
			Name: name, Version: version, Description: description, Collections: options.Collections,
			SourceRoot: provenance, SourceType: "openapi-catalog", ImportedFrom: provenance, Sources: sourceCount,
		}
		return writeCanonicalFile(stage, "_index.md", metadata, "This document set was generated from a finite RapiDoc API catalog.")
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind = "openapi-catalog"
	result.Framework = "rapidoc"
	result.Source = provenance
	result.Pages = pages
	result.Sources = sourceCount
	reportProgress(options, Progress{Stage: "completed", Framework: "rapidoc", URL: provenance, Pages: pages})
	return result, nil
}

func writeAPIDescription(stage, prefix, namespace, name, version, provenance, kind string, document apiDescription, pages *int, queued int, options Options) error {
	overviewTitle := "Overview"
	overviewID := "overview"
	overviewPath := ""
	if namespace != "" {
		overviewTitle = name + " Overview"
		overviewID = stableID("api-", namespace)
		overviewPath = filepath.ToSlash(prefix)
	}
	description := strings.TrimSpace(document.Info.Description)
	if err := writeCanonicalFile(stage, filepath.Join(prefix, "overview.md"), pageFront{
		Title: overviewTitle, PageID: overviewID, Path: overviewPath, Description: description,
		Source: provenance, SourceType: kind, ImportedFrom: provenance,
	}, openAPIOverview(name, version, provenance, document)); err != nil {
		return err
	}
	*pages++
	reportProgress(options, Progress{Stage: "page", Framework: kind, URL: provenance, Pages: *pages, Queued: queued})

	if err := writeAPIOperations(stage, prefix, namespace, provenance, kind, document.Paths, false, pages, queued, options); err != nil {
		return err
	}
	if err := writeAPIOperations(stage, prefix, namespace, provenance, kind, document.Webhooks, true, pages, queued, options); err != nil {
		return err
	}

	schemas := document.Components.Schemas
	if kind == "swagger" {
		schemas = document.Definitions
	}
	for _, schemaName := range sortedKeys(schemas) {
		node := schemas[schemaName]
		description := nodeString(node, "description")
		front := pageFront{
			Title: schemaName, PageID: stableID("schema-", strings.TrimSpace(namespace+" "+schemaName)), Path: filepath.ToSlash(filepath.Join(prefix, "schemas")),
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
		if err := writeCanonicalFile(stage, filepath.Join(prefix, "schemas", stableID("", schemaName)+".md"), front, body); err != nil {
			return err
		}
		*pages++
		reportProgress(options, Progress{Stage: "page", Framework: kind, URL: provenance, Pages: *pages, Queued: queued})
	}
	return nil
}

func writeAPIOperations(stage, prefix, namespace, provenance, kind string, items map[string]map[string]yaml.Node, webhook bool, pages *int, queued int, options Options) error {
	for _, endpoint := range sortedKeys(items) {
		pathItem := items[endpoint]
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
			pageIDPrefix, outputRoot, displayEndpoint := "operation-", "operations", endpoint
			identity := strings.TrimSpace(namespace + " " + strings.ToUpper(method) + " " + endpoint)
			filenameIdentity := endpoint
			if webhook {
				pageIDPrefix, outputRoot, displayEndpoint = "webhook-operation-", "webhooks", "webhook:"+endpoint
				identity = strings.TrimSpace(namespace + " webhook " + strings.ToUpper(method) + " " + endpoint)
				filenameIdentity = "webhook " + endpoint
			}
			front := pageFront{
				Title: title, PageID: stableID(pageIDPrefix, identity), Path: filepath.ToSlash(filepath.Join(prefix, outputRoot, tag)),
				Description: description, Source: provenance, HTTPMethods: []string{strings.ToUpper(method)},
				APIEndpoints: []string{displayEndpoint}, SourceType: kind, ImportedFrom: provenance,
			}
			if operation.OperationID != "" {
				front.OperationIDs = []string{operation.OperationID}
			}
			filename := stableID(method+"-", filenameIdentity) + ".md"
			body, err := operationMarkdown(method, displayEndpoint, operation, node, pathItem["parameters"])
			if err != nil {
				return err
			}
			if err := writeCanonicalFile(stage, filepath.Join(prefix, outputRoot, tag, filename), front, body); err != nil {
				return err
			}
			*pages++
			reportProgress(options, Progress{Stage: "page", Framework: kind, URL: provenance, Pages: *pages, Queued: queued})
		}
	}
	return nil
}

func rapidocSources(ctx context.Context, root *htmlNode, base *url.URL, reader *sourceReader) ([]apiSpecSource, bool, error) {
	static := rapidocSpecCandidates(root, base)
	roots := rapidocCatalogRoots(root, base)
	if len(roots) == 0 && len(static) < 2 {
		return nil, false, nil
	}
	sources := staticAPISpecSources(static)
	for _, catalogRoot := range roots {
		expanded, err := expandOVHCatalog(ctx, catalogRoot, reader)
		if err != nil {
			return nil, true, err
		}
		sources = append(sources, expanded...)
	}
	if len(sources) == 0 {
		return nil, true, errors.New("RapiDoc catalog contains no specifications")
	}
	byURL := make(map[string]string, len(sources))
	byNamespace := make(map[string]string, len(sources))
	for _, source := range sources {
		if previous, exists := byURL[source.URL]; exists {
			return nil, true, fmt.Errorf("RapiDoc catalog repeats specification URL %s in %s and %s", source.URL, previous, source.Namespace)
		}
		if previous, exists := byNamespace[source.Namespace]; exists {
			return nil, true, fmt.Errorf("RapiDoc catalog namespace %q is shared by %s and %s", source.Namespace, previous, source.URL)
		}
		byURL[source.URL] = source.Namespace
		byNamespace[source.Namespace] = source.URL
	}
	sort.Slice(sources, func(left, right int) bool { return sources[left].Namespace < sources[right].Namespace })
	return sources, true, nil
}

func staticAPISpecSources(candidates []string) []apiSpecSource {
	sources := make([]apiSpecSource, 0, len(candidates))
	counts := make(map[string]int, len(candidates))
	for _, candidate := range candidates {
		counts[specNamespace(candidate)]++
	}
	for _, candidate := range candidates {
		namespace := specNamespace(candidate)
		if counts[namespace] > 1 {
			namespace += "/" + stableID("source-", candidate)
		}
		sources = append(sources, apiSpecSource{URL: candidate, Namespace: namespace, Label: namespace})
	}
	return sources
}

func rapidocSpecCandidates(root *htmlNode, base *url.URL) []string {
	var candidates []string
	walkHTML(root, func(node *htmlNode) {
		if node.tag == "rapi-doc" || node.tag == "rapi-doc-mini" {
			appendOpenAPIURL(&candidates, base, node.attrs["spec-url"])
		}
	})
	return uniqueStrings(candidates)
}

func hasRapiDocComponent(root *htmlNode) bool {
	found := false
	walkHTML(root, func(node *htmlNode) {
		if node.tag == "rapi-doc" || node.tag == "rapi-doc-mini" {
			found = true
		}
	})
	return found
}

func rapidocCatalogRoots(root *htmlNode, base *url.URL) []string {
	var roots []string
	walkHTML(root, func(node *htmlNode) {
		if node.tag != "rapi-doc" && node.tag != "rapi-doc-mini" {
			return
		}
		for _, value := range strings.Split(node.attrs["spec-roots"], ",") {
			appendOpenAPIURL(&roots, base, value)
		}
	})
	return uniqueStrings(roots)
}

func expandOVHCatalog(ctx context.Context, source string, reader *sourceReader) ([]apiSpecSource, error) {
	raw, provenance, err := reader.read(ctx, source, nil)
	if err != nil {
		return nil, fmt.Errorf("fetch RapiDoc catalog %s: %w", source, err)
	}
	var catalog apiCatalog
	if err := yaml.Unmarshal(raw, &catalog); err != nil {
		return nil, fmt.Errorf("parse RapiDoc catalog %s: %w", provenance, err)
	}
	if len(catalog.APIs) == 0 || strings.TrimSpace(catalog.BasePath) == "" {
		return nil, fmt.Errorf("RapiDoc catalog %s requires non-empty apis and basePath", provenance)
	}
	catalogURL, err := url.Parse(provenance)
	if err != nil {
		return nil, fmt.Errorf("parse RapiDoc catalog URL %s: %w", provenance, err)
	}
	baseURL, err := url.Parse(strings.TrimSpace(catalog.BasePath))
	if err != nil || baseURL.Scheme != "http" && baseURL.Scheme != "https" || baseURL.Host == "" || baseURL.RawQuery != "" || baseURL.Fragment != "" || baseURL.RawPath != "" {
		return nil, fmt.Errorf("RapiDoc catalog %s has invalid HTTP(S) basePath", provenance)
	}
	if !sameHTTPOrigin(catalogURL, baseURL) {
		return nil, fmt.Errorf("RapiDoc catalog %s basePath changes origin to %s", provenance, baseURL.String())
	}
	rootNamespace := SafeSlug(path.Base(strings.TrimRight(catalogURL.Path, "/")))
	if rootNamespace == "" {
		return nil, fmt.Errorf("RapiDoc catalog %s has no stable root namespace", provenance)
	}
	sources := make([]apiSpecSource, 0, len(catalog.APIs))
	for index, entry := range catalog.APIs {
		entryPath := strings.TrimSpace(entry.Path)
		template := strings.TrimSpace(entry.Schema)
		if entryPath == "" || template == "" || !containsFold(entry.Formats, "json") {
			return nil, fmt.Errorf("RapiDoc catalog %s entry %d requires path, schema, and JSON format", provenance, index)
		}
		rendered := strings.ReplaceAll(strings.ReplaceAll(template, "{path}", entryPath), "{format}", "json")
		if strings.ContainsAny(rendered, "{}?#%") || !strings.HasPrefix(rendered, "/") || hasDotPathSegment(rendered) {
			return nil, fmt.Errorf("RapiDoc catalog %s entry %q has unsupported schema template %q", provenance, entryPath, template)
		}
		specURL := *baseURL
		specURL.Path = path.Join(baseURL.Path, rendered)
		specURL.RawPath = ""
		if specURL.Scheme != "http" && specURL.Scheme != "https" || specURL.Host == "" || !sameHTTPOrigin(baseURL, &specURL) {
			return nil, fmt.Errorf("RapiDoc catalog %s entry %q resolves outside its origin", provenance, entryPath)
		}
		query := specURL.Query()
		query.Set("format", "openapi3")
		specURL.RawQuery = query.Encode()
		namespace := rootNamespace + "/" + safeURLPath(entryPath)
		if strings.HasSuffix(namespace, "/") || safeURLPath(entryPath) == "" {
			return nil, fmt.Errorf("RapiDoc catalog %s entry %q has no stable namespace", provenance, entryPath)
		}
		sources = append(sources, apiSpecSource{URL: specURL.String(), Namespace: namespace, Label: entryPath, Origin: httpOrigin(baseURL)})
	}
	return sources, nil
}

func hasDotPathSegment(value string) bool {
	for _, segment := range strings.Split(value, "/") {
		if segment == "." || segment == ".." {
			return true
		}
	}
	return false
}

func sameHTTPOrigin(left, right *url.URL) bool {
	return httpOrigin(left) == httpOrigin(right)
}

func containsFold(values []string, wanted string) bool {
	for _, value := range values {
		if strings.EqualFold(strings.TrimSpace(value), wanted) {
			return true
		}
	}
	return false
}

func safeURLPath(value string) string {
	var segments []string
	for _, segment := range strings.Split(strings.Trim(value, "/"), "/") {
		if slug := SafeSlug(segment); slug != "" {
			segments = append(segments, slug)
		}
	}
	return strings.Join(segments, "/")
}

func specNamespace(source string) string {
	parsed, err := url.Parse(source)
	if err == nil {
		if namespace := safeURLPath(strings.TrimSuffix(parsed.Path, path.Ext(parsed.Path))); namespace != "" {
			return namespace
		}
		if namespace := SafeSlug(parsed.Hostname()); namespace != "" {
			return namespace
		}
	}
	return stableID("spec-", source)
}

func openAPICatalogOverview(name, version string, sources []apiSpecSource) string {
	var body strings.Builder
	fmt.Fprintf(&body, "# %s %s\n\n", name, version)
	fmt.Fprintf(&body, "This catalog contains %d API specifications.\n\n", len(sources))
	body.WriteString("## Specifications\n\n")
	for _, source := range sources {
		fmt.Fprintf(&body, "- `%s`: `%s`\n", source.Namespace, source.URL)
	}
	return body.String()
}

func parseAPIDescription(raw []byte) (apiDescription, string, error) {
	document, kind, err := parseAPIDescriptionDocument(raw)
	if err != nil {
		return apiDescription{}, "", err
	}
	if len(document.Paths) == 0 && len(document.Webhooks) == 0 {
		return apiDescription{}, "", errors.New("OpenAPI document has no paths or webhooks")
	}
	return document, kind, nil
}

func hasAPIDescriptionOperations(document apiDescription) bool {
	for _, items := range []map[string]map[string]yaml.Node{document.Paths, document.Webhooks} {
		for _, pathItem := range items {
			for method := range pathItem {
				if operationMethods[strings.ToLower(method)] {
					return true
				}
			}
		}
	}
	return false
}

func parseCatalogAPIDescription(raw []byte) (apiDescription, string, error) {
	document, kind, err := parseAPIDescriptionDocument(raw)
	if err != nil {
		return apiDescription{}, "", err
	}
	if document.Paths == nil {
		return apiDescription{}, "", errors.New("OpenAPI catalog document has no paths member")
	}
	return document, kind, nil
}

func parseAPIDescriptionDocument(raw []byte) (apiDescription, string, error) {
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
	return document, kind, nil
}

func openAPISpecCandidates(root *htmlNode, base *url.URL) []string {
	var candidates []string
	add := func(value string) {
		appendOpenAPIURL(&candidates, base, value)
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
		for _, candidate := range openAPIConfigCandidates(script+" "+node.attrs["id"]+" "+node.attrs["class"], base) {
			add(candidate)
		}
	})
	return uniqueStrings(candidates)
}

func openAPIConfigScripts(root *htmlNode, base *url.URL) []string {
	var scripts []string
	walkHTML(root, func(node *htmlNode) {
		if node.tag != "script" || node.attrs["src"] == "" {
			return
		}
		identity := strings.ToLower(node.attrs["src"] + " " + node.attrs["id"] + " " + node.attrs["class"])
		if strings.Contains(identity, "swagger-initializer") || strings.Contains(identity, "swagger-config") {
			appendOpenAPIURL(&scripts, base, node.attrs["src"])
		}
	})
	return uniqueStrings(scripts)
}

func openAPIConfigCandidates(script string, base *url.URL) []string {
	lower := strings.ToLower(script)
	var candidates []string
	if strings.Contains(lower, "swaggeruibundle") || strings.Contains(lower, "swaggerui(") || strings.Contains(lower, "swagger-config") {
		for _, match := range scriptConfigURL.FindAllStringSubmatch(script, -1) {
			appendOpenAPIURL(&candidates, base, match[1])
		}
		if len(candidates) == 0 {
			stringsByVariable := make(map[string]string)
			for _, match := range scriptStringVariable.FindAllStringSubmatch(script, -1) {
				stringsByVariable[match[1]] = match[2]
			}
			aliases := make(map[string]string)
			for _, match := range scriptVariableAlias.FindAllStringSubmatch(script, -1) {
				aliases[match[1]] = match[2]
			}
			for _, match := range scriptConfigVariable.FindAllStringSubmatch(script, -1) {
				variable := match[1]
				for range len(aliases) + 1 {
					if value, ok := stringsByVariable[variable]; ok {
						appendOpenAPIURL(&candidates, base, value)
						break
					}
					next, ok := aliases[variable]
					if !ok || next == variable {
						break
					}
					variable = next
				}
			}
		}
		if len(candidates) == 0 && strings.Contains(lower, "location.host") {
			for _, match := range hostConfigURL.FindAllStringSubmatch(script, -1) {
				candidate, err := url.Parse(match[2])
				if err == nil && strings.EqualFold(match[1], base.Hostname()) && likelyOpenAPISpecURL(candidate) {
					appendOpenAPIURL(&candidates, base, match[2])
				}
			}
		}
	}
	if strings.Contains(lower, "redoc") {
		for _, match := range redocInitURL.FindAllStringSubmatch(script, -1) {
			appendOpenAPIURL(&candidates, base, match[1])
		}
		for _, match := range scriptConfigURL.FindAllStringSubmatch(script, -1) {
			appendOpenAPIURL(&candidates, base, match[1])
		}
	}
	return uniqueStrings(candidates)
}

func likelyOpenAPISpecURL(candidate *url.URL) bool {
	path := strings.ToLower(candidate.Path)
	extension := strings.ToLower(filepath.Ext(path))
	return extension == ".json" || extension == ".yaml" || extension == ".yml" ||
		strings.Contains(path, "openapi") || strings.Contains(path, "swagger") || strings.Contains(path, "api-docs")
}

func appendOpenAPIURL(candidates *[]string, base *url.URL, value string) {
	value = strings.TrimSpace(strings.ReplaceAll(value, `\/`, `/`))
	if value == "" || strings.ContainsAny(value, "{}\n\r") {
		return
	}
	resolved, err := base.Parse(value)
	if err != nil || resolved.Host == "" || resolved.Scheme != "http" && resolved.Scheme != "https" {
		return
	}
	resolved.Fragment = ""
	*candidates = append(*candidates, resolved.String())
}

func uniqueStrings(values []string) []string {
	unique := make(map[string]bool)
	output := make([]string, 0, len(values))
	for _, value := range values {
		if !unique[value] {
			unique[value] = true
			output = append(output, value)
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
