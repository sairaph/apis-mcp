---
title: BaseContentPartAddedEvent
page_id: schema-basecontentpartaddedevent-f97ffba0
path: schemas
description: Event emitted when a new content part is added to an output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseContentPartAddedEvent

Event emitted when a new content part is added to an output item

```yaml
{"description": "Event emitted when a new content part is added to an output item", "example": {"content_index": 0, "item_id": "item-1", "output_index": 0, "part": {"annotations": [], "text": "", "type": "output_text"}, "sequence_number": 3, "type": "response.content_part.added"}, "properties": {"content_index": {"type": "integer"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "part": {"anyOf": [{"$ref": "#/components/schemas/ResponseOutputText"}, {"$ref": "#/components/schemas/OpenAIResponsesRefusalContent"}]}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.content_part.added"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "part", "sequence_number"], "type": "object"}
```
