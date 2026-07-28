---
title: ApplyPatchCallOperation
page_id: schema-applypatchcalloperation-02da9d17
path: schemas
description: The patch operation requested by an `apply_patch_call`. `create_file` and `update_file` carry a V4A diff; `delete_file` omits it.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchCallOperation

The patch operation requested by an `apply_patch_call`. `create_file` and `update_file` carry a V4A diff; `delete_file` omits it.

```yaml
{"description": "The patch operation requested by an `apply_patch_call`. `create_file` and `update_file` carry a V4A diff; `delete_file` omits it.", "discriminator": {"mapping": {"create_file": "#/components/schemas/ApplyPatchCreateFileOperation", "delete_file": "#/components/schemas/ApplyPatchDeleteFileOperation", "update_file": "#/components/schemas/ApplyPatchUpdateFileOperation"}, "propertyName": "type"}, "example": {"diff": "@@ function main() {\n+  console.log(\"hi\");\n }", "path": "/src/main.ts", "type": "update_file"}, "oneOf": [{"$ref": "#/components/schemas/ApplyPatchCreateFileOperation"}, {"$ref": "#/components/schemas/ApplyPatchUpdateFileOperation"}, {"$ref": "#/components/schemas/ApplyPatchDeleteFileOperation"}]}
```
