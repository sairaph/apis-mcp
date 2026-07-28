---
title: ShellCallItem
page_id: schema-shellcallitem-b46a71ce
path: schemas
description: A shell command execution call (newer variant)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ShellCallItem

A shell command execution call (newer variant)

```yaml
{"description": "A shell command execution call (newer variant)", "example": {"action": {"commands": ["ls", "-la"], "max_output_length": 10000}, "call_id": "call-abc123", "status": "completed", "type": "shell_call"}, "properties": {"action": {"properties": {"commands": {"items": {"type": "string"}, "type": "array"}, "max_output_length": {"type": ["integer", "null"]}, "timeout_ms": {"type": ["integer", "null"]}}, "required": ["commands"], "type": "object"}, "call_id": {"type": "string"}, "environment": {}, "id": {"type": ["string", "null"]}, "status": {"anyOf": [{"$ref": "#/components/schemas/ToolCallStatus"}, {"type": "null"}]}, "type": {"enum": ["shell_call"], "type": "string"}}, "required": ["type", "call_id", "action"], "type": "object"}
```
