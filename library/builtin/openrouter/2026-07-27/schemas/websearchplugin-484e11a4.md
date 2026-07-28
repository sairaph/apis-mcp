---
title: WebSearchPlugin
page_id: schema-websearchplugin-484e11a4
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WebSearchPlugin

```yaml
{"example": {"enabled": true, "id": "web", "max_results": 5}, "properties": {"enabled": {"description": "Set to false to disable the web-search plugin for this request. Defaults to true.", "type": "boolean"}, "engine": {"$ref": "#/components/schemas/WebSearchEngine"}, "exclude_domains": {"description": "A list of domains to exclude from web search results. Supports wildcards (e.g. \"*.substack.com\") and path filtering (e.g. \"openai.com/blog\").", "example": ["example.com", "*.substack.com", "openai.com/blog"], "items": {"type": "string"}, "type": "array"}, "id": {"enum": ["web"], "type": "string"}, "include_domains": {"description": "A list of domains to restrict web search results to. Supports wildcards (e.g. \"*.substack.com\") and path filtering (e.g. \"openai.com/blog\").", "example": ["example.com", "*.substack.com", "openai.com/blog"], "items": {"type": "string"}, "type": "array"}, "max_results": {"type": "integer"}, "max_uses": {"description": "Maximum number of times the model can invoke web search in a single turn. Passed through to native providers that support it (e.g. Anthropic).", "type": "integer"}, "search_prompt": {"type": "string"}, "user_location": {"allOf": [{"$ref": "#/components/schemas/WebSearchUserLocation"}, {"description": "Approximate user location for location-biased search results. Passed through to native providers that support it (e.g. Anthropic).", "example": {"city": "San Francisco", "country": "US", "region": "California", "timezone": "America/Los_Angeles", "type": "approximate"}, "required": ["type"]}]}}, "required": ["id"], "type": "object"}
```
