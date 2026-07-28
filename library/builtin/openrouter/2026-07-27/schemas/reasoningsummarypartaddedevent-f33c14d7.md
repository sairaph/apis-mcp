---
title: ReasoningSummaryPartAddedEvent
page_id: schema-reasoningsummarypartaddedevent-f33c14d7
path: schemas
description: Event emitted when a reasoning summary part is added
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningSummaryPartAddedEvent

Event emitted when a reasoning summary part is added

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseReasoningSummaryPartAddedEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when a reasoning summary part is added", "example": {"item_id": "item-1", "output_index": 0, "part": {"text": "", "type": "summary_text"}, "sequence_number": 3, "summary_index": 0, "type": "response.reasoning_summary_part.added"}}
```
