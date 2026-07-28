---
title: FusionCallInProgressEvent
page_id: schema-fusioncallinprogressevent-34b837fb
path: schemas
description: Emitted when an openrouter:fusion tool call begins executing.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallInProgressEvent

Emitted when an openrouter:fusion tool call begins executing.

```yaml
{"description": "Emitted when an openrouter:fusion tool call begins executing.", "example": {"item_id": "st_fusion_abc", "output_index": 0, "sequence_number": 3, "type": "response.fusion_call.in_progress"}, "properties": {"item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.in_progress"], "type": "string"}}, "required": ["type", "output_index", "item_id", "sequence_number"], "type": "object"}
```
