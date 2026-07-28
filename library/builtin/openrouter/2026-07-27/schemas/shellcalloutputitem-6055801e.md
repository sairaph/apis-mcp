---
title: ShellCallOutputItem
page_id: schema-shellcalloutputitem-6055801e
path: schemas
description: Output from a shell command execution (newer variant)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ShellCallOutputItem

Output from a shell command execution (newer variant)

```yaml
{"description": "Output from a shell command execution (newer variant)", "example": {"call_id": "call-abc123", "output": [{"content": "total 0\n", "type": "stdout"}], "status": "completed", "type": "shell_call_output"}, "properties": {"call_id": {"type": "string"}, "id": {"type": ["string", "null"]}, "max_output_length": {"type": ["integer", "null"]}, "output": {"items": {"additionalProperties": {}, "properties": {"content": {"type": ["string", "null"]}, "exit_code": {"type": ["integer", "null"]}, "type": {"type": "string"}}, "required": ["type"], "type": "object"}, "type": "array"}, "status": {"anyOf": [{"$ref": "#/components/schemas/ToolCallStatus"}, {"type": "null"}]}, "type": {"enum": ["shell_call_output"], "type": "string"}}, "required": ["type", "call_id", "output"], "type": "object"}
```
