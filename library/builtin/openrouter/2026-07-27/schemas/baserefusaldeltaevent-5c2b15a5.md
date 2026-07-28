---
title: BaseRefusalDeltaEvent
page_id: schema-baserefusaldeltaevent-5c2b15a5
path: schemas
description: Event emitted when a refusal delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseRefusalDeltaEvent

Event emitted when a refusal delta is streamed

```yaml
{"description": "Event emitted when a refusal delta is streamed", "example": {"content_index": 0, "delta": "I'm sorry", "item_id": "item-1", "output_index": 0, "sequence_number": 4, "type": "response.refusal.delta"}, "properties": {"content_index": {"type": "integer"}, "delta": {"type": "string"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.refusal.delta"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "delta", "sequence_number"], "type": "object"}
```
