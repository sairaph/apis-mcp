package importer

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var astroGenerator = regexp.MustCompile(`(?i)^astro v[0-9]+\.[0-9]+\.[0-9]+(?:-[0-9a-z.-]+)?(?:\+[0-9a-z.-]+)?$`)

// Detection identifies the ingestion engine and framework advertised by a URL.
type Detection struct {
	Engine          string `json:"engine"`
	Framework       string `json:"framework,omitempty"`
	Format          string `json:"format"`
	Source          string `json:"source"`
	DownloadedBytes int64  `json:"downloaded_bytes,omitempty"`
}

// DetectURL inspects one HTTP(S) source without publishing documentation.
func DetectURL(ctx context.Context, source string, options Options) (Detection, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return Detection{}, err
	}
	parsed, err := url.Parse(strings.TrimSpace(source))
	if err != nil || parsed.Host == "" || parsed.User != nil || parsed.Scheme != "http" && parsed.Scheme != "https" {
		return Detection{}, errors.New("detection source must be an HTTP(S) URL")
	}
	reader := newSourceReader(options)
	detected := func(result Detection) Detection {
		result.DownloadedBytes = reader.used
		return result
	}
	raw, provenance, err := reader.read(ctx, parsed.String(), nil)
	if err != nil {
		return Detection{}, err
	}
	if _, kind, parseErr := parseAPIDescription(raw); parseErr == nil {
		return detected(Detection{Engine: "openapi", Framework: kind, Format: kind, Source: provenance}), nil
	}
	document, err := parseHTML(raw)
	if err != nil {
		return Detection{}, fmt.Errorf("parse detection source: %w", err)
	}
	if !isHTMLDocument(document) {
		return Detection{}, errors.New("source is neither OpenAPI nor HTML documentation")
	}
	if framework := detectHTMLFramework(document); framework != "" {
		engine := "html"
		if framework == "docsify" {
			engine = "docsify"
		}
		return detected(Detection{Engine: engine, Framework: framework, Format: "html", Source: provenance}), nil
	}
	if framework := detectProductHTMLFramework(document); framework == "mintlify" || framework == "stripe-docs" {
		return detected(Detection{Engine: "html", Framework: framework, Format: "html", Source: provenance}), nil
	}
	base, err := url.Parse(provenance)
	if err == nil {
		openAPI := len(openAPISpecCandidates(document, base)) > 0 || len(openAPIConfigScripts(document, base)) > 0 || len(rapidocCatalogRoots(document, base)) > 0 || hasRapiDocComponent(document)
		framework := detectOpenAPIHTMLFramework(document)
		if openAPI && framework != "openapi-ui" {
			return detected(Detection{Engine: "openapi", Framework: framework, Format: "html", Source: provenance}), nil
		}
		scalar, scalarErr := resolveScalarSchema(ctx, document, base, reader)
		if scalarErr != nil {
			return Detection{}, scalarErr
		}
		if scalar.Proven {
			if scalar.URL != "" {
				return detected(Detection{Engine: "openapi", Framework: "scalar", Format: "html", Source: scalar.URL}), nil
			}
			return detected(Detection{Engine: "html", Framework: "scalar", Format: "html", Source: provenance}), nil
		}
		if framework := detectProductHTMLFramework(document); framework != "" {
			return detected(Detection{Engine: "html", Framework: framework, Format: "html", Source: provenance}), nil
		}
		if openAPI {
			return detected(Detection{Engine: "openapi", Framework: framework, Format: "html", Source: provenance}), nil
		}
	} else if framework := detectProductHTMLFramework(document); framework != "" {
		return detected(Detection{Engine: "html", Framework: framework, Format: "html", Source: provenance}), nil
	}
	if framework := detectRuntimeHTMLFramework(document); framework != "" {
		return detected(Detection{Engine: "html", Framework: framework, Format: "html", Source: provenance}), nil
	}
	return detected(Detection{Engine: "html", Framework: "unknown", Format: "html", Source: provenance}), nil
}

// Product shells are checked before generic OpenAPI discovery because their
// embedded configuration does not constitute a complete static spec inventory.
func detectProductHTMLFramework(root *htmlNode) string {
	switch {
	case looksLikeStripeDocs(root):
		return "stripe-docs"
	case looksLikeMintlify(root):
		return "mintlify"
	case looksLikeScalar(root):
		return "scalar"
	default:
		return ""
	}
}

// Runtime frameworks remain below OpenAPI UI discovery so a Swagger, Redoc,
// or RapiDoc shell is not classified by the framework used to render it.
func detectRuntimeHTMLFramework(root *htmlNode) string {
	if hasGenerator(root, astroGenerator.MatchString) {
		return "astro"
	}
	if looksLikeSvelteKit(root) {
		return "sveltekit"
	}
	return ""
}

