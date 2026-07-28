---
title: ApplyPatchCallOperationDiffDoneEvent
page_id: schema-applypatchcalloperationdiffdoneevent-ee777688
path: schemas
description: Emitted when `operation.diff` streaming completes for an `apply_patch_call`.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ApplyPatchCallOperationDiffDoneEvent

Emitted when `operation.diff` streaming completes for an `apply_patch_call`.

```yaml
{"description": "Emitted when `operation.diff` streaming completes for an `apply_patch_call`.", "example": {"diff": "@@\n+console.log(\"hi\");\n", "item_id": "apc_abc123", "output_index": 0, "sequence_number": 12, "type": "response.apply_patch_call_operation_diff.done"}, "properties": {"diff": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.apply_patch_call_operation_diff.done"], "type": "string"}}, "required": ["type", "item_id", "output_index", "diff", "sequence_number"], "type": "object"}
```
