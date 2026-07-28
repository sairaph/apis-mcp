---
title: StreamEventsResponseOutputItemDone
page_id: schema-streameventsresponseoutputitemdone-bbb3eb64
path: schemas
description: Event emitted when an output item is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StreamEventsResponseOutputItemDone

Event emitted when an output item is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/OutputItemDoneEvent"}, {"properties": {"item": {"$ref": "#/components/schemas/OutputItems"}}, "type": "object"}], "description": "Event emitted when an output item is complete", "example": {"item": {"content": [{"annotations": [], "text": "Hello! How can I help you?", "type": "output_text"}], "id": "item-1", "role": "assistant", "status": "completed", "type": "message"}, "output_index": 0, "sequence_number": 8, "type": "response.output_item.done"}}
```
