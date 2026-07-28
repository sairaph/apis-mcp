---
title: OpenAIResponsesImageGenCallPartialImage
page_id: schema-openairesponsesimagegencallpartialimage-0132e9d1
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponsesImageGenCallPartialImage

```yaml
{"example": {"item_id": "ig_abc123", "output_index": 0, "partial_image_b64": "iVBORw0KGgo...", "partial_image_index": 0, "sequence_number": 3, "type": "response.image_generation_call.partial_image"}, "properties": {"item_id": {"type": "string"}, "output_index": {"type": "integer"}, "partial_image_b64": {"type": "string"}, "partial_image_index": {"type": "integer"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.image_generation_call.partial_image"], "type": "string"}}, "required": ["type", "item_id", "output_index", "sequence_number", "partial_image_b64", "partial_image_index"], "type": "object"}
```
