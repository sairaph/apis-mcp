---
title: ContentPartAddedEvent
page_id: schema-contentpartaddedevent-d93833f6
path: schemas
description: Event emitted when a new content part is added to an output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContentPartAddedEvent

Event emitted when a new content part is added to an output item

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseContentPartAddedEvent"}, {"properties": {"part": {"anyOf": [{"$ref": "#/components/schemas/ResponseOutputText"}, {"$ref": "#/components/schemas/ReasoningTextContent"}, {"$ref": "#/components/schemas/OpenAIResponsesRefusalContent"}]}}, "type": "object"}], "description": "Event emitted when a new content part is added to an output item", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "part": {"annotations": [], "text": "", "type": "output_text"}, "sequence_number": 3, "type": "response.content_part.added"}}
```
