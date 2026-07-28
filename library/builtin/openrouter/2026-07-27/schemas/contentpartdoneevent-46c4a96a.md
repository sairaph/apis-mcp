---
title: ContentPartDoneEvent
page_id: schema-contentpartdoneevent-46c4a96a
path: schemas
description: Event emitted when a content part is complete
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContentPartDoneEvent

Event emitted when a content part is complete

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseContentPartDoneEvent"}, {"properties": {"part": {"anyOf": [{"$ref": "#/components/schemas/ResponseOutputText"}, {"$ref": "#/components/schemas/ReasoningTextContent"}, {"$ref": "#/components/schemas/OpenAIResponsesRefusalContent"}]}}, "type": "object"}], "description": "Event emitted when a content part is complete", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "part": {"annotations": [], "text": "Hello! How can I help you?", "type": "output_text"}, "sequence_number": 7, "type": "response.content_part.done"}}
```
