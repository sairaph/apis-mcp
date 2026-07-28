---
title: ApplyPatchUpdateFileOperation
page_id: schema-applypatchupdatefileoperation-43ca4cf2
path: schemas
description: The `update_file` variant of an `apply_patch_call.operation`. Carries a V4A diff describing edits to an existing file.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchUpdateFileOperation

The `update_file` variant of an `apply_patch_call.operation`. Carries a V4A diff describing edits to an existing file.

```yaml
{"description": "The `update_file` variant of an `apply_patch_call.operation`. Carries a V4A diff describing edits to an existing file.", "example": {"diff": "@@ function main() {\n+  console.log(\"hi\");\n }", "path": "/src/main.ts", "type": "update_file"}, "properties": {"diff": {"type": "string"}, "path": {"type": "string"}, "type": {"enum": ["update_file"], "type": "string"}}, "required": ["type", "path", "diff"], "type": "object"}
```
