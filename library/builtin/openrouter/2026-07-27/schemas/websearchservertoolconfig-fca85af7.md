---
title: WebSearchServerToolConfig
page_id: schema-websearchservertoolconfig-fca85af7
path: schemas
description: Configuration for the openrouter:web_search server tool
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebSearchServerToolConfig

Configuration for the openrouter:web_search server tool

```yaml
{"description": "Configuration for the openrouter:web_search server tool", "example": {"max_results": 5, "search_context_size": "medium"}, "properties": {"allowed_domains": {"description": "Limit search results to these domains. Supported by Exa, Firecrawl, Parallel, Perplexity, and most native providers (Anthropic, OpenAI, xAI). Cannot be used with excluded_domains.", "items": {"type": "string"}, "type": "array"}, "engine": {"$ref": "#/components/schemas/WebSearchEngineEnum"}, "excluded_domains": {"description": "Exclude search results from these domains. Supported by Exa, Firecrawl, Parallel, Perplexity, Anthropic, and xAI. Not supported with OpenAI (silently ignored). Cannot be used with allowed_domains.", "items": {"type": "string"}, "type": "array"}, "max_characters": {"description": "Exact maximum number of characters of content per search result. Applies to the Exa, Parallel, and Perplexity engines; ignored with native provider search and Firecrawl. For Exa, caps highlight content per result. For Parallel, caps excerpt content per result (default 1,500 when omitted). For Perplexity, maps to the native `max_tokens_per_page` parameter (converted from characters to tokens) and trims the response to the exact character cap. When both `max_characters` and `search_context_size` are set, `max_characters` takes precedence. When omitted, falls back to `search_context_size` mapping (Exa) or engine defaults (Parallel, Perplexity).", "example": 2000, "type": "integer"}, "max_results": {"description": "Maximum number of search results to return per search call. Defaults to 5. Applies to Exa, Firecrawl, Parallel, and Perplexity engines; ignored with native provider search. Perplexity supports a maximum of 20; values above 20 are clamped.", "example": 5, "type": "integer"}, "max_total_results": {"description": "Maximum total number of search results across all search calls in a single request. Once this limit is reached, the tool will stop returning new results. Useful for controlling cost and context size in agentic loops. Defaults to 50 when not specified.", "example": 50, "type": "integer"}, "max_uses": {"description": "Maximum number of web searches the model may perform in a single request. Once reached, further search calls return an error result instead of executing. Applies to the Exa, Firecrawl, Parallel, and Perplexity engines. With native provider search, forwarded only to Anthropic (as `max_uses`); other native search providers have no equivalent parameter and ignore it.", "example": 3, "type": "integer"}, "search_context_size": {"$ref": "#/components/schemas/SearchQualityLevel"}, "user_location": {"$ref": "#/components/schemas/WebSearchUserLocationServerTool"}}, "type": "object"}
```
