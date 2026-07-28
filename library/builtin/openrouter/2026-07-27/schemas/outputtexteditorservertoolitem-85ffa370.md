---
title: OutputTextEditorServerToolItem
page_id: schema-outputtexteditorservertoolitem-85ffa370
path: schemas
description: An openrouter:text_editor server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputTextEditorServerToolItem

An openrouter:text_editor server tool output item

```yaml
{"description": "An openrouter:text_editor server tool output item", "example": {"command": "view", "filePath": "/src/main.ts", "id": "te_tmp_abc123", "status": "completed", "type": "openrouter:text_editor"}, "properties": {"command": {"enum": ["view", "create", "str_replace", "insert"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "filePath": {"type": "string"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:text_editor"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
