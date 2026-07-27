package importer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	maxOpenAPIRefDocuments = 2_000
	maxOpenAPIBundleDepth  = 512
	maxOpenAPIBundleNodes  = 2_000_000
)

type openAPISchemaEntry struct {
	source   string
	name     string
	identity string
	node     *yaml.Node
}

type openAPIRefDocument struct {
	source string
	root   *yaml.Node
}

type openAPIRefGraph struct {
	documents         map[string]*openAPIRefDocument
	aliases           map[string]string
	schemaTargets     map[string]string
	schemaDefinitions map[string]string
	scope             openAPIRefScope
	kind              string
	openAPI31         bool
	bundleNodes       int
}

type openAPIBundleContext uint8

const (
	openAPIDocumentContext openAPIBundleContext = iota
	openAPIComponentsContext
	openAPIReferenceContext
	openAPIReferenceMapContext
	openAPIReferenceListContext
	openAPIParameterContext
	openAPIParameterMapContext
	openAPIParameterListContext
	openAPIHeaderContext
	openAPIHeaderMapContext
	openAPIRequestBodyContext
	openAPIRequestBodyMapContext
	openAPIResponseContext
	openAPIResponseMapContext
	openAPISecuritySchemeContext
	openAPISecuritySchemeMapContext
	openAPIPathItemContext
	openAPISchemaContext
	openAPISchemaMapContext
	openAPISchemaListContext
	openAPIPathItemMapContext
	openAPIOperationContext
	openAPICallbackContext
	openAPICallbackMapContext
	openAPIMediaContext
	openAPIMediaMapContext
	openAPIEncodingContext
	openAPIEncodingMapContext
	openAPIExampleContext
	openAPIExampleMapContext
	openAPILinkContext
	openAPILinkMapContext
	openAPIDataContext
)

type openAPISemanticReference struct {
	value   string
	context openAPIBundleContext
}

type openAPIPendingTraversal struct {
	source   string
	node     *yaml.Node
	context  openAPIBundleContext
	identity string
}

type openAPIRefScope struct {
	remote    bool
	origin    string
	path      string
	localRoot string
}

type openAPIReference struct {
	source   string
	fragment string
}

// resolveExternalOpenAPIRefs loads and bundles an external reference graph. A
// document without external references follows the existing single-file path.
func resolveExternalOpenAPIRefs(ctx context.Context, raw []byte, provenance, kind string, document apiDescription, reader *sourceReader) (apiDescription, int, error) {
	root, _, err := parseOpenAPIRefDocument(raw, provenance)
	if err != nil {
		return apiDescription{}, 0, err
	}
	probe := &openAPIRefGraph{
		kind:      kind,
		openAPI31: kind == "openapi" && strings.HasPrefix(strings.TrimSpace(document.OpenAPI), "3.1"),
	}
	if err := probe.externalSchemaResourceReferenceError(root, openAPIDocumentContext); err != nil {
		return apiDescription{}, 0, fmt.Errorf("validate OpenAPI references in %s: %w", provenance, err)
	}
	refs, err := probe.semanticReferences(root, openAPIDocumentContext)
	if err != nil {
		return apiDescription{}, 0, fmt.Errorf("parse OpenAPI references in %s: %w", provenance, err)
	}
	if len(refs) == 0 {
		return document, 1, nil
	}

	scope, rootSource, err := newOpenAPIRefScope(provenance)
	if err != nil {
		return apiDescription{}, 0, fmt.Errorf("resolve OpenAPI reference scope: %w", err)
	}
	graph := &openAPIRefGraph{
		documents: map[string]*openAPIRefDocument{rootSource: {source: rootSource, root: root}},
		aliases:   map[string]string{rootSource: rootSource},
		scope:     scope,
		kind:      kind,
		openAPI31: kind == "openapi" && strings.HasPrefix(strings.TrimSpace(document.OpenAPI), "3.1"),
	}
	queue := []openAPIPendingTraversal{{source: rootSource, node: root, context: openAPIDocumentContext, identity: openAPIRefIdentity(rootSource, "document")}}
	visited := make(map[string]bool)
	fetchedSources := 1
	hasExternal := false
	var schemaResourceErr error
	for len(queue) > 0 {
		if err := ctx.Err(); err != nil {
			return apiDescription{}, 0, err
		}
		pending := queue[0]
		queue = queue[1:]
		visitKey := pending.identity + "\x00" + strconv.Itoa(int(pending.context))
		if visited[visitKey] {
			continue
		}
		visited[visitKey] = true
		if issue := graph.schemaResourceError(pending.node, pending.context); issue != nil {
			if hasExternal {
				return apiDescription{}, 0, fmt.Errorf("validate OpenAPI references in %s: %w", pending.source, issue)
			}
			if schemaResourceErr == nil {
				schemaResourceErr = fmt.Errorf("validate OpenAPI references in %s: %w", pending.source, issue)
			}
		}
		references, parseErr := graph.semanticReferences(pending.node, pending.context)
		if parseErr != nil {
			return apiDescription{}, 0, fmt.Errorf("parse OpenAPI references in %s: %w", pending.source, parseErr)
		}
		for _, semanticRef := range references {
			external := externalOpenAPIReference(semanticRef.value)
			hasExternal = hasExternal || external
			if issue := graph.schemaReferenceResourceError(semanticRef); issue != nil {
				if external || hasExternal {
					return apiDescription{}, 0, fmt.Errorf("validate OpenAPI references in %s: %w", pending.source, issue)
				}
				if schemaResourceErr == nil {
					schemaResourceErr = fmt.Errorf("validate OpenAPI references in %s: %w", pending.source, issue)
				}
				continue
			}
			if external && schemaResourceErr != nil {
				return apiDescription{}, 0, schemaResourceErr
			}
			ref, resolveErr := graph.scope.resolve(pending.source, semanticRef.value)
			if resolveErr != nil {
				return apiDescription{}, 0, fmt.Errorf("resolve OpenAPI reference %q in %s: %w", semanticRef.value, pending.source, resolveErr)
			}
			if ref.source != pending.source && graph.aliases[ref.source] == "" {
				if fetchedSources >= maxOpenAPIRefDocuments {
					return apiDescription{}, 0, fmt.Errorf("OpenAPI reference graph exceeds %d source URLs", maxOpenAPIRefDocuments)
				}
				fetchedSources++
				referencedRaw, resolved, readErr := graph.scope.read(ctx, ref.source, reader)
				if readErr != nil {
					return apiDescription{}, 0, fmt.Errorf("fetch OpenAPI reference %s: %w", ref.source, readErr)
				}
				resolvedSource, canonicalErr := graph.scope.canonical(resolved)
				if canonicalErr != nil {
					return apiDescription{}, 0, fmt.Errorf("validate OpenAPI reference response %s: %w", resolved, canonicalErr)
				}
				graph.aliases[ref.source] = resolvedSource
				if graph.documents[resolvedSource] == nil {
					referenced, _, parseErr := parseOpenAPIRefDocument(referencedRaw, resolvedSource)
					if parseErr != nil {
						return apiDescription{}, 0, parseErr
					}
					graph.aliases[resolvedSource] = resolvedSource
					graph.documents[resolvedSource] = &openAPIRefDocument{source: resolvedSource, root: referenced}
					queue = append(queue, openAPIPendingTraversal{
						source: resolvedSource, node: referenced, context: openAPIDocumentContext,
						identity: openAPIRefIdentity(resolvedSource, "document"),
					})
				}
			}
			target, targetSource, pointer, resolveErr := graph.resolveReference(pending.source, semanticRef.value)
			if resolveErr != nil {
				return apiDescription{}, 0, fmt.Errorf("resolve OpenAPI reference %q in %s: %w", semanticRef.value, pending.source, resolveErr)
			}
			queue = append(queue, openAPIPendingTraversal{
				source: targetSource, node: target, context: semanticRef.context,
				identity: openAPIRefIdentity(targetSource, pointer),
			})
		}
	}
	if !hasExternal {
		return document, 1, nil
	}
	if schemaResourceErr != nil {
		return apiDescription{}, 0, schemaResourceErr
	}

	if err := graph.indexSchemas(kind); err != nil {
		return apiDescription{}, 0, err
	}
	bundled, err := graph.bundleNode(rootSource, root, make(map[string]bool), openAPIDocumentContext)
	if err != nil {
		return apiDescription{}, 0, fmt.Errorf("bundle OpenAPI reference graph: %w", err)
	}
	var aggregate apiDescription
	if err := bundled.Decode(&aggregate); err != nil {
		return apiDescription{}, 0, fmt.Errorf("decode bundled OpenAPI document: %w", err)
	}
	if err := graph.addSchemas(&aggregate, kind); err != nil {
		return apiDescription{}, 0, err
	}
	return aggregate, len(graph.documents), nil
}

