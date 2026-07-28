---
title: OutputApplyPatchServerToolItem
page_id: schema-outputapplypatchservertoolitem-f12b936f
path: schemas
description: An openrouter:apply_patch server tool output item. The turn halts when validation succeeds so the client can apply the patch and echo an `apply_patch_call_output` on the next turn.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputApplyPatchServerToolItem

An openrouter:apply_patch server tool output item. The turn halts when validation succeeds so the client can apply the patch and echo an `apply_patch_call_output` on the next turn.

```yaml
{"description": "An openrouter:apply_patch server tool output item. The turn halts when validation succeeds so the client can apply the patch and echo an `apply_patch_call_output` on the next turn.", "example": {"call_id": "call_abc123", "id": "apc_abc123", "operation": {"diff": "@@ function main() {\n+  console.log(\"hi\");\n }", "path": "/src/main.ts", "type": "update_file"}, "status": "completed", "type": "openrouter:apply_patch"}, "properties": {"call_id": {"type": "string"}, "id": {"type": "string"}, "operation": {"$ref": "#/components/schemas/ApplyPatchCallOperation"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:apply_patch"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
