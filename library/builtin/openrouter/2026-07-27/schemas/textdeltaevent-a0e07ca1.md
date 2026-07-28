---
title: TextDeltaEvent
page_id: schema-textdeltaevent-a0e07ca1
path: schemas
description: Event emitted when a text delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TextDeltaEvent

Event emitted when a text delta is streamed

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseTextDeltaEvent"}, {"properties": {"logprobs": {"items": {"$ref": "#/components/schemas/StreamLogprob"}, "type": "array"}}, "type": "object"}], "description": "Event emitted when a text delta is streamed", "example": {"content_index": 0, "delta": "Hello", "item_id": "item-1", "logprobs": [], "output_index": 0, "sequence_number": 4, "type": "response.output_text.delta"}}
```