func parseOpenAPIRefDocument(raw []byte, source string) (*yaml.Node, []string, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(raw))
	var document yaml.Node
	if err := decoder.Decode(&document); err != nil {
		return nil, nil, fmt.Errorf("parse referenced OpenAPI document %s: %w", source, err)
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return nil, nil, fmt.Errorf("parse referenced OpenAPI document %s: multiple YAML documents are not supported", source)
		}
		return nil, nil, fmt.Errorf("parse referenced OpenAPI document %s: %w", source, err)
	}
	if document.Kind != yaml.DocumentNode || len(document.Content) != 1 || document.Content[0].Kind != yaml.MappingNode {
		return nil, nil, fmt.Errorf("parse referenced OpenAPI document %s: document root must be a mapping", source)
	}
	_, err := inspectOpenAPIYAMLNode(&document)
	if err != nil {
		return nil, nil, fmt.Errorf("parse referenced OpenAPI document %s: %w", source, err)
	}
	return &document, nil, nil
}

func inspectOpenAPIYAMLNode(root *yaml.Node) (int, error) {
	stack := []*yaml.Node{root}
	seen := make(map[*yaml.Node]bool)
	nodes := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil || seen[node] {
			continue
		}
		seen[node] = true
		nodes++
		if node.Kind == yaml.AliasNode {
			stack = append(stack, node.Alias)
			continue
		}
		if node.Kind == yaml.MappingNode {
			keys := make(map[string]bool, len(node.Content)/2)
			for index := 0; index+1 < len(node.Content); index += 2 {
				key := node.Content[index]
				if key.Kind != yaml.ScalarNode {
					return nodes, errors.New("mapping keys must be scalar")
				}
				identity := key.Tag + "\x00" + key.Value
				if keys[identity] {
					return nodes, fmt.Errorf("duplicate mapping key %q", key.Value)
				}
				keys[identity] = true
			}
		}
		stack = append(stack, node.Content...)
	}
	return nodes, nil
}

func (graph *openAPIRefGraph) semanticReferences(root *yaml.Node, context openAPIBundleContext) ([]openAPISemanticReference, error) {
	type pendingNode struct {
		node    *yaml.Node
		context openAPIBundleContext
	}
	stack := []pendingNode{{node: root, context: context}}
	seen := make(map[pendingNode]bool)
	var references []openAPISemanticReference
	for len(stack) > 0 {
		pending := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pending.node == nil || pending.context == openAPIDataContext || seen[pending] {
			continue
		}
		seen[pending] = true
		if pending.node.Kind == yaml.AliasNode {
			stack = append(stack, pendingNode{node: pending.node.Alias, context: pending.context})
			continue
		}
		if pending.node.Kind == yaml.MappingNode {
			refIndex := -1
			if graph.contextAllowsReference(pending.context) {
				refIndex = yamlRefIndex(pending.node)
				if refIndex >= 0 {
					value := pending.node.Content[refIndex+1]
					if value.Kind != yaml.ScalarNode || value.Tag != "!!str" || strings.TrimSpace(value.Value) == "" {
						return nil, errors.New("$ref must be a non-empty string")
					}
					references = append(references, openAPISemanticReference{value: strings.TrimSpace(value.Value), context: graph.referenceTargetContext(pending.context)})
				}
			}
			for index := 0; index+1 < len(pending.node.Content); index += 2 {
				key := pending.node.Content[index].Value
				if key == "$ref" && refIndex >= 0 || refIndex >= 0 && !graph.referenceSiblingAllowed(pending.context, key) {
					continue
				}
				childContext := graph.childContext(pending.context, key)
				if childContext != openAPIDataContext {
					stack = append(stack, pendingNode{node: pending.node.Content[index+1], context: childContext})
				}
			}
			continue
		}
		childContext := graph.sequenceChildContext(pending.context)
		if pending.node.Kind == yaml.DocumentNode {
			childContext = pending.context
		}
		if childContext == openAPIDataContext {
			continue
		}
		for _, child := range pending.node.Content {
			stack = append(stack, pendingNode{node: child, context: childContext})
		}
	}
	return references, nil
}

