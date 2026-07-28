---
title: ApplyPatchCallOperationDiffDeltaEvent
page_id: schema-applypatchcalloperationdiffdeltaevent-b9d11d02
path: schemas
description: Incremental chunk of `operation.diff` for an `apply_patch_call`. Matches OpenAI's streaming shape.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchCallOperationDiffDeltaEvent

Incremental chunk of `operation.diff` for an `apply_patch_call`. Matches OpenAI's streaming shape.

```yaml
{"description": "Incremental chunk of `operation.diff` for an `apply_patch_call`. Matches OpenAI's streaming shape.", "example": {"delta": "+console.log(\"hi\");\n", "item_id": "apc_abc123", "output_index": 0, "sequence_number": 5, "type": "response.apply_patch_call_operation_diff.delta"}, "properties": {"delta": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.apply_patch_call_operation_diff.delta"], "type": "string"}}, "required": ["type", "item_id", "output_index", "delta", "sequence_number"], "type": "object"}
```
