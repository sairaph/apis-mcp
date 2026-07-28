---
title: ApplyPatchDeleteFileOperation
page_id: schema-applypatchdeletefileoperation-a58a15bd
path: schemas
description: The `delete_file` variant of an `apply_patch_call.operation`. Identifies the file to remove; no diff is required.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchDeleteFileOperation

The `delete_file` variant of an `apply_patch_call.operation`. Identifies the file to remove; no diff is required.

```yaml
{"description": "The `delete_file` variant of an `apply_patch_call.operation`. Identifies the file to remove; no diff is required.", "example": {"path": "/src/main.ts", "type": "delete_file"}, "properties": {"path": {"type": "string"}, "type": {"enum": ["delete_file"], "type": "string"}}, "required": ["type", "path"], "type": "object"}
```