func (graph *openAPIRefGraph) schemaResourceError(root *yaml.Node, context openAPIBundleContext) error {
	if !graph.openAPI31 {
		return nil
	}
	type pendingNode struct {
		node    *yaml.Node
		context openAPIBundleContext
	}
	stack := []pendingNode{{node: root, context: context}}
	seen := make(map[pendingNode]bool)
	for len(stack) > 0 {
		pending := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pending.node == nil || pending.context == openAPIDataContext || seen[pending] {
			continue
		}
		seen[pending] = true
		if pending.node.Kind == yaml.AliasNode {
			stack = append(stack, pendingNode{node: pending.node.Alias, context: pending.context})
			continue
		}
		if pending.node.Kind == yaml.MappingNode {
			refIndex := -1
			if graph.contextAllowsReference(pending.context) {
				refIndex = yamlRefIndex(pending.node)
			}
			for index := 0; index+1 < len(pending.node.Content); index += 2 {
				key := pending.node.Content[index].Value
				if pending.context == openAPISchemaContext {
					switch key {
					case "$id", "$anchor", "$dynamicAnchor", "$dynamicRef", "$recursiveAnchor", "$recursiveRef":
						return fmt.Errorf("OpenAPI 3.1 external reference graphs do not support schema resource keyword %q", key)
					}
				}
				if key == "$ref" && refIndex >= 0 || refIndex >= 0 && !graph.referenceSiblingAllowed(pending.context, key) {
					continue
				}
				childContext := graph.childContext(pending.context, key)
				if childContext != openAPIDataContext {
					stack = append(stack, pendingNode{node: pending.node.Content[index+1], context: childContext})
				}
			}
			continue
		}
		childContext := graph.sequenceChildContext(pending.context)
		if pending.node.Kind == yaml.DocumentNode {
			childContext = pending.context
		}
		if childContext != openAPIDataContext {
			for _, child := range pending.node.Content {
				stack = append(stack, pendingNode{node: child, context: childContext})
			}
		}
	}
	return nil
}

func (graph *openAPIRefGraph) externalSchemaResourceReferenceError(root *yaml.Node, context openAPIBundleContext) error {
	if !graph.openAPI31 {
		return nil
	}
	type pendingNode struct {
		node    *yaml.Node
		context openAPIBundleContext
	}
	stack := []pendingNode{{node: root, context: context}}
	seen := make(map[pendingNode]bool)
	for len(stack) > 0 {
		pending := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pending.node == nil || pending.context == openAPIDataContext || seen[pending] {
			continue
		}
		seen[pending] = true
		if pending.node.Kind == yaml.AliasNode {
			stack = append(stack, pendingNode{node: pending.node.Alias, context: pending.context})
			continue
		}
		if pending.node.Kind == yaml.MappingNode {
			refIndex := -1
			if graph.contextAllowsReference(pending.context) {
				refIndex = yamlRefIndex(pending.node)
			}
			for index := 0; index+1 < len(pending.node.Content); index += 2 {
				key, value := pending.node.Content[index].Value, pending.node.Content[index+1]
				if pending.context == openAPISchemaContext && (key == "$dynamicRef" || key == "$recursiveRef") && value.Kind == yaml.ScalarNode && value.Tag == "!!str" && externalOpenAPIReference(value.Value) {
					return fmt.Errorf("OpenAPI 3.1 external reference graphs do not support external schema resource reference %s: %q", key, value.Value)
				}
				if key == "$ref" && refIndex >= 0 || refIndex >= 0 && !graph.referenceSiblingAllowed(pending.context, key) {
					continue
				}
				childContext := graph.childContext(pending.context, key)
				if childContext != openAPIDataContext {
					stack = append(stack, pendingNode{node: value, context: childContext})
				}
			}
			continue
		}
		childContext := graph.sequenceChildContext(pending.context)
		if pending.node.Kind == yaml.DocumentNode {
			childContext = pending.context
		}
		if childContext != openAPIDataContext {
			for _, child := range pending.node.Content {
				stack = append(stack, pendingNode{node: child, context: childContext})
			}
		}
	}
	return nil
}

func (graph *openAPIRefGraph) schemaReferenceResourceError(ref openAPISemanticReference) error {
	if !graph.openAPI31 || ref.context != openAPISchemaContext {
		return nil
	}
	parsed, err := url.Parse(strings.TrimSpace(ref.value))
	if err != nil || parsed.Fragment == "" || strings.HasPrefix(parsed.Fragment, "/") {
		return nil
	}
	return fmt.Errorf("OpenAPI 3.1 external reference graphs do not support non-pointer schema fragment %q in $ref %q", parsed.Fragment, ref.value)
}

func (graph *openAPIRefGraph) contextAllowsReference(context openAPIBundleContext) bool {
	switch context {
	case openAPIReferenceContext, openAPIParameterContext, openAPIHeaderContext, openAPIRequestBodyContext, openAPIResponseContext, openAPISecuritySchemeContext,
		openAPIPathItemContext, openAPICallbackContext, openAPIExampleContext, openAPILinkContext, openAPISchemaContext:
		return true
	default:
		return false
	}
}

func (graph *openAPIRefGraph) referenceTargetContext(context openAPIBundleContext) openAPIBundleContext {
	return context
}

