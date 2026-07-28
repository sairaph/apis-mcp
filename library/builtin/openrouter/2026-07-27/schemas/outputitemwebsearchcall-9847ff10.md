---
title: OutputItemWebSearchCall
page_id: schema-outputitemwebsearchcall-9847ff10
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemWebSearchCall

```yaml
{"example": {"action": {"query": "OpenAI API", "type": "search"}, "id": "search-abc123", "status": "completed", "type": "web_search_call"}, "properties": {"action": {"oneOf": [{"properties": {"queries": {"items": {"type": "string"}, "type": "array"}, "query": {"type": "string"}, "sources": {"items": {"$ref": "#/components/schemas/WebSearchSource"}, "type": "array"}, "type": {"enum": ["search"], "type": "string"}}, "required": ["type", "query"], "type": "object"}, {"properties": {"type": {"enum": ["open_page"], "type": "string"}, "url": {"type": ["string", "null"]}}, "required": ["type"], "type": "object"}, {"properties": {"pattern": {"type": "string"}, "type": {"enum": ["find_in_page"], "type": "string"}, "url": {"type": "string"}}, "required": ["type", "pattern", "url"], "type": "object"}]}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/WebSearchStatus"}, "type": {"enum": ["web_search_call"], "type": "string"}}, "required": ["type", "id", "status"], "type": "object"}
```
