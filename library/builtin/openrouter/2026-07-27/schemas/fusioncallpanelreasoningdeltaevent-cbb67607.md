---
title: FusionCallPanelReasoningDeltaEvent
page_id: schema-fusioncallpanelreasoningdeltaevent-cbb67607
path: schemas
description: Incremental reasoning token from a fusion panel model.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallPanelReasoningDeltaEvent

Incremental reasoning token from a fusion panel model.

```yaml
{"description": "Incremental reasoning token from a fusion panel model.", "example": {"delta": "Considering both sides", "item_id": "st_fusion_abc", "model": "openai/gpt-5", "output_index": 0, "sequence_number": 6, "type": "response.fusion_call.panel.reasoning.delta"}, "properties": {"delta": {"type": "string"}, "item_id": {"type": "string"}, "model": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.panel.reasoning.delta"], "type": "string"}}, "required": ["type", "model", "delta", "output_index", "item_id", "sequence_number"], "type": "object"}
```
