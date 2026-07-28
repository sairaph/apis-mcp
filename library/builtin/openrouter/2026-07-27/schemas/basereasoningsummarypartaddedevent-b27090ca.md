---
title: BaseReasoningSummaryPartAddedEvent
page_id: schema-basereasoningsummarypartaddedevent-b27090ca
path: schemas
description: Event emitted when a reasoning summary part is added
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseReasoningSummaryPartAddedEvent

Event emitted when a reasoning summary part is added

```yaml
{"description": "Event emitted when a reasoning summary part is added", "example": {"item_id": "item-1", "output_index": 0, "part": {"text": "", "type": "summary_text"}, "sequence_number": 3, "summary_index": 0, "type": "response.reasoning_summary_part.added"}, "properties": {"item_id": {"type": "string"}, "output_index": {"type": "integer"}, "part": {"$ref": "#/components/schemas/ReasoningSummaryText"}, "sequence_number": {"type": "integer"}, "summary_index": {"type": "integer"}, "type": {"enum": ["response.reasoning_summary_part.added"], "type": "string"}}, "required": ["type", "output_index", "item_id", "summary_index", "part", "sequence_number"], "type": "object"}
```
