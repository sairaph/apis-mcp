---
title: OutputSearchModelsServerToolItem
page_id: schema-outputsearchmodelsservertoolitem-b2b52e79
path: schemas
description: An openrouter:experimental__search_models server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputSearchModelsServerToolItem

An openrouter:experimental__search_models server tool output item

```yaml
{"description": "An openrouter:experimental__search_models server tool output item", "example": {"arguments": "{\"query\":\"Claude Opus\"}", "id": "sm_tmp_abc123", "query": "Claude Opus", "status": "completed", "type": "openrouter:experimental__search_models"}, "properties": {"arguments": {"description": "The JSON arguments submitted to the search tool (e.g. {\"query\":\"Claude\"})", "type": "string"}, "id": {"type": "string"}, "query": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:experimental__search_models"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
