---
title: StreamEventsResponseOutputItemAdded
page_id: schema-streameventsresponseoutputitemadded-33dc328d
path: schemas
description: Event emitted when a new output item is added to the response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StreamEventsResponseOutputItemAdded

Event emitted when a new output item is added to the response

```yaml
{"allOf": [{"$ref": "#/components/schemas/OutputItemAddedEvent"}, {"properties": {"item": {"$ref": "#/components/schemas/OutputItems"}}, "type": "object"}], "description": "Event emitted when a new output item is added to the response", "example": {"item": {"content": [], "id": "item-1", "role": "assistant", "status": "in_progress", "type": "message"}, "output_index": 0, "sequence_number": 2, "type": "response.output_item.added"}}
```
