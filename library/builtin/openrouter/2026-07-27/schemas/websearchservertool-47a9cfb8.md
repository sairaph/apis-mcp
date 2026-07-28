---
title: WebSearchServerTool
page_id: schema-websearchservertool-47a9cfb8
path: schemas
description: Web search tool configuration (2025-08-26 version)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebSearchServerTool

Web search tool configuration (2025-08-26 version)

```yaml
{"description": "Web search tool configuration (2025-08-26 version)", "example": {"engine": "auto", "filters": {"allowed_domains": ["example.com"]}, "type": "web_search_2025_08_26"}, "properties": {"engine": {"$ref": "#/components/schemas/WebSearchEngineEnum"}, "filters": {"$ref": "#/components/schemas/WebSearchDomainFilter"}, "max_results": {"description": "Maximum number of search results to return per search call. Defaults to 5. Applies to Exa, Firecrawl, Parallel, and Perplexity engines; ignored with native provider search. Perplexity supports a maximum of 20; values above 20 are clamped.", "example": 5, "type": "integer"}, "max_uses": {"description": "Maximum number of web searches the model may perform in a single request. Once reached, further search calls return an error result instead of executing. Applies to the Exa, Firecrawl, Parallel, and Perplexity engines. With native provider search, forwarded only to Anthropic (as `max_uses`); other native search providers have no equivalent parameter and ignore it.", "example": 3, "type": "integer"}, "search_context_size": {"$ref": "#/components/schemas/SearchContextSizeEnum"}, "type": {"enum": ["web_search_2025_08_26"], "type": "string"}, "user_location": {"$ref": "#/components/schemas/WebSearchUserLocation"}}, "required": ["type"], "type": "object"}
```
