---
title: FusionCallPanelAddedEvent
page_id: schema-fusioncallpaneladdedevent-100a3806
path: schemas
description: Emitted when a fusion analysis-panel model starts.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallPanelAddedEvent

Emitted when a fusion analysis-panel model starts.

```yaml
{"description": "Emitted when a fusion analysis-panel model starts.", "example": {"item_id": "st_fusion_abc", "model": "openai/gpt-5", "output_index": 0, "sequence_number": 4, "type": "response.fusion_call.panel.added"}, "properties": {"item_id": {"type": "string"}, "model": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.panel.added"], "type": "string"}}, "required": ["type", "model", "output_index", "item_id", "sequence_number"], "type": "object"}
```
