package importer

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

const maxLLMSLinks = 256

var markdownLink = regexp.MustCompile(`(?m)\[([^\]]+)\]\(([^\s)]+)(?:\s+"[^"]*")?\)(?:[ \t]*(?::|-)[ \t]*(.+))?[ \t]*$`)

// ImportLLMSTxt imports Markdown documents linked by a local or HTTP(S)
// llms.txt index. It intentionally does not render JavaScript or crawl HTML.
func ImportLLMSTxt(ctx context.Context, name, version, source string, options Options) (Result, error) {
	options, err := normalizeOptions(options)
	if err != nil {
		return Result{}, err
	}
	if strings.TrimSpace(name) == "" || strings.TrimSpace(version) == "" {
		return Result{}, errors.New("llms.txt import requires API name and version")
	}
	reader := newSourceReader(options)
	indexRaw, provenance, err := reader.read(ctx, source, nil)
	if err != nil {
		return Result{}, err
	}
	links := markdownLink.FindAllStringSubmatch(string(indexRaw), -1)
	if len(links) == 0 {
		return Result{}, errors.New("llms.txt contains no Markdown links")
	}
	if len(links) > maxLLMSLinks {
		return Result{}, fmt.Errorf("llms.txt contains more than %d links", maxLLMSLinks)
	}
	base, err := llmsBase(provenance)
	if err != nil {
		return Result{}, err
	}
	type linkedDocument struct {
		title, description, source, body string
		front                            map[string]any
	}
	seen := make(map[string]bool)
	documents := make([]linkedDocument, 0, len(links))
	for _, match := range links {
		resolved, _, err := resolveSource(match[2], base)
		if err != nil {
			return Result{}, err
		}
		if seen[resolved] {
			continue
		}
		raw, resolved, err := reader.read(ctx, match[2], base)
		if err != nil {
			return Result{}, err
		}
		seen[resolved] = true
		front := make(map[string]any)
		body := string(raw)
		if strings.HasPrefix(strings.TrimSpace(body), "---") {
			parsedBody, parseErr := splitFrontmatter(raw, &front)
			if parseErr != nil {
				return Result{}, fmt.Errorf("parse frontmatter from %s: %w", resolved, parseErr)
			}
			body = parsedBody
		}
		title := strings.TrimSpace(match[1])
		if heading := firstHeading(body); heading != "" {
			title = heading
		}
		if existing, ok := front["title"].(string); ok && strings.TrimSpace(existing) != "" {
			title = strings.TrimSpace(existing)
		}
		if title == "" {
			title = "Documentation"
		}
		description := ""
		if len(match) > 3 {
			description = strings.TrimSpace(match[3])
		}
		if existing, ok := front["description"].(string); ok && strings.TrimSpace(existing) != "" {
			description = strings.TrimSpace(existing)
		}
		documents = append(documents, linkedDocument{
			title: title, description: description, source: resolved, body: body, front: front,
		})
	}
	if len(documents) == 0 {
		return Result{}, errors.New("llms.txt contains no unique documents")
	}

	result, err := publish(ctx, options, name, version, func(stage string) error {
		metadata := manifest{
			Name: name, Version: version, Collections: options.Collections,
			SourceRoot: provenance, SourceType: "llms.txt", ImportedFrom: provenance,
		}
		if err := writeCanonicalFile(stage, "_index.md", metadata, "This document set was imported from llms.txt."); err != nil {
			return err
		}
		for _, document := range documents {
			front := document.front
			if upstream, ok := front["source"]; ok && fmt.Sprint(upstream) != document.source {
				front["upstream_source"] = upstream
			}
			front["title"] = document.title
			if document.description != "" {
				front["description"] = document.description
			} else {
				delete(front, "description")
			}
			front["page_id"] = stableID("page-", document.source)
			front["path"] = "documentation"
			front["source"] = document.source
			front["source_type"] = "llms.txt"
			front["imported_from"] = provenance
			filename := stableID("", document.source) + ".md"
			if err := writeCanonicalFile(stage, filepath.Join("documentation", filename), front, document.body); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return Result{}, err
	}
	result.Kind, result.Source, result.Pages = "llms.txt", provenance, len(documents)
	return result, nil
}

func llmsBase(source string) (*url.URL, error) {
	parsed, err := url.Parse(source)
	if err != nil {
		return nil, err
	}
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	if parsed.Scheme == "http" || parsed.Scheme == "https" {
		return parsed, nil
	}
	directory := filepath.Dir(source)
	return &url.URL{Scheme: "file", Path: filepath.ToSlash(directory) + "/"}, nil
}

func firstHeading(markdown string) string {
	for _, line := range strings.Split(markdown, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
		}
	}
	return ""
}

func firstParagraph(markdown string) string {
	var lines []string
	for _, line := range strings.Split(markdown, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if len(lines) > 0 {
				break
			}
			continue
		}
		if strings.HasPrefix(trimmed, "#") || markdownLink.MatchString(trimmed) {
			continue
		}
		lines = append(lines, trimmed)
	}
	return strings.Join(lines, " ")
}
