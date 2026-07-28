---
title: ApplyPatchEngineEnum
page_id: schema-applypatchengineenum-66bb3ae6
path: schemas
description: Which apply_patch engine to use. "auto" (default) uses native passthrough when the endpoint advertises native apply_patch support, otherwise falls back to OpenRouter's HITL validator. "native" forces native passthrough — when the endpoint does not support native, the request falls back to HITL. "openrouter" always runs the HITL validator. Native passthrough streams the diff incrementally via `apply_patch_call_operation_diff.delta` events; HITL buffers the diff for atomic delivery as a single delta.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchEngineEnum

Which apply_patch engine to use. "auto" (default) uses native passthrough when the endpoint advertises native apply_patch support, otherwise falls back to OpenRouter's HITL validator. "native" forces native passthrough — when the endpoint does not support native, the request falls back to HITL. "openrouter" always runs the HITL validator. Native passthrough streams the diff incrementally via `apply_patch_call_operation_diff.delta` events; HITL buffers the diff for atomic delivery as a single delta.

```yaml
{"description": "Which apply_patch engine to use. \"auto\" (default) uses native passthrough when the endpoint advertises native apply_patch support, otherwise falls back to OpenRouter's HITL validator. \"native\" forces native passthrough — when the endpoint does not support native, the request falls back to HITL. \"openrouter\" always runs the HITL validator. Native passthrough streams the diff incrementally via `apply_patch_call_operation_diff.delta` events; HITL buffers the diff for atomic delivery as a single delta.", "enum": ["auto", "native", "openrouter"], "example": "auto", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
