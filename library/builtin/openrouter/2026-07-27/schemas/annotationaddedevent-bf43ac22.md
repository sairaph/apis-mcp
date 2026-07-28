---
title: AnnotationAddedEvent
page_id: schema-annotationaddedevent-bf43ac22
path: schemas
description: Event emitted when a text annotation is added to output
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnnotationAddedEvent

Event emitted when a text annotation is added to output

```yaml
{"allOf": [{"$ref": "#/components/schemas/BaseAnnotationAddedEvent"}, {"properties": {}, "type": "object"}], "description": "Event emitted when a text annotation is added to output", "example": {"annotation": {"end_index": 7, "start_index": 0, "title": "Example", "type": "url_citation", "url": "https://example.com"}, "annotation_index": 0, "content_index": 0, "item_id": "item-1", "output_index": 0, "sequence_number": 5, "type": "response.output_text.annotation.added"}}
```
