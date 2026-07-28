---
title: ReasoningSummaryPartDoneEvent
page_id: schema-reasoningsummarypartdoneevent-8876e3fb
path: schemas
description: Event emitted when a reasoning summary part is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningSummaryPartDoneEvent

Event emitted when a reasoning summary part is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseReasoningSummaryPartDoneEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when a reasoning summary part is complete", "example": {"item_id": "item-1", "output_index": 0, "part": {"text": "Analyzing the problem step by step to find the optimal solution.", "type": "summary_text"}, "sequence_number": 7, "summary_index": 0, "type": "response.reasoning_summary_part.done"}}
```
