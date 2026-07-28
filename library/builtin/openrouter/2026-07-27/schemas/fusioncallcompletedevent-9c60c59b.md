---
title: FusionCallCompletedEvent
page_id: schema-fusioncallcompletedevent-9c60c59b
path: schemas
description: Emitted when the openrouter:fusion tool call finishes.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallCompletedEvent

Emitted when the openrouter:fusion tool call finishes.

```yaml
{"description": "Emitted when the openrouter:fusion tool call finishes.", "example": {"item_id": "st_fusion_abc", "output_index": 0, "sequence_number": 41, "type": "response.fusion_call.completed"}, "properties": {"item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.completed"], "type": "string"}}, "required": ["type", "output_index", "item_id", "sequence_number"], "type": "object"}
```
