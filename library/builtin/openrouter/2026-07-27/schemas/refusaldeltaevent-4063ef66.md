---
title: RefusalDeltaEvent
page_id: schema-refusaldeltaevent-4063ef66
path: schemas
description: Event emitted when a refusal delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RefusalDeltaEvent

Event emitted when a refusal delta is streamed

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseRefusalDeltaEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when a refusal delta is streamed", "example": {"content_index": 0, "delta": "I'm sorry", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.refusal.delta"}}
```
