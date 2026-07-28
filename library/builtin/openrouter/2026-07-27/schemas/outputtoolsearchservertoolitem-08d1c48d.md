---
title: OutputToolSearchServerToolItem
page_id: schema-outputtoolsearchservertoolitem-08d1c48d
path: schemas
description: An openrouter:tool_search server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputToolSearchServerToolItem

An openrouter:tool_search server tool output item

```yaml
{"description": "An openrouter:tool_search server tool output item", "example": {"id": "ts_tmp_abc123", "query": "weather tools", "status": "completed", "type": "openrouter:tool_search"}, "properties": {"id": {"type": "string"}, "query": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:tool_search"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
