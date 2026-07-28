---
title: TextDoneEvent
page_id: schema-textdoneevent-ed5121a9
path: schemas
description: Event emitted when text streaming is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TextDoneEvent

Event emitted when text streaming is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseTextDoneEvent"}, {"properties": {"logprobs": {"items": {"$ref": "#/components/schemas/StreamLogprob"}, "type": "array"}}, "type": "object"}], "description": "Event emitted when text streaming is complete", "example": {"content_index": 0, "item_id": "item-1", "logprobs": [], "output_index": 0, "sequence_number": 6, "text": "Hello! How can I help you?", "type": "response.output_text.done"}}
```
