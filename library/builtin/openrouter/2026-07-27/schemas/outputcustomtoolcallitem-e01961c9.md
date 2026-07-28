---
title: OutputCustomToolCallItem
page_id: schema-outputcustomtoolcallitem-e01961c9
path: schemas
description: A call to a custom (freeform-grammar) tool created by the model — distinct from `function_call`. Used for tools like Codex CLI's `apply_patch` whose payload is opaque text rather than JSON arguments.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputCustomToolCallItem

A call to a custom (freeform-grammar) tool created by the model — distinct from `function_call`. Used for tools like Codex CLI's `apply_patch` whose payload is opaque text rather than JSON arguments.

```yaml
{"description": "A call to a custom (freeform-grammar) tool created by the model — distinct from `function_call`. Used for tools like Codex CLI's `apply_patch` whose payload is opaque text rather than JSON arguments.", "example": {"call_id": "call-abc123", "id": "ctc-abc123", "input": "*** Begin Patch\n*** End Patch", "name": "apply_patch", "type": "custom_tool_call"}, "properties": {"call_id": {"type": "string"}, "id": {"type": "string"}, "input": {"type": "string"}, "name": {"type": "string"}, "namespace": {"description": "Namespace qualifier for tools registered as part of a namespace tool group (e.g. an MCP server)", "type": "string"}, "type": {"enum": ["custom_tool_call"], "type": "string"}}, "required": ["type", "name", "input", "call_id"], "type": "object"}
```
