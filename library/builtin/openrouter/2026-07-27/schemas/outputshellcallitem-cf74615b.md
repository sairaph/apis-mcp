---
title: OutputShellCallItem
page_id: schema-outputshellcallitem-cf74615b
path: schemas
description: A native `shell_call` output item matching OpenAI's Responses API shape. Emitted for the sandbox-backed `shell` tool.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputShellCallItem

A native `shell_call` output item matching OpenAI's Responses API shape. Emitted for the sandbox-backed `shell` tool.

```yaml
{"description": "A native `shell_call` output item matching OpenAI's Responses API shape. Emitted for the sandbox-backed `shell` tool.", "example": {"action": {"commands": ["echo hello"], "max_output_length": null, "timeout_ms": null}, "call_id": "call_abc123", "id": "shc_abc123", "status": "completed", "type": "shell_call"}, "properties": {"action": {"properties": {"commands": {"items": {"type": "string"}, "type": "array"}, "max_output_length": {"type": ["integer", "null"]}, "timeout_ms": {"type": ["integer", "null"]}}, "required": ["commands", "max_output_length", "timeout_ms"], "type": "object"}, "call_id": {"type": "string"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ShellCallStatus"}, "type": {"enum": ["shell_call"], "type": "string"}}, "required": ["type", "id", "call_id", "status"], "type": "object"}
```