func (graph *openAPIRefGraph) childContext(parent openAPIBundleContext, key string) openAPIBundleContext {
	if strings.HasPrefix(strings.ToLower(key), "x-") {
		return openAPIDataContext
	}
	switch parent {
	case openAPIDocumentContext:
		switch key {
		case "paths", "webhooks":
			return openAPIPathItemMapContext
		case "components":
			return openAPIComponentsContext
		case "definitions":
			return openAPISchemaMapContext
		case "parameters":
			return openAPIParameterMapContext
		case "responses":
			return openAPIResponseMapContext
		}
	case openAPIComponentsContext:
		switch key {
		case "schemas":
			return openAPISchemaMapContext
		case "callbacks":
			return openAPICallbackMapContext
		case "examples":
			return openAPIExampleMapContext
		case "links":
			return openAPILinkMapContext
		case "pathItems":
			return openAPIPathItemMapContext
		case "responses":
			return openAPIResponseMapContext
		case "parameters":
			return openAPIParameterMapContext
		case "requestBodies":
			return openAPIRequestBodyMapContext
		case "headers":
			return openAPIHeaderMapContext
		case "securitySchemes":
			return openAPISecuritySchemeMapContext
		}
	case openAPIReferenceMapContext:
		return openAPIReferenceContext
	case openAPIParameterMapContext:
		return openAPIParameterContext
	case openAPIHeaderMapContext:
		return openAPIHeaderContext
	case openAPIRequestBodyMapContext:
		return openAPIRequestBodyContext
	case openAPIResponseMapContext:
		return openAPIResponseContext
	case openAPISecuritySchemeMapContext:
		return openAPISecuritySchemeContext
	case openAPIPathItemMapContext:
		return openAPIPathItemContext
	case openAPICallbackMapContext:
		return openAPICallbackContext
	case openAPIExampleMapContext:
		return openAPIExampleContext
	case openAPILinkMapContext:
		return openAPILinkContext
	case openAPISchemaMapContext:
		return openAPISchemaContext
	case openAPIPathItemContext:
		if operationMethods[strings.ToLower(key)] {
			return openAPIOperationContext
		}
		if key == "parameters" {
			return openAPIParameterListContext
		}
	case openAPIOperationContext:
		switch key {
		case "parameters":
			return openAPIParameterListContext
		case "requestBody":
			return openAPIRequestBodyContext
		case "responses":
			return openAPIResponseMapContext
		case "callbacks":
			return openAPICallbackMapContext
		}
	case openAPICallbackContext:
		return openAPIPathItemContext
	case openAPIReferenceContext, openAPIParameterContext, openAPIHeaderContext, openAPIRequestBodyContext, openAPIResponseContext:
		switch key {
		case "schema":
			return openAPISchemaContext
		case "content":
			return openAPIMediaMapContext
		case "headers":
			return openAPIHeaderMapContext
		case "links":
			return openAPILinkMapContext
		case "examples":
			if graph.kind != "swagger" {
				return openAPIExampleMapContext
			}
		case "example", "value", "default", "enum", "const":
			return openAPIDataContext
		}
	case openAPIMediaMapContext:
		return openAPIMediaContext
	case openAPIMediaContext:
		switch key {
		case "schema":
			return openAPISchemaContext
		case "examples":
			return openAPIExampleMapContext
		case "example":
			return openAPIDataContext
		case "encoding":
			return openAPIEncodingMapContext
		}
	case openAPIEncodingMapContext:
		return openAPIEncodingContext
	case openAPIEncodingContext:
		if key == "headers" {
			return openAPIHeaderMapContext
		}
	case openAPIExampleContext, openAPILinkContext:
		return openAPIDataContext
	case openAPISchemaContext:
		switch key {
		case "not", "items", "additionalProperties", "unevaluatedProperties", "unevaluatedItems", "contains", "propertyNames", "contentSchema", "if", "then", "else":
			return openAPISchemaContext
		case "properties", "patternProperties", "dependentSchemas", "$defs", "definitions":
			return openAPISchemaMapContext
		case "allOf", "anyOf", "oneOf", "prefixItems":
			return openAPISchemaListContext
		case "example", "examples", "default", "enum", "const":
			return openAPIDataContext
		}
	}
	return openAPIDataContext
}

func (graph *openAPIRefGraph) sequenceChildContext(parent openAPIBundleContext) openAPIBundleContext {
	switch parent {
	case openAPIReferenceListContext:
		return openAPIReferenceContext
	case openAPIParameterListContext:
		return openAPIParameterContext
	case openAPISchemaListContext:
		return openAPISchemaContext
	default:
		return openAPIDataContext
	}
}

func externalOpenAPIReference(value string) bool {
	parsed, err := url.Parse(strings.TrimSpace(value))
	return err != nil || parsed.Scheme != "" || parsed.Host != "" || parsed.Path != "" || parsed.RawQuery != ""
}

func newOpenAPIRefScope(provenance string) (openAPIRefScope, string, error) {
	parsed, err := url.Parse(provenance)
	if err == nil && (parsed.Scheme == "http" || parsed.Scheme == "https") {
		if parsed.User != nil || parsed.Host == "" {
			return openAPIRefScope{}, "", errors.New("HTTP source URL must not contain credentials")
		}
		parsed.Fragment = ""
		parsed.RawFragment = ""
		parsed.RawPath = ""
		canonical, err := canonicalHTTPURL(parsed.String())
		if err != nil {
			return openAPIRefScope{}, "", err
		}
		parsed, _ = url.Parse(canonical)
		scopePath := path.Dir(parsed.Path)
		if scopePath != "/" {
			scopePath += "/"
		}
		return openAPIRefScope{remote: true, origin: httpOrigin(parsed), path: scopePath}, canonical, nil
	}

	absolute, err := filepath.Abs(provenance)
	if err != nil {
		return openAPIRefScope{}, "", err
	}
	root, err := filepath.EvalSymlinks(filepath.Dir(absolute))
	if err != nil {
		return openAPIRefScope{}, "", err
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err != nil {
		return openAPIRefScope{}, "", err
	}
	return openAPIRefScope{localRoot: root}, resolved, nil
}

func (scope openAPIRefScope) resolve(current, value string) (openAPIReference, error) {
	parsed, err := url.Parse(strings.TrimSpace(value))
	if err != nil {
		return openAPIReference{}, fmt.Errorf("invalid URI reference: %w", err)
	}
	if parsed.User != nil {
		return openAPIReference{}, errors.New("reference URL must not contain credentials")
	}
	fragment := parsed.Fragment
	parsed.Fragment = ""
	parsed.RawFragment = ""
	if scope.remote {
		base, _ := url.Parse(current)
		resolved := base.ResolveReference(parsed)
		resolved.RawPath = ""
		canonical, err := scope.canonical(resolved.String())
		if err != nil {
			return openAPIReference{}, err
		}
		return openAPIReference{source: canonical, fragment: fragment}, nil
	}
	if parsed.Scheme != "" || parsed.Host != "" || parsed.RawQuery != "" {
		return openAPIReference{}, errors.New("local OpenAPI references must be local files without query parameters")
	}
	target := current
	if parsed.Path != "" {
		if filepath.IsAbs(filepath.FromSlash(parsed.Path)) {
			target = filepath.FromSlash(parsed.Path)
		} else {
			target = filepath.Join(filepath.Dir(current), filepath.FromSlash(parsed.Path))
		}
	}
	canonical, err := scope.canonical(target)
	if err != nil {
		return openAPIReference{}, err
	}
	return openAPIReference{source: canonical, fragment: fragment}, nil
}

func (scope openAPIRefScope) canonical(value string) (string, error) {
	if scope.remote {
		canonical, err := canonicalHTTPURL(value)
		if err != nil {
			return "", err
		}
		parsed, _ := url.Parse(canonical)
		if parsed.User != nil {
			return "", errors.New("reference URL must not contain credentials")
		}
		if httpOrigin(parsed) != scope.origin {
			return "", fmt.Errorf("reference changes origin from %s to %s", scope.origin, httpOrigin(parsed))
		}
		cleanPath := path.Clean(parsed.Path)
		if scope.path != "/" && !strings.HasPrefix(cleanPath, scope.path) {
			return "", fmt.Errorf("reference path %s is outside specification scope %s", cleanPath, scope.path)
		}
		return canonical, nil
	}
	absolute, err := filepath.Abs(value)
	if err != nil {
		return "", err
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err != nil {
		return "", err
	}
	relative, err := filepath.Rel(scope.localRoot, resolved)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("reference path %s is outside specification scope %s", resolved, scope.localRoot)
	}
	return resolved, nil
}

func (scope openAPIRefScope) read(ctx context.Context, source string, reader *sourceReader) ([]byte, string, error) {
	if !scope.remote {
		return reader.read(ctx, source, nil)
	}
	client := *reader.client
	originalRedirect := client.CheckRedirect
	client.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if request.URL.User != nil {
			return errors.New("OpenAPI reference redirect URL must not contain credentials")
		}
		if _, err := scope.canonical(request.URL.String()); err != nil {
			return fmt.Errorf("OpenAPI reference redirect leaves specification scope: %w", err)
		}
		if originalRedirect != nil {
			return originalRedirect(request, via)
		}
		return nil
	}
	scopedReader := *reader
	scopedReader.client = &client
	raw, resolved, err := scopedReader.readFromOrigin(ctx, source, nil, scope.origin)
	reader.used = scopedReader.used
	return raw, resolved, err
}

