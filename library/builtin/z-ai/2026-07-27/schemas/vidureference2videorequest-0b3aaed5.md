---
title: ViduReference2VideoRequest
page_id: schema-vidureference2videorequest-0b3aaed5
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ViduReference2VideoRequest

```yaml
{"allOf": [{"type": "object", "properties": {"model": {"type": "string", "description": "The model code to be called.", "enum": ["vidu2-reference"]}, "prompt": {"type": "string", "description": "Text description of the video, maximum input length of 512 characters. Either image_url or prompt must be provided, or both."}, "image_url": {"type": "array", "description": "Image reference\nSupports input of 1 to 3 images. The model will use the themes from the images provided in this parameter as references to generate a video with consistent subjects.\n1. Supports image URLs or images encoded in Base64 (ensure accessibility; it is recommended to prioritize using image URLs).\n2. Supported formats: `png`, `jpeg`, `.jpg`, `webp`.\n3. Image resolution must not be smaller than `128x128`, and the aspect ratio must be less than `1:4` or `4:1`.\n4. Image file size must not exceed `50 MB`.\n5. Note: After Base64 decoding, the byte length must be less than 50 MB, and the encoding must include the proper content-type string, such as `data:image/png;base64,{base64_encode}`.", "items": {"type": "string", "minLength": 1}, "minItems": 1, "maxItems": 3}, "duration": {"title": "vidu2-reference", "type": "integer", "description": "Video duration parameter.\nDefault: `4` , Optional: `4`.", "example": 4, "enum": [4]}, "aspect_ratio": {"type": "string", "description": "Aspect ratio\nDefault: `16:9`, Optional values: `16:9`, `9:16`, `1:1`", "example": "16:9", "enum": ["16:9", "9:16", "1:1"]}, "size": {"title": "vidu2-reference ", "type": "string", "description": "Resolution parameter\nDefault: `1280x720`, Optional: `1280x720`", "example": "1280x720", "enum": ["1280x720"]}, "movement_amplitude": {"type": "string", "description": "Motion amplitude\nDefault: `auto` , Optional values:  `auto` ,`small` ,`medium` ,`large`", "example": "auto", "enum": ["auto", "small", "medium", "large"]}, "with_audio": {"type": "boolean", "description": "Add background music to the generated video."}}, "required": ["model"]}, {"$ref": "#/components/schemas/VideoCommonRequest"}]}
```
