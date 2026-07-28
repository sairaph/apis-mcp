---
title: FusionCallPanelFailedEvent
page_id: schema-fusioncallpanelfailedevent-479f4211
path: schemas
description: Emitted when a fusion panel model fails.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionCallPanelFailedEvent

Emitted when a fusion panel model fails.

```yaml
{"description": "Emitted when a fusion panel model fails.", "example": {"error": "Upstream provider error", "item_id": "st_fusion_abc", "model": "openai/gpt-5", "output_index": 0, "sequence_number": 18, "status_code": 502, "type": "response.fusion_call.panel.failed"}, "properties": {"error": {"type": "string"}, "item_id": {"type": "string"}, "model": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "status_code": {"type": "integer"}, "type": {"enum": ["response.fusion_call.panel.failed"], "type": "string"}}, "required": ["type", "model", "error", "output_index", "item_id", "sequence_number"], "type": "object"}
```
