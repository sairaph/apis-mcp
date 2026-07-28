---
title: OutputApplyPatchCallItem
page_id: schema-outputapplypatchcallitem-dc9160a9
path: schemas
description: A native `apply_patch_call` output item matching OpenAI's Responses API shape. Emitted when the client requested the `apply_patch` shorthand.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputApplyPatchCallItem

A native `apply_patch_call` output item matching OpenAI's Responses API shape. Emitted when the client requested the `apply_patch` shorthand.

```yaml
{"description": "A native `apply_patch_call` output item matching OpenAI's Responses API shape. Emitted when the client requested the `apply_patch` shorthand.", "example": {"call_id": "call_abc123", "id": "apc_abc123", "operation": {"diff": "@@ function main() {\n+  console.log(\"hi\");\n }", "path": "/src/main.ts", "type": "update_file"}, "status": "completed", "type": "apply_patch_call"}, "properties": {"call_id": {"type": "string"}, "id": {"type": "string"}, "operation": {"$ref": "#/components/schemas/ApplyPatchCallOperation"}, "status": {"$ref": "#/components/schemas/ApplyPatchCallStatus"}, "type": {"enum": ["apply_patch_call"], "type": "string"}}, "required": ["type", "id", "call_id", "status", "operation"], "type": "object"}
```
