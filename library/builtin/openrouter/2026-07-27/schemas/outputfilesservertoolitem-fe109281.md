---
title: OutputFilesServerToolItem
page_id: schema-outputfilesservertoolitem-fe109281
path: schemas
description: An openrouter:files server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputFilesServerToolItem

An openrouter:files server tool output item

```yaml
{"description": "An openrouter:files server tool output item", "example": {"filename": "notes.txt", "id": "fl_tmp_abc123", "operation": "read", "result": "{\"id\":\"file_abc\",\"filename\":\"notes.txt\",\"content\":\"hello\"}", "status": "completed", "type": "openrouter:files"}, "properties": {"error": {"description": "Error message when the file operation failed.", "type": "string"}, "file_id": {"description": "The target file id supplied in the tool-call arguments.", "type": "string"}, "filename": {"description": "The target filename supplied in the tool-call arguments.", "type": "string"}, "id": {"type": "string"}, "operation": {"description": "The file operation performed (list, read, write, or edit).", "type": "string"}, "result": {"description": "JSON-serialized result of the file operation.", "type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:files"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
