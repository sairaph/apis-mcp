---
title: FusionCallPanelDeltaEvent
page_id: schema-fusioncallpaneldeltaevent-a5048111
path: schemas
description: Incremental content token from a fusion panel model.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallPanelDeltaEvent

Incremental content token from a fusion panel model.

```yaml
{"description": "Incremental content token from a fusion panel model.", "example": {"delta": "Carbon taxes", "item_id": "st_fusion_abc", "model": "openai/gpt-5", "output_index": 0, "sequence_number": 5, "type": "response.fusion_call.panel.delta"}, "properties": {"delta": {"type": "string"}, "item_id": {"type": "string"}, "model": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.panel.delta"], "type": "string"}}, "required": ["type", "model", "delta", "output_index", "item_id", "sequence_number"], "type": "object"}
```