func hasGenerator(root *htmlNode, match func(string) bool) bool {
	found := false
	walkHTML(root, func(node *htmlNode) {
		if !found && node.tag == "meta" && strings.EqualFold(strings.TrimSpace(node.attrs["name"]), "generator") {
			found = match(strings.TrimSpace(node.attrs["content"]))
		}
	})
	return found
}

func looksLikeMintlify(root *htmlNode) bool {
	if hasGenerator(root, func(value string) bool { return strings.EqualFold(value, "Mintlify") }) {
		return true
	}
	asset, cdnAsset, config, nextRuntime := false, false, false, false
	walkHTML(root, func(node *htmlNode) {
		if node.tag == "script" || node.tag == "link" {
			value := strings.ToLower(node.attrs["src"] + " " + node.attrs["href"])
			asset = asset || containsAny(value, "/_mintlify/", "/mintlify-assets/", "@mintlify/")
			cdnAsset = cdnAsset || containsAny(value, "mintcdn.com/", "mintlify.s3.")
			nextRuntime = nextRuntime || strings.Contains(value, "/_next/static/")
		}
		if node.tag == "script" {
			asset = asset || strings.HasPrefix(strings.ToLower(node.attrs["id"]), "_mintlify-")
			text := strings.ToLower(htmlNodeText(node))
			config = config || containsAny(text, "__mintlify", "mintlifyconfig", "mintlify_config", "window.mintlify")
		}
	})
	return asset || nextRuntime && (config || cdnAsset)
}

func looksLikeScalar(root *htmlNode) bool {
	found := false
	walkHTML(root, func(node *htmlNode) {
		if found {
			return
		}
		if node.tag == "scalar-api-reference" {
			found = true
			return
		}
		if node.tag == "script" {
			found = scalarScriptEvidence(node, htmlNodeText(node))
		}
	})
	return found
}

func looksLikeSvelteKit(root *htmlNode) bool {
	runtime, immutable, preload, hydration := false, false, false, false
	walkHTML(root, func(node *htmlNode) {
		for name := range node.attrs {
			preload = preload || strings.HasPrefix(name, "data-sveltekit-")
			hydration = hydration || name == "data-svelte-h"
		}
		if node.tag != "script" && node.tag != "link" {
			return
		}
		value := strings.ToLower(node.attrs["src"] + " " + node.attrs["href"])
		if node.tag == "script" {
			value += " " + strings.ToLower(htmlNodeText(node))
		}
		runtime = runtime || strings.Contains(value, "__sveltekit")
		immutable = immutable || strings.Contains(value, "/_app/immutable/")
	})
	markers := 0
	for _, present := range []bool{runtime, immutable, preload, hydration} {
		if present {
			markers++
		}
	}
	return markers >= 2 && immutable
}

func looksLikeStripeDocs(root *htmlNode) bool {
	stripe, apiReference, sail := false, false, false
	walkHTML(root, func(node *htmlNode) {
		identity := strings.ToLower(node.attrs["id"] + " " + node.attrs["class"] + " " + node.attrs["data-testid"] + " " + node.attrs["data-page-type"] + " " + node.attrs["data-js-controller"])
		asset := ""
		if node.tag == "script" || node.tag == "link" {
			asset = strings.ToLower(node.attrs["src"] + " " + node.attrs["href"])
		}
		text := ""
		if node.tag == "script" {
			text = strings.ToLower(htmlNodeText(node))
		}
		combined := identity + " " + asset + " " + text
		stripe = stripe || containsAny(asset, "stripecdn.com/docs-", "/docs-statics-srv/")
		apiReference = apiReference || containsAny(combined, "api-reference", "api_reference", "apireference")
		sail = sail || containsAny(combined, "sail-", "/sail/", "/assets/sail.", "sail_app", "sailapp")
	})
	return stripe && apiReference && sail
}

func containsAny(value string, markers ...string) bool {
	for _, marker := range markers {
		if strings.Contains(value, marker) {
			return true
		}
	}
	return false
}

func detectOpenAPIHTMLFramework(root *htmlNode) string {
	framework := "openapi-ui"
	walkHTML(root, func(node *htmlNode) {
		identity := strings.ToLower(node.tag + " " + node.attrs["id"] + " " + node.attrs["class"] + " " + node.attrs["src"])
		switch {
		case node.tag == "redoc" || strings.Contains(identity, "redoc"):
			framework = "redoc"
		case node.tag == "rapi-doc" || node.tag == "rapi-doc-mini" || strings.Contains(identity, "rapidoc"):
			framework = "rapidoc"
		case strings.Contains(identity, "swagger") || node.tag == "script" && strings.Contains(strings.ToLower(htmlNodeText(node)), "swaggerui"):
			framework = "swagger-ui"
		}
	})
	return framework
}