func (graph *openAPIRefGraph) canonicalSource(source string) string {
	if resolved := graph.aliases[source]; resolved != "" {
		return resolved
	}
	return source
}

func (graph *openAPIRefGraph) resolveReference(current, value string) (*yaml.Node, string, string, error) {
	ref, err := graph.scope.resolve(current, value)
	if err != nil {
		return nil, "", "", err
	}
	ref.source = graph.canonicalSource(ref.source)
	document := graph.documents[ref.source]
	if document == nil {
		return nil, "", "", fmt.Errorf("referenced document %s was not loaded", ref.source)
	}
	target, pointer, err := openAPIJSONPointer(document.root, ref.fragment)
	if err != nil {
		return nil, "", "", fmt.Errorf("reference %q in %s: %w", value, current, err)
	}
	return target, ref.source, pointer, nil
}

func openAPIJSONPointer(document *yaml.Node, fragment string) (*yaml.Node, string, error) {
	current := document
	if current.Kind == yaml.DocumentNode {
		if len(current.Content) != 1 {
			return nil, "", errors.New("invalid YAML document")
		}
		current = current.Content[0]
	}
	if fragment == "" {
		return current, "", nil
	}
	if !strings.HasPrefix(fragment, "/") {
		return nil, "", fmt.Errorf("fragment %q is not a JSON pointer", fragment)
	}
	segments := strings.Split(fragment[1:], "/")
	decoded := make([]string, 0, len(segments))
	for _, segment := range segments {
		value, err := decodeJSONPointerSegment(segment)
		if err != nil {
			return nil, "", err
		}
		decoded = append(decoded, value)
		if current.Kind == yaml.AliasNode {
			current = current.Alias
		}
		switch current.Kind {
		case yaml.MappingNode:
			var next *yaml.Node
			for index := 0; index+1 < len(current.Content); index += 2 {
				if current.Content[index].Value == value {
					next = current.Content[index+1]
					break
				}
			}
			if next == nil {
				return nil, "", fmt.Errorf("JSON pointer %q does not exist", fragment)
			}
			current = next
		case yaml.SequenceNode:
			index, err := strconv.Atoi(value)
			if err != nil || index < 0 || index >= len(current.Content) || strconv.Itoa(index) != value {
				return nil, "", fmt.Errorf("JSON pointer %q has invalid array index %q", fragment, value)
			}
			current = current.Content[index]
		default:
			return nil, "", fmt.Errorf("JSON pointer %q traverses a scalar", fragment)
		}
	}
	normalized := ""
	for _, segment := range decoded {
		normalized += "/" + encodeJSONPointerSegment(segment)
	}
	return current, normalized, nil
}

func decodeJSONPointerSegment(value string) (string, error) {
	var output strings.Builder
	for index := 0; index < len(value); index++ {
		if value[index] != '~' {
			output.WriteByte(value[index])
			continue
		}
		if index+1 >= len(value) || value[index+1] != '0' && value[index+1] != '1' {
			return "", fmt.Errorf("invalid JSON pointer escape in %q", value)
		}
		index++
		if value[index] == '0' {
			output.WriteByte('~')
		} else {
			output.WriteByte('/')
		}
	}
	return output.String(), nil
}

func encodeJSONPointerSegment(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "~", "~0"), "/", "~1")
}

func (graph *openAPIRefGraph) indexSchemas(kind string) error {
	var schemas []openAPISchemaEntry
	for source, document := range graph.documents {
		var partial apiDescription
		if err := document.root.Decode(&partial); err != nil {
			return fmt.Errorf("decode referenced OpenAPI document %s: %w", source, err)
		}
		values := partial.Components.Schemas
		if kind == "swagger" {
			values = partial.Definitions
		}
		for name := range values {
			prefix := "/components/schemas/"
			if kind == "swagger" {
				prefix = "/definitions/"
			}
			pointer := prefix + encodeJSONPointerSegment(name)
			node, _, err := openAPIJSONPointer(document.root, pointer)
			if err != nil {
				return err
			}
			schemas = append(schemas, openAPISchemaEntry{source: source, name: name, identity: openAPIRefIdentity(source, pointer), node: node})
		}
	}
	sort.Slice(schemas, func(left, right int) bool {
		if schemas[left].source == schemas[right].source {
			return schemas[left].name < schemas[right].name
		}
		return schemas[left].source < schemas[right].source
	})
	reachable, err := graph.externalRootSchemas()
	if err != nil {
		return err
	}
	firstLevel := make(map[string]openAPISchemaEntry, len(schemas))
	for _, schema := range schemas {
		firstLevel[schema.identity] = schema
	}
	meaningfulReachable := make(map[string]bool, len(reachable))
	for identity := range reachable {
		node, err := graph.schemaNode(identity)
		if err != nil {
			return err
		}
		if graph.recognizedSchemaNode(node) {
			meaningfulReachable[identity] = true
		}
	}
	emitted := make(map[string]bool, len(schemas)+len(meaningfulReachable))
	for _, schema := range schemas {
		if graph.recognizedSchemaNode(schema.node) {
			emitted[schema.identity] = true
		}
	}
	for identity := range meaningfulReachable {
		emitted[identity] = true
	}
	aliasTargets := make(map[string]string)
	for identity := range emitted {
		node, err := graph.schemaNode(identity)
		if err != nil {
			return err
		}
		source, _ := splitOpenAPIRefIdentity(identity)
		if target, ok, err := graph.collapsibleSchemaAlias(source, node); err != nil {
			return err
		} else if ok && target != identity && emitted[target] {
			aliasTargets[identity] = target
		}
	}
	canonical := make(map[string]string, len(emitted))
	for identity := range emitted {
		current := identity
		seen := make(map[string]bool)
		for aliasTargets[current] != "" && !seen[current] {
			seen[current] = true
			current = aliasTargets[current]
		}
		canonical[identity] = current
	}
	canonicalSet := make(map[string]bool)
	for _, identity := range canonical {
		canonicalSet[identity] = true
	}
	rootAliasNames := make(map[string][]string)
	for alias, target := range aliasTargets {
		_, pointer := splitOpenAPIRefIdentity(target)
		if pointer == "" {
			if entry, ok := firstLevel[alias]; ok {
				rootAliasNames[target] = append(rootAliasNames[target], entry.name)
			}
		}
	}
	baseNames := make(map[string]string, len(canonicalSet))
	for identity := range canonicalSet {
		if aliases := rootAliasNames[identity]; len(aliases) > 0 {
			sort.Strings(aliases)
			baseNames[identity] = aliases[0]
			continue
		}
		source, pointer := splitOpenAPIRefIdentity(identity)
		baseNames[identity] = openAPISchemaLogicalName(source, pointer)
	}
	names := deterministicSchemaNames(baseNames)
	graph.schemaTargets = make(map[string]string, len(emitted))
	graph.schemaDefinitions = make(map[string]string, len(canonicalSet))
	for identity, target := range canonical {
		graph.schemaTargets[identity] = names[target]
	}
	for identity := range canonicalSet {
		graph.schemaTargets[identity] = names[identity]
		graph.schemaDefinitions[names[identity]] = identity
	}
	return nil
}

