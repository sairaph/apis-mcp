---
title: ApplyPatchCallOutputItem
page_id: schema-applypatchcalloutputitem-cd0e6bb7
path: schemas
description: The client's echo of an `apply_patch_call` after applying the patch. `output` is an optional human-readable log; `status` is `completed` when the patch was applied successfully, `failed` otherwise.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchCallOutputItem

The client's echo of an `apply_patch_call` after applying the patch. `output` is an optional human-readable log; `status` is `completed` when the patch was applied successfully, `failed` otherwise.

```yaml
{"description": "The client's echo of an `apply_patch_call` after applying the patch. `output` is an optional human-readable log; `status` is `completed` when the patch was applied successfully, `failed` otherwise.", "example": {"call_id": "call_abc123", "output": "Applied patch to /src/main.ts", "status": "completed", "type": "apply_patch_call_output"}, "properties": {"call_id": {"type": "string"}, "id": {"type": ["string", "null"]}, "output": {"type": ["string", "null"]}, "status": {"enum": ["completed", "failed"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["apply_patch_call_output"], "type": "string"}}, "required": ["type", "call_id", "status"], "type": "object"}
```
