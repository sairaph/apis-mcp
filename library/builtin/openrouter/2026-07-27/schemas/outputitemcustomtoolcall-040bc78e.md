---
title: OutputItemCustomToolCall
page_id: schema-outputitemcustomtoolcall-040bc78e
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemCustomToolCall

```yaml
{"example": {"call_id": "call-abc123", "id": "ctc-abc123", "input": "*** Begin Patch\n*** End Patch", "name": "apply_patch", "type": "custom_tool_call"}, "properties": {"call_id": {"type": "string"}, "id": {"type": "string"}, "input": {"type": "string"}, "name": {"type": "string"}, "namespace": {"description": "Namespace qualifier for tools registered as part of a namespace tool group (e.g. an MCP server)", "type": "string"}, "type": {"enum": ["custom_tool_call"], "type": "string"}}, "required": ["type", "name", "input", "call_id"], "type": "object"}
```
