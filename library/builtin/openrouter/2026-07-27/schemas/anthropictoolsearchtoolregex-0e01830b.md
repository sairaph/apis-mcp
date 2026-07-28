---
title: AnthropicToolSearchToolRegex
page_id: schema-anthropictoolsearchtoolregex-0e01830b
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicToolSearchToolRegex

```yaml
{"properties": {"allowed_callers": {"$ref": "#/components/schemas/AnthropicAllowedCallers"}, "cache_control": {"$ref": "#/components/schemas/AnthropicCacheControlDirective"}, "defer_loading": {"type": "boolean"}, "name": {"enum": ["tool_search_tool_regex"], "type": "string"}, "strict": {"type": "boolean"}, "type": {"enum": ["tool_search_tool_regex_20251119", "tool_search_tool_regex"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["type", "name"], "type": "object"}
```
