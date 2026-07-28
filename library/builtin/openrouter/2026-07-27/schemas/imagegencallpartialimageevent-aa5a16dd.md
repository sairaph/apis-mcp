---
title: ImageGenCallPartialImageEvent
page_id: schema-imagegencallpartialimageevent-aa5a16dd
path: schemas
description: Image generation call with partial image
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenCallPartialImageEvent

Image generation call with partial image

```yaml
{"allOf": [{"$ref": "#/components/schemas/OpenAIResponsesImageGenCallPartialImage"}, {"properties": {}, "type": "object"}], "description": "Image generation call with partial image", "example": {"item_id": "call-123", "output_index": 0, "partial_image_b64": "base64encodedimage...", "partial_image_index": 0, "sequence_number": 3, "type": "response.image_generation_call.partial_image"}}
```
