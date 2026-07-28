---
title: OutputCodeInterpreterServerToolItem
page_id: schema-outputcodeinterpreterservertoolitem-14b38ffe
path: schemas
description: An openrouter:code_interpreter server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputCodeInterpreterServerToolItem

An openrouter:code_interpreter server tool output item

```yaml
{"description": "An openrouter:code_interpreter server tool output item", "example": {"code": "print(\"hello\")", "id": "ci_tmp_abc123", "language": "python", "status": "completed", "stdout": "hello\n", "type": "openrouter:code_interpreter"}, "properties": {"code": {"type": "string"}, "exitCode": {"type": "integer"}, "id": {"type": "string"}, "language": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "stderr": {"type": "string"}, "stdout": {"type": "string"}, "type": {"enum": ["openrouter:code_interpreter"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
