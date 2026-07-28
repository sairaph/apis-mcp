---
title: RefusalDoneEvent
page_id: schema-refusaldoneevent-31588275
path: schemas
description: Event emitted when refusal streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RefusalDoneEvent

Event emitted when refusal streaming is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseRefusalDoneEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when refusal streaming is complete", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "refusal": "I'm sorry, but I can't assist with that request.", "sequence_number": 6, "type": "response.refusal.done"}}
```
