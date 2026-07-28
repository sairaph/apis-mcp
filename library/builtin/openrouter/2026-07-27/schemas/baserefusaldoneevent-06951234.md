---
title: BaseRefusalDoneEvent
page_id: schema-baserefusaldoneevent-06951234
path: schemas
description: Event emitted when refusal streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseRefusalDoneEvent

Event emitted when refusal streaming is complete

```yaml
{"description": "Event emitted when refusal streaming is complete", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "refusal": "I'm sorry, but I can't assist with that request.", "sequence_number": 6, "type": "response.refusal.done"}, "properties": {"content_index": {"type": "integer"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "refusal": {"type": "string"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.refusal.done"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "refusal", "sequence_number"], "type": "object"}
```
