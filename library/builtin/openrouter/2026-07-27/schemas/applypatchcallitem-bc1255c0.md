---
title: ApplyPatchCallItem
page_id: schema-applypatchcallitem-bc1255c0
path: schemas
description: A tool call emitted by the model requesting a V4A patch operation. The client applies the patch and echoes an `apply_patch_call_output` on the next turn.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchCallItem

A tool call emitted by the model requesting a V4A patch operation. The client applies the patch and echoes an `apply_patch_call_output` on the next turn.

```yaml
{"description": "A tool call emitted by the model requesting a V4A patch operation. The client applies the patch and echoes an `apply_patch_call_output` on the next turn.", "example": {"call_id": "call_abc123", "id": "apc_abc123", "operation": {"diff": "@@ function main() {\n+  console.log(\"hi\");\n }", "path": "/src/main.ts", "type": "update_file"}, "status": "completed", "type": "apply_patch_call"}, "properties": {"call_id": {"type": "string"}, "id": {"type": ["string", "null"]}, "operation": {"$ref": "#/components/schemas/ApplyPatchCallOperation"}, "status": {"$ref": "#/components/schemas/ApplyPatchCallStatus"}, "type": {"enum": ["apply_patch_call"], "type": "string"}}, "required": ["type", "call_id", "status", "operation"], "type": "object"}
```
