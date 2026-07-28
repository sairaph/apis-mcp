---
title: CreateImageRequest
page_id: schema-createimagerequest-342e9f24
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# CreateImageRequest

```yaml
{"type": "object", "required": ["model", "prompt"], "properties": {"model": {"type": "string", "description": "Model code", "enum": ["glm-image", "cogview-4-250304"], "example": "glm-image"}, "prompt": {"type": "string", "description": "The text description of the image to be generated.", "example": "A cute little kitten."}, "quality": {"type": "string", "description": "The quality of the generated image. `glm-image` default is `hd`, others model is `standard`. `hd`: Generates a more detailed and rich image with higher overall consistency, but takes about `20` seconds. `standard`: Generates an image quickly, suitable for scenarios with higher requirements for generation speed, takes about `5-10` seconds.", "enum": ["hd", "standard"], "default": "hd"}, "size": {"type": "string", "description": "Image size. `glm-image` recommended enum values: `1280x1280` (default), `1568x1056`, `1056x1568`, `1472x1088`, `1088x1472`, `1728x960`, `960x1728`. Custom parameter: Both width and height must be between `1024px-2048px`, and must be divisible by `32`, and the maximum pixel count must not exceed `2^22px`. \nOthers model recommended enum values: `1024x1024` (default), `768x1344`, `864x1152`, `1344x768`, `1152x864`, `1440x720`, `720x1440`. Custom parameter: Both width and height must be between `512px-2048px`, and must be divisible by `16`, and the maximum pixel count must not exceed `2^21px`.", "default": "1280x1280", "example": "1280x1280"}, "user_id": {"type": "string", "description": "Unique ID of the end user, helping the platform intervene in illegal activities, inappropriate content generation, or other abuses. ID length: 6 to 128 characters.", "minLength": 6, "maxLength": 128}}}
```
