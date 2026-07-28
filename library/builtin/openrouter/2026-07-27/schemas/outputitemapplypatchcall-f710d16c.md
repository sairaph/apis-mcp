---
title: OutputItemApplyPatchCall
page_id: schema-outputitemapplypatchcall-f710d16c
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputItemApplyPatchCall

```yaml
{"example": {"call_id": "call_abc123", "id": "apc_abc123", "operation": {"diff": "@@ function main() {\n+  console.log(\"hi\");\n }", "path": "/src/main.ts", "type": "update_file"}, "status": "completed", "type": "apply_patch_call"}, "properties": {"call_id": {"type": "string"}, "created_by": {"type": "string"}, "id": {"type": "string"}, "operation": {"discriminator": {"mapping": {"create_file": "#/components/schemas/ApplyPatchCreateFileOperation", "delete_file": "#/components/schemas/ApplyPatchDeleteFileOperation", "update_file": "#/components/schemas/ApplyPatchUpdateFileOperation"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/ApplyPatchCreateFileOperation"}, {"$ref": "#/components/schemas/ApplyPatchUpdateFileOperation"}, {"$ref": "#/components/schemas/ApplyPatchDeleteFileOperation"}]}, "status": {"enum": ["in_progress", "completed"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["apply_patch_call"], "type": "string"}}, "required": ["type", "id", "call_id", "operation", "status"], "type": "object"}
```
