---
title: BaseAnnotationAddedEvent
page_id: schema-baseannotationaddedevent-9140df40
path: schemas
description: Event emitted when a text annotation is added to output
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BaseAnnotationAddedEvent

Event emitted when a text annotation is added to output

```yaml
{"description": "Event emitted when a text annotation is added to output", "example": {"annotation": {"end_index": 7, "start_index": 0, "title": "Example", "type": "url_citation", "url": "https://example.com"}, "annotation_index": 0, "content_index": 0, "item_id": "item-1", "output_index": 0, "sequence_number": 5, "type": "response.output_text.annotation.added"}, "properties": {"annotation": {"$ref": "#/components/schemas/OpenAIResponsesAnnotation"}, "annotation_index": {"type": "integer"}, "content_index": {"type": "integer"}, "item_id": {"type": "string"}, "output_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.output_text.annotation.added"], "type": "string"}}, "required": ["type", "output_index", "item_id", "content_index", "sequence_number", "annotation_index", "annotation"], "type": "object"}
```