func (graph *openAPIRefGraph) collapsibleSchemaAlias(source string, node *yaml.Node) (string, bool, error) {
	refIndex := yamlRefIndex(node)
	if refIndex < 0 || graph.openAPI31 && len(node.Content) != 2 {
		return "", false, nil
	}
	_, targetSource, pointer, err := graph.resolveReference(source, strings.TrimSpace(node.Content[refIndex+1].Value))
	if err != nil {
		return "", false, err
	}
	return openAPIRefIdentity(targetSource, pointer), true, nil
}

func (graph *openAPIRefGraph) externalRootSchemas() (map[string]bool, error) {
	queue := make([]openAPIPendingTraversal, 0, len(graph.documents))
	for source, document := range graph.documents {
		queue = append(queue, openAPIPendingTraversal{
			source: source, node: document.root, context: openAPIDocumentContext,
			identity: openAPIRefIdentity(source, "document"),
		})
	}
	targets := make(map[string]bool)
	seen := make(map[string]bool)
	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		visitKey := current.identity + "\x00" + strconv.Itoa(int(current.context))
		if current.node == nil || seen[visitKey] {
			continue
		}
		seen[visitKey] = true
		references, err := graph.semanticReferences(current.node, current.context)
		if err != nil {
			return nil, err
		}
		for _, ref := range references {
			target, targetSource, pointer, err := graph.resolveReference(current.source, ref.value)
			if err != nil {
				return nil, err
			}
			identity := openAPIRefIdentity(targetSource, pointer)
			if ref.context == openAPISchemaContext {
				targets[identity] = true
			}
			queue = append(queue, openAPIPendingTraversal{
				source: targetSource, node: target, context: ref.context, identity: identity,
			})
		}
	}
	return targets, nil
}

func openAPISchemaNamespace(source string) string {
	parsed, err := url.Parse(source)
	value := source
	if err == nil && parsed.Path != "" {
		value = parsed.Path
	}
	base := path.Base(strings.TrimSuffix(value, path.Ext(value)))
	if namespace := SafeSlug(base); namespace != "" {
		return namespace
	}
	return stableID("source-", source)
}

func (graph *openAPIRefGraph) schemaNode(identity string) (*yaml.Node, error) {
	source, pointer := splitOpenAPIRefIdentity(identity)
	document := graph.documents[source]
	if document == nil {
		return nil, fmt.Errorf("schema source %s was not loaded", source)
	}
	node, _, err := openAPIJSONPointer(document.root, pointer)
	return node, err
}

func (graph *openAPIRefGraph) recognizedSchemaNode(node *yaml.Node) bool {
	seen := make(map[*yaml.Node]bool)
	for node != nil && node.Kind == yaml.AliasNode {
		if seen[node] {
			return false
		}
		seen[node] = true
		node = node.Alias
	}
	if node == nil {
		return false
	}
	if graph.openAPI31 && node.Kind == yaml.ScalarNode && (node.Tag == "!!bool" || node.Value == "true" || node.Value == "false") {
		return true
	}
	if node.Kind != yaml.MappingNode {
		return false
	}
	if len(node.Content) == 0 {
		return true
	}
	for index := 0; index+1 < len(node.Content); index += 2 {
		key := node.Content[index].Value
		if recognizedOpenAPISchemaKeywords[key] || strings.HasPrefix(strings.ToLower(key), "x-") {
			return true
		}
	}
	return false
}

var recognizedOpenAPISchemaKeywords = map[string]bool{
	"$anchor": true, "$comment": true, "$defs": true, "$dynamicAnchor": true, "$dynamicRef": true, "$id": true, "$recursiveAnchor": true, "$recursiveRef": true, "$ref": true, "$schema": true,
	"additionalItems": true, "additionalProperties": true, "allOf": true, "anyOf": true,
	"const": true, "contains": true, "contentEncoding": true, "contentMediaType": true, "contentSchema": true,
	"default": true, "dependentRequired": true, "dependentSchemas": true, "deprecated": true, "description": true, "discriminator": true,
	"else": true, "enum": true, "example": true, "examples": true, "exclusiveMaximum": true, "exclusiveMinimum": true, "externalDocs": true,
	"format": true, "if": true, "items": true,
	"maxContains": true, "maximum": true, "maxItems": true, "maxLength": true, "maxProperties": true,
	"minContains": true, "minimum": true, "minItems": true, "minLength": true, "minProperties": true, "multipleOf": true,
	"not": true, "nullable": true, "oneOf": true, "pattern": true, "patternProperties": true, "prefixItems": true, "properties": true, "propertyNames": true,
	"readOnly": true, "required": true, "then": true, "title": true, "type": true,
	"unevaluatedItems": true, "unevaluatedProperties": true, "uniqueItems": true,
	"writeOnly": true, "xml": true,
}

