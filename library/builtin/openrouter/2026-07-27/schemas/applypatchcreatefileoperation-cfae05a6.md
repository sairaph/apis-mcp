---
title: ApplyPatchCreateFileOperation
page_id: schema-applypatchcreatefileoperation-cfae05a6
path: schemas
description: The `create_file` variant of an `apply_patch_call.operation`. Carries a V4A diff describing the new file contents.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchCreateFileOperation

The `create_file` variant of an `apply_patch_call.operation`. Carries a V4A diff describing the new file contents.

```yaml
{"description": "The `create_file` variant of an `apply_patch_call.operation`. Carries a V4A diff describing the new file contents.", "example": {"diff": "@@\n+console.log(\"hi\");\n", "path": "/src/main.ts", "type": "create_file"}, "properties": {"diff": {"type": "string"}, "path": {"type": "string"}, "type": {"enum": ["create_file"], "type": "string"}}, "required": ["type", "path", "diff"], "type": "object"}
```
