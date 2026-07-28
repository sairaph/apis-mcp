---
title: BaseReasoningSummaryPartDoneEvent
page_id: schema-basereasoningsummarypartdoneevent-cf0b4a33
path: schemas
description: Event emitted when a reasoning summary part is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseReasoningSummaryPartDoneEvent

Event emitted when a reasoning summary part is complete

```yaml
{"description": "Event emitted when a reasoning summary part is complete", "example": {"item_id": "item-1", "output_index": 0, "part": {"text": "Analyzing the problem step by step to find the optimal solution.", "type": "summary_text"}, "sequence_number": 7, "summary_index": 0, "type": "response.reasoning_summary_part.done"}, "properties": {"item_id": {"type": "string"}, "output_index": {"type": "integer"}, "part": {"$ref": "#/components/schemas/ReasoningSummaryText"}, "sequence_number": {"type": "integer"}, "summary_index": {"type": "integer"}, "type": {"enum": ["response.reasoning_summary_part.done"], "type": "string"}}, "required": ["type", "output_index", "item_id", "summary_index", "part", "sequence_number"], "type": "object"}
```
