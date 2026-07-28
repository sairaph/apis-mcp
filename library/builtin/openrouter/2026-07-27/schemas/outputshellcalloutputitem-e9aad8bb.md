---
title: OutputShellCallOutputItem
page_id: schema-outputshellcalloutputitem-e9aad8bb
path: schemas
description: A native `shell_call_output` item matching OpenAI's Responses API shape. Carries per-command stdout, stderr, and the exit/timeout outcome.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputShellCallOutputItem

A native `shell_call_output` item matching OpenAI's Responses API shape. Carries per-command stdout, stderr, and the exit/timeout outcome.

```yaml
{"description": "A native `shell_call_output` item matching OpenAI's Responses API shape. Carries per-command stdout, stderr, and the exit/timeout outcome.", "example": {"call_id": "call_abc123", "id": "sho_abc123", "output": [{"outcome": {"exit_code": 0, "type": "exit"}, "stderr": "", "stdout": "hello\n"}], "status": "completed", "type": "shell_call_output"}, "properties": {"call_id": {"type": "string"}, "id": {"type": "string"}, "max_output_length": {"type": ["integer", "null"]}, "output": {"items": {"properties": {"outcome": {"oneOf": [{"properties": {"exit_code": {"type": "integer"}, "type": {"enum": ["exit"], "type": "string"}}, "required": ["type", "exit_code"], "type": "object"}, {"properties": {"type": {"enum": ["timeout"], "type": "string"}}, "required": ["type"], "type": "object"}]}, "stderr": {"type": "string"}, "stdout": {"type": "string"}}, "required": ["stdout", "stderr", "outcome"], "type": "object"}, "type": "array"}, "status": {"$ref": "#/components/schemas/ShellCallStatus"}, "type": {"enum": ["shell_call_output"], "type": "string"}}, "required": ["type", "id", "call_id", "status", "output"], "type": "object"}
```
