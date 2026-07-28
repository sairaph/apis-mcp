---
title: FusionCallPanelCompletedEvent
page_id: schema-fusioncallpanelcompletedevent-d41d18d0
path: schemas
description: Emitted when a fusion panel model finishes with its full content.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallPanelCompletedEvent

Emitted when a fusion panel model finishes with its full content.

```yaml
{"description": "Emitted when a fusion panel model finishes with its full content.", "example": {"content": "Full panel response text...", "item_id": "st_fusion_abc", "model": "openai/gpt-5", "output_index": 0, "sequence_number": 20, "type": "response.fusion_call.panel.completed"}, "properties": {"content": {"type": "string"}, "item_id": {"type": "string"}, "model": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.fusion_call.panel.completed"], "type": "string"}}, "required": ["type", "model", "content", "output_index", "item_id", "sequence_number"], "type": "object"}
```
