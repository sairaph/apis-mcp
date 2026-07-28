---
title: BaseReasoningSummaryTextDoneEvent
page_id: schema-basereasoningsummarytextdoneevent-41ece609
path: schemas
description: Event emitted when reasoning summary text streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseReasoningSummaryTextDoneEvent

Event emitted when reasoning summary text streaming is complete

```yaml
{"description": "Event emitted when reasoning summary text streaming is complete", "example": {"item_id": "item-1", "output_index": 0, "sequence_number": 6, "summary_index": 0, "text": "Analyzing the problem step by step to find the optimal solution.", "type": "response.reasoning_summary_text.done"}, "properties": {"item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "summary_index": {"type": "integer"}, "text": {"type": "string"}, "type": {"enum": ["response.reasoning_summary_text.done"], "type": "string"}}, "required": ["type", "item_id", "output_index", "summary_index", "text", "sequence_number"], "type": "object"}
```
