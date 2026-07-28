---
title: OutputFileSearchServerToolItem
page_id: schema-outputfilesearchservertoolitem-bc64707b
path: schemas
description: An openrouter:file_search server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputFileSearchServerToolItem

An openrouter:file_search server tool output item

```yaml
{"description": "An openrouter:file_search server tool output item", "example": {"id": "fs_tmp_abc123", "queries": ["search term"], "status": "completed", "type": "openrouter:file_search"}, "properties": {"id": {"type": "string"}, "queries": {"items": {"type": "string"}, "type": "array"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:file_search"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
