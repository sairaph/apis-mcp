---
title: LocalShellCallItem
page_id: schema-localshellcallitem-7e256d1b
path: schemas
description: A local shell command execution call
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# LocalShellCallItem

A local shell command execution call

```yaml
{"description": "A local shell command execution call", "example": {"action": {"command": ["ls", "-la"], "env": {"PATH": "/usr/bin"}, "timeout_ms": 5000, "type": "exec"}, "call_id": "call-abc123", "id": "shell-abc123", "status": "completed", "type": "local_shell_call"}, "properties": {"action": {"properties": {"command": {"items": {"type": "string"}, "type": "array"}, "env": {"additionalProperties": {"type": "string"}, "type": "object"}, "timeout_ms": {"type": ["integer", "null"]}, "type": {"enum": ["exec"], "type": "string"}, "user": {"type": ["string", "null"]}, "working_directory": {"type": ["string", "null"]}}, "required": ["type", "command", "env"], "type": "object"}, "call_id": {"type": "string"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["local_shell_call"], "type": "string"}}, "required": ["type", "id", "call_id", "action", "status"], "type": "object"}
```
