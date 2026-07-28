---
title: BaseContentPartDoneEvent
page_id: schema-basecontentpartdoneevent-00b435c0
path: schemas
description: Event emitted when a content part is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseContentPartDoneEvent

Event emitted when a content part is complete

```yaml
{"description": "Event emitted when a content part is complete", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "part": {"annotations": [], "text": "Hello! How can I help you?", "type": "output_text"}, "sequence_number": 7, "type": "response.content_part.done"}, "properties": {"content_index": {"type": "integer"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "part": {"anyOf": [{"$ref": "#/components/schemas/ResponseOutputText"}, {"$ref": "#/components/schemas/OpenAIResponsesRefusalContent"}]}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.content_part.done"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "part", "sequence_number"], "type": "object"}
```
