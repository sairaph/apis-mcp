package importer

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Detection identifies the ingestion engine and framework advertised by a URL.
type Detection struct {
	Engine    string `json:"engine"`
	Framework string `json:"framework,omitempty"`
	Format    string `json:"format"`
	Source    string `json:"source"`
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
	raw, provenance, err := newSourceReader(options).read(ctx, parsed.String(), nil)
	if err != nil {
		return Detection{}, err
	}
	if _, kind, parseErr := parseAPIDescription(raw); parseErr == nil {
		return Detection{Engine: "openapi", Framework: kind, Format: kind, Source: provenance}, nil
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
		return Detection{Engine: engine, Framework: framework, Format: "html", Source: provenance}, nil
	}
	base, err := url.Parse(provenance)
	if err == nil && (len(openAPISpecCandidates(document, base)) > 0 || len(openAPIConfigScripts(document, base)) > 0 || len(rapidocCatalogRoots(document, base)) > 0 || hasRapiDocComponent(document)) {
		return Detection{Engine: "openapi", Framework: detectOpenAPIHTMLFramework(document), Format: "html", Source: provenance}, nil
	}
	return Detection{Engine: "html", Framework: "unknown", Format: "html", Source: provenance}, nil
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
