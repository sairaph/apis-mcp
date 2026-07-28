---
title: Preview_WebSearchServerTool
page_id: schema-preview-websearchservertool-44b729c9
path: schemas
description: Web search preview tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Preview_WebSearchServerTool

Web search preview tool configuration

```yaml
{"description": "Web search preview tool configuration", "example": {"type": "web_search_preview"}, "properties": {"engine": {"$ref": "#/components/schemas/WebSearchEngineEnum"}, "filters": {"$ref": "#/components/schemas/WebSearchDomainFilter"}, "max_results": {"description": "Maximum number of search results to return per search call. Defaults to 5. Applies to Exa, Firecrawl, Parallel, and Perplexity engines; ignored with native provider search. Perplexity supports a maximum of 20; values above 20 are clamped.", "example": 5, "type": "integer"}, "max_uses": {"description": "Maximum number of web searches the model may perform in a single request. Once reached, further search calls return an error result instead of executing. Applies to the Exa, Firecrawl, Parallel, and Perplexity engines. With native provider search, forwarded only to Anthropic (as `max_uses`); other native search providers have no equivalent parameter and ignore it.", "example": 3, "type": "integer"}, "search_context_size": {"$ref": "#/components/schemas/SearchContextSizeEnum"}, "type": {"enum": ["web_search_preview"], "type": "string"}, "user_location": {"$ref": "#/components/schemas/Preview_WebSearchUserLocation"}}, "required": ["type"], "type": "object"}
```
