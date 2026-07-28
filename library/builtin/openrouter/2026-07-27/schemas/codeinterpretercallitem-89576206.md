---
title: CodeInterpreterCallItem
page_id: schema-codeinterpretercallitem-89576206
path: schemas
description: A code interpreter execution call with outputs
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CodeInterpreterCallItem

A code interpreter execution call with outputs

```yaml
{"description": "A code interpreter execution call with outputs", "example": {"code": "print(\"Hello, World!\")", "container_id": "container-xyz789", "id": "code-abc123", "outputs": [{"logs": "Hello, World!", "type": "logs"}], "status": "completed", "type": "code_interpreter_call"}, "properties": {"code": {"type": ["string", "null"]}, "container_id": {"type": "string"}, "id": {"type": "string"}, "outputs": {"items": {"anyOf": [{"properties": {"type": {"enum": ["image"], "type": "string"}, "url": {"type": "string"}}, "required": ["type", "url"], "type": "object"}, {"properties": {"logs": {"type": "string"}, "type": {"enum": ["logs"], "type": "string"}}, "required": ["type", "logs"], "type": "object"}]}, "type": ["array", "null"]}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["code_interpreter_call"], "type": "string"}}, "required": ["type", "id", "code", "outputs", "status", "container_id"], "type": "object"}
```