func openAPISchemaLogicalName(source, pointer string) string {
	segments := openAPIJSONPointerSegments(pointer)
	prefixSource := true
	if len(segments) >= 2 && segments[0] == "components" && segments[1] == "schemas" {
		segments = segments[2:]
		prefixSource = false
	} else if len(segments) >= 1 && segments[0] == "definitions" {
		segments = segments[1:]
		prefixSource = false
	} else if len(segments) >= 1 && segments[0] == "components" {
		segments = segments[1:]
	}
	for index := range segments {
		segments[index] = strings.ReplaceAll(strings.TrimSpace(segments[index]), "/", ".")
	}
	if prefixSource || len(segments) == 0 {
		segments = append([]string{openAPISchemaNamespace(source)}, segments...)
	}
	return strings.Join(segments, ".")
}

func openAPIJSONPointerSegments(pointer string) []string {
	if pointer == "" {
		return nil
	}
	encoded := strings.Split(strings.TrimPrefix(pointer, "/"), "/")
	segments := make([]string, 0, len(encoded))
	for _, segment := range encoded {
		decoded, err := decodeJSONPointerSegment(segment)
		if err == nil {
			segments = append(segments, decoded)
		}
	}
	return segments
}

func deterministicSchemaNames(baseNames map[string]string) map[string]string {
	counts := make(map[string]int, len(baseNames))
	for _, name := range baseNames {
		counts[name]++
	}
	identities := make([]string, 0, len(baseNames))
	for identity := range baseNames {
		identities = append(identities, identity)
	}
	sort.Strings(identities)
	used := make(map[string]bool, len(identities))
	names := make(map[string]string, len(identities))
	for _, identity := range identities {
		name := baseNames[identity]
		source, _ := splitOpenAPIRefIdentity(identity)
		if counts[name] > 1 && !strings.HasPrefix(name, openAPISchemaNamespace(source)+".") {
			name = openAPISchemaNamespace(source) + "." + name
		}
		if used[name] {
			name = openAPISchemaSourceNamespace(source) + "." + name
		}
		base := name
		for suffix := 2; used[name]; suffix++ {
			name = base + "." + strconv.Itoa(suffix)
		}
		used[name] = true
		names[identity] = name
	}
	return names
}

func openAPISchemaSourceNamespace(source string) string {
	parsed, err := url.Parse(source)
	value := source
	if err == nil && parsed.Path != "" {
		value = strings.TrimSuffix(parsed.Path, path.Ext(parsed.Path))
	}
	if namespace := safeURLPath(value); namespace != "" {
		return strings.ReplaceAll(namespace, "/", ".")
	}
	return openAPISchemaNamespace(source)
}

func (graph *openAPIRefGraph) addSchemas(document *apiDescription, kind string) error {
	if kind == "swagger" {
		document.Definitions = make(map[string]yaml.Node, len(graph.schemaDefinitions))
	} else {
		document.Components.Schemas = make(map[string]yaml.Node, len(graph.schemaDefinitions))
	}
	names := make([]string, 0, len(graph.schemaDefinitions))
	for name := range graph.schemaDefinitions {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		source, pointer := splitOpenAPIRefIdentity(graph.schemaDefinitions[name])
		target, _, err := openAPIJSONPointer(graph.documents[source].root, pointer)
		if err != nil {
			return err
		}
		bundled, err := graph.bundleNode(source, target, make(map[string]bool), openAPISchemaContext)
		if err != nil {
			return fmt.Errorf("bundle schema %s: %w", name, err)
		}
		if kind == "swagger" {
			document.Definitions[name] = *bundled
		} else {
			document.Components.Schemas[name] = *bundled
		}
	}
	return nil
}

func (graph *openAPIRefGraph) bundleNode(source string, node *yaml.Node, resolving map[string]bool, context openAPIBundleContext) (*yaml.Node, error) {
	return graph.bundleNodeAtDepth(source, node, resolving, make(map[*yaml.Node]bool), context, 0)
}

func (graph *openAPIRefGraph) bundleNodeAtDepth(source string, node *yaml.Node, resolving map[string]bool, aliases map[*yaml.Node]bool, context openAPIBundleContext, depth int) (*yaml.Node, error) {
	if node == nil {
		return nil, errors.New("nil referenced node")
	}
	if depth > maxOpenAPIBundleDepth {
		return nil, fmt.Errorf("bundled OpenAPI reference exceeds depth %d", maxOpenAPIBundleDepth)
	}
	if err := graph.consumeBundleNode(); err != nil {
		return nil, err
	}
	if node.Kind == yaml.AliasNode {
		if aliases[node] {
			return nil, errors.New("cyclic YAML alias in OpenAPI document")
		}
		aliases[node] = true
		bundled, err := graph.bundleNodeAtDepth(source, node.Alias, resolving, aliases, context, depth+1)
		delete(aliases, node)
		return bundled, err
	}
	if context == openAPIDataContext {
		return graph.cloneYAMLNodeBounded(node, depth, false, make(map[*yaml.Node]bool))
	}
	if node.Kind == yaml.MappingNode && graph.contextAllowsReference(context) {
		refIndex := yamlRefIndex(node)
		if refIndex >= 0 {
			value := strings.TrimSpace(node.Content[refIndex+1].Value)
			target, targetSource, pointer, err := graph.resolveReference(source, value)
			if err != nil {
				return nil, err
			}
			identity := openAPIRefIdentity(targetSource, pointer)
			if schemaName := graph.schemaTargets[identity]; schemaName != "" {
				clone := *node
				clone.Content = nil
				for index := 0; index+1 < len(node.Content); index += 2 {
					if node.Content[index].Value == "$ref" {
						prefix := "#/components/schemas/"
						if graph.kind == "swagger" {
							prefix = "#/definitions/"
						}
						refValue, err := graph.cloneYAMLScalarBounded(node.Content[index+1])
						if err != nil {
							return nil, err
						}
						refValue.Value = prefix + encodeJSONPointerSegment(schemaName)
						refValue.Tag = "!!str"
						key, err := graph.cloneYAMLScalarBounded(node.Content[index])
						if err != nil {
							return nil, err
						}
						clone.Content = append(clone.Content, key, refValue)
						continue
					}
					if !graph.referenceSiblingAllowed(context, node.Content[index].Value) {
						continue
					}
					childContext := graph.childContext(context, node.Content[index].Value)
					bundledSibling, err := graph.bundleNodeAtDepth(source, node.Content[index+1], resolving, aliases, childContext, depth+1)
					if err != nil {
						return nil, err
					}
					key, err := graph.cloneYAMLScalarBounded(node.Content[index])
					if err != nil {
						return nil, err
					}
					clone.Content = append(clone.Content, key, bundledSibling)
				}
				return &clone, nil
			}
			if context == openAPISchemaContext {
				return nil, fmt.Errorf("schema reference %q has no indexed internal target", value)
			}
			if resolving[identity] {
				return nil, fmt.Errorf("cyclic non-schema reference %s", identity)
			}
			resolving[identity] = true
			targetContext := graph.referenceTargetContext(context)
			bundled, err := graph.bundleNodeAtDepth(targetSource, target, resolving, aliases, targetContext, depth+1)
			delete(resolving, identity)
			if err != nil {
				return nil, err
			}
			if !graph.hasAllowedReferenceSiblings(node, context) {
				return bundled, nil
			}
			if bundled.Kind != yaml.MappingNode {
				return nil, fmt.Errorf("reference %q has sibling fields but target is not a mapping", value)
			}
			merged, err := graph.cloneYAMLNodeBounded(bundled, depth+1, true, make(map[*yaml.Node]bool))
			if err != nil {
				return nil, err
			}
			for index := 0; index+1 < len(node.Content); index += 2 {
				keyValue := node.Content[index].Value
				if keyValue == "$ref" || !graph.referenceSiblingAllowed(context, keyValue) {
					continue
				}
				key, err := graph.cloneYAMLScalarBounded(node.Content[index])
				if err != nil {
					return nil, err
				}
				childContext := graph.childContext(context, keyValue)
				value, err := graph.bundleNodeAtDepth(source, node.Content[index+1], resolving, aliases, childContext, depth+1)
				if err != nil {
					return nil, err
				}
				setYAMLMappingValue(merged, key, value)
			}
			return merged, nil
		}
	}

	clone := *node
	clone.Content = nil
	clone.Alias = nil
	if node.Kind == yaml.MappingNode {
		for index := 0; index+1 < len(node.Content); index += 2 {
			key, err := graph.cloneYAMLScalarBounded(node.Content[index])
			if err != nil {
				return nil, err
			}
			childContext := graph.childContext(context, node.Content[index].Value)
			value, err := graph.bundleNodeAtDepth(source, node.Content[index+1], resolving, aliases, childContext, depth+1)
			if err != nil {
				return nil, err
			}
			clone.Content = append(clone.Content, key, value)
		}
		return &clone, nil
	}
	childContext := graph.sequenceChildContext(context)
	if node.Kind == yaml.DocumentNode {
		childContext = context
	}
	for _, child := range node.Content {
		bundled, err := graph.bundleNodeAtDepth(source, child, resolving, aliases, childContext, depth+1)
		if err != nil {
			return nil, err
		}
		clone.Content = append(clone.Content, bundled)
	}
	return &clone, nil
}

