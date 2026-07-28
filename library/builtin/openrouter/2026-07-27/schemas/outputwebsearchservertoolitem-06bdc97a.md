---
title: OutputWebSearchServerToolItem
page_id: schema-outputwebsearchservertoolitem-06bdc97a
path: schemas
description: An openrouter:web_search server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputWebSearchServerToolItem

An openrouter:web_search server tool output item

```yaml
{"description": "An openrouter:web_search server tool output item", "example": {"action": {"query": "latest AI news", "type": "search"}, "id": "ws_tmp_abc123", "status": "completed", "type": "openrouter:web_search"}, "properties": {"action": {"description": "The search action performed, matching OpenAI web_search_call.action shape. Includes the query the model issued and optional source URLs returned by the search provider.", "properties": {"query": {"type": "string"}, "sources": {"items": {"properties": {"type": {"enum": ["url"], "type": "string"}, "url": {"type": "string"}}, "required": ["type", "url"], "type": "object"}, "type": "array"}, "type": {"enum": ["search"], "type": "string"}}, "required": ["type", "query"], "type": "object"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:web_search"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
