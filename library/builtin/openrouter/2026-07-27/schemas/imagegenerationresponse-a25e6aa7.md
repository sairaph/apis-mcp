---
title: ImageGenerationResponse
page_id: schema-imagegenerationresponse-a25e6aa7
path: schemas
description: Image generation response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ImageGenerationResponse

Image generation response

```yaml
{"description": "Image generation response", "example": {"created": 1748372400, "data": [{"b64_json": "<base64-encoded-image>"}], "usage": {"completion_tokens": 4175, "cost": 0.04, "prompt_tokens": 0, "total_tokens": 4175}}, "properties": {"created": {"description": "Unix timestamp (seconds) when the image was generated", "example": 1748372400, "type": "integer"}, "data": {"description": "Generated images", "items": {"properties": {"b64_json": {"description": "Base64-encoded image bytes", "type": "string"}, "media_type": {"description": "Media type (MIME type) of the image, e.g. `image/png`, `image/jpeg`, `image/webp`, `image/svg+xml`. May be omitted if the format could not be determined.", "example": "image/png", "type": "string"}}, "required": ["b64_json"], "type": "object"}, "type": "array"}, "usage": {"$ref": "#/components/schemas/ImageGenerationUsage"}}, "required": ["created", "data"], "type": "object"}
```
