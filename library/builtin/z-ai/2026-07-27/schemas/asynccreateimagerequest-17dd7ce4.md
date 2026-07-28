---
title: AsyncCreateImageRequest
page_id: schema-asynccreateimagerequest-17dd7ce4
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# AsyncCreateImageRequest

```yaml
{"type": "object", "required": ["model", "prompt"], "properties": {"model": {"type": "string", "description": "Model code", "enum": ["glm-image"], "example": "glm-image"}, "prompt": {"type": "string", "description": "The text description of the image to be generated.", "example": "A cute little kitten."}, "quality": {"type": "string", "description": "The quality of the generated image. `hd`: Generates a more detailed and rich image with higher overall consistency, takes about `20` seconds.", "enum": ["hd"], "default": "hd"}, "size": {"type": "string", "description": "Image size, recommended enum values: `1280x1280` (default), `1568x1056`, `1056x1568`, `1472x1088`, `1088x1472`, `1728x960`, `960x1728`.\nCustom parameter: Both width and height must be between `1024px-2048px`, and must be divisible by `32`, and the maximum pixel count must not exceed `2^22px`.", "default": "1280x1280", "example": "1280x1280"}, "user_id": {"type": "string", "description": "Unique ID of the end user, helping the platform intervene in illegal activities, inappropriate content generation, or other abuses. ID length: 6 to 128 characters.", "minLength": 6, "maxLength": 128}}}
```