func (graph *openAPIRefGraph) referenceSiblingAllowed(context openAPIBundleContext, key string) bool {
	if key == "$ref" {
		return false
	}
	if context == openAPIPathItemContext {
		return operationMethods[strings.ToLower(key)] || key == "summary" || key == "description" || key == "servers" || key == "parameters" || strings.HasPrefix(strings.ToLower(key), "x-")
	}
	if graph.kind == "swagger" || !graph.openAPI31 {
		return false
	}
	if context == openAPISchemaContext {
		return true
	}
	switch context {
	case openAPIExampleContext, openAPIPathItemContext:
		return key == "summary" || key == "description"
	case openAPIParameterContext, openAPIHeaderContext, openAPIRequestBodyContext, openAPIResponseContext, openAPISecuritySchemeContext, openAPILinkContext:
		return key == "description"
	default:
		return false
	}
}

func (graph *openAPIRefGraph) hasAllowedReferenceSiblings(node *yaml.Node, context openAPIBundleContext) bool {
	for index := 0; index+1 < len(node.Content); index += 2 {
		if graph.referenceSiblingAllowed(context, node.Content[index].Value) {
			return true
		}
	}
	return false
}

func yamlRefIndex(node *yaml.Node) int {
	if node == nil || node.Kind != yaml.MappingNode {
		return -1
	}
	for index := 0; index+1 < len(node.Content); index += 2 {
		if node.Content[index].Value == "$ref" {
			return index
		}
	}
	return -1
}

func (graph *openAPIRefGraph) consumeBundleNode() error {
	graph.bundleNodes++
	if graph.bundleNodes > maxOpenAPIBundleNodes {
		return fmt.Errorf("bundled OpenAPI reference exceeds %d nodes", maxOpenAPIBundleNodes)
	}
	return nil
}

func (graph *openAPIRefGraph) cloneYAMLScalarBounded(node *yaml.Node) (*yaml.Node, error) {
	if err := graph.consumeBundleNode(); err != nil {
		return nil, err
	}
	clone := *node
	clone.Content = nil
	clone.Alias = nil
	return &clone, nil
}

func (graph *openAPIRefGraph) cloneYAMLNodeBounded(node *yaml.Node, depth int, countRoot bool, aliases map[*yaml.Node]bool) (*yaml.Node, error) {
	if node == nil {
		return nil, errors.New("nil YAML node")
	}
	if depth > maxOpenAPIBundleDepth {
		return nil, fmt.Errorf("bundled OpenAPI reference exceeds depth %d", maxOpenAPIBundleDepth)
	}
	if countRoot {
		if err := graph.consumeBundleNode(); err != nil {
			return nil, err
		}
	}
	if node.Kind == yaml.AliasNode {
		if aliases[node] {
			return nil, errors.New("cyclic YAML alias in OpenAPI document")
		}
		aliases[node] = true
		clone, err := graph.cloneYAMLNodeBounded(node.Alias, depth+1, true, aliases)
		delete(aliases, node)
		return clone, err
	}
	clone := *node
	clone.Content = make([]*yaml.Node, len(node.Content))
	clone.Alias = nil
	for index, child := range node.Content {
		cloned, err := graph.cloneYAMLNodeBounded(child, depth+1, true, aliases)
		if err != nil {
			return nil, err
		}
		clone.Content[index] = cloned
	}
	return &clone, nil
}

func setYAMLMappingValue(mapping, key, value *yaml.Node) {
	for index := 0; index+1 < len(mapping.Content); index += 2 {
		if mapping.Content[index].Value == key.Value {
			mapping.Content[index], mapping.Content[index+1] = key, value
			return
		}
	}
	mapping.Content = append(mapping.Content, key, value)
}

func openAPIRefIdentity(source, pointer string) string {
	return source + "\x00" + pointer
}

func splitOpenAPIRefIdentity(identity string) (string, string) {
	separator := strings.IndexByte(identity, 0)
	return identity[:separator], identity[separator+1:]
}
