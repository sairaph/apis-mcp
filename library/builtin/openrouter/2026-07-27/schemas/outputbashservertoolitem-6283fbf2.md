---
title: OutputBashServerToolItem
page_id: schema-outputbashservertoolitem-6283fbf2
path: schemas
description: An openrouter:bash server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputBashServerToolItem

An openrouter:bash server tool output item

```yaml
{"description": "An openrouter:bash server tool output item", "example": {"command": "ls -la", "exitCode": 0, "id": "bash_tmp_abc123", "status": "completed", "stdout": "total 0\n", "type": "openrouter:bash"}, "properties": {"arguments": {"description": "The raw tool-call arguments string as emitted by the model.", "type": "string"}, "call_id": {"description": "The model-generated tool call id from the originating turn.", "type": "string"}, "command": {"type": "string"}, "exitCode": {"type": "integer"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "stderr": {"type": "string"}, "stdout": {"type": "string"}, "type": {"enum": ["openrouter:bash"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
