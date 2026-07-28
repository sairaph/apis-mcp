---
title: BaseReasoningSummaryTextDeltaEvent
page_id: schema-basereasoningsummarytextdeltaevent-238f6381
path: schemas
description: Event emitted when reasoning summary text delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseReasoningSummaryTextDeltaEvent

Event emitted when reasoning summary text delta is streamed

```yaml
{"description": "Event emitted when reasoning summary text delta is streamed", "example": {"delta": "Analyzing", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "summary_index": 0, "type": "response.reasoning_summary_text.delta"}, "properties": {"delta": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "summary_index": {"type": "integer"}, "type": {"enum": ["response.reasoning_summary_text.delta"], "type": "string"}}, "required": ["type", "item_id", "output_index", "summary_index", "delta", "sequence_number"], "type": "object"}
```
