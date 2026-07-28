---
title: BaseTextDoneEvent
page_id: schema-basetextdoneevent-6abd5730
path: schemas
description: Event emitted when text streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseTextDoneEvent

Event emitted when text streaming is complete

```yaml
{"description": "Event emitted when text streaming is complete", "example": {"content_index": 0, "item_id": "item-1", "logprobs": [], "output_index": 0, "sequence_number": 6, "text": "Hello! How can I help you?", "type": "response.output_text.done"}, "properties": {"content_index": {"type": "integer"}, "item_id": {"type": "string"}, "logprobs": {"items": {"$ref": "#/components/schemas/OpenResponsesLogProbs"}, "type": "array"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "text": {"type": "string"}, "type": {"enum": ["response.output_text.done"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "text", "sequence_number", "logprobs"], "type": "object"}
```
