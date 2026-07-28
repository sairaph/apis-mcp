---
title: BaseTextDeltaEvent
page_id: schema-basetextdeltaevent-417fbb03
path: schemas
description: Event emitted when a text delta is streamed
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseTextDeltaEvent

Event emitted when a text delta is streamed

```yaml
{"description": "Event emitted when a text delta is streamed", "example": {"content_index": 0, "delta": "Hello", "item_id": "item-1", "logprobs": [], "output_index": 0, "sequence_number": 4, "type": "response.output_text.delta"}, "properties": {"content_index": {"type": "integer"}, "delta": {"type": "string"}, "item_id": {"type": "string"}, "logprobs": {"items": {"$ref": "#/components/schemas/OpenResponsesLogProbs"}, "type": "array"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.output_text.delta"], "type": "string"}}, "required": ["type", "logprobs", "output_index", "item_id", "content_index", "delta", "sequence_number"], "type": "object"}
```
