---
title: ImageGenCompletedEvent
page_id: schema-imagegencompletedevent-31f1e6e8
path: schemas
description: Emitted when generation completes and the final image is available
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenCompletedEvent

Emitted when generation completes and the final image is available

```yaml
{"description": "Emitted when generation completes and the final image is available", "example": {"b64_json": "<base64-encoded-final-image>", "created": 1748372400, "type": "image_generation.completed", "usage": {"completion_tokens": 4175, "cost": 0.04, "prompt_tokens": 0, "total_tokens": 4175}}, "properties": {"b64_json": {"description": "Base64-encoded final image data", "type": "string"}, "created": {"description": "Unix timestamp (seconds) when the image was generated", "type": "integer"}, "media_type": {"description": "Media type (MIME type) of the image, e.g. `image/png`, `image/jpeg`, `image/webp`, `image/svg+xml`. May be omitted if the format could not be determined.", "example": "image/png", "type": "string"}, "type": {"description": "The event type", "enum": ["image_generation.completed"], "type": "string"}, "usage": {"$ref": "#/components/schemas/ImageGenerationUsage"}}, "required": ["type", "b64_json", "created"], "type": "object"}
```
