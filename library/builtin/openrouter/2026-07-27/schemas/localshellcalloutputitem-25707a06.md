---
title: LocalShellCallOutputItem
page_id: schema-localshellcalloutputitem-25707a06
path: schemas
description: Output from a local shell command execution
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# LocalShellCallOutputItem

Output from a local shell command execution

```yaml
{"description": "Output from a local shell command execution", "example": {"id": "output-abc123", "output": "total 24\ndrwxr-xr-x  5 user  staff  160 Jan  1 12:00 .", "status": "completed", "type": "local_shell_call_output"}, "properties": {"id": {"type": "string"}, "output": {"type": "string"}, "status": {"anyOf": [{"$ref": "#/components/schemas/ToolCallStatus"}, {"type": "null"}]}, "type": {"enum": ["local_shell_call_output"], "type": "string"}}, "required": ["type", "id", "output"], "type": "object"}
```
