---
title: ViduImage2VideoRequest
page_id: schema-viduimage2videorequest-eb860158
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ViduImage2VideoRequest

```yaml
{"allOf": [{"type": "object", "properties": {"model": {"type": "string", "description": "The model code to be called.", "enum": ["viduq1-image", "vidu2-image"]}, "prompt": {"type": "string", "description": "Text description of the video, maximum input length of 512 characters. Either image_url or prompt must be provided, or both."}, "image_url": {"type": "string", "description": "The model will use the image provided in this parameter as the first frame to generate the video.\nOnly `1` image is supported.\nSupported formats: `png` , `jpeg` , `jpg` , `webp` .\nImage aspect ratio must be less than `1:4` or `4:1`.\nImage file size must not exceed `50MB`.\nNote: After Base64 decoding, the byte length must be less than 50 MB, and the encoding must include the appropriate content type string (e.g., `data:image/png;base64,{base64_encode}`).", "oneOf": [{"title": "Image URL", "type": "string", "format": "uri", "example": "https://example.com/image.png"}, {"title": "Base64 Encoded Image", "type": "string", "format": "byte", "example": "data:image/png;base64, XXX"}]}, "duration": {"oneOf": [{"title": "viduq1-image", "type": "integer", "description": "Video duration parameter.\nDefault: `5` , Optional: `5`.", "example": 5, "enum": [5]}, {"title": "viduq2-image", "type": "integer", "description": "Video duration parameter.\nDefault: `4` , Optional: `4`.", "example": 4, "enum": [4]}]}, "size": {"oneOf": [{"title": "viduq1-image", "type": "string", "description": "Resolution parameter\nDefault: `1920x1080`, Optional: `1920x1080`", "example": "1920x1080", "enum": ["1920x1080"]}, {"title": "viduq2-image", "type": "string", "description": "Resolution parameter\nDefault: `1280x720`, Optional: `1280x720`", "example": "1280x720", "default": "1280x720", "enum": ["1280x720"]}]}, "movement_amplitude": {"type": "string", "description": "Motion amplitude\nDefault: `auto` , Optional values:  `auto` ,`small` ,`medium` ,`large`", "example": "auto", "enum": ["auto", "small", "medium", "large"]}, "with_audio": {"type": "boolean", "description": "Add background music to the generated video."}}, "required": ["model"]}, {"$ref": "#/components/schemas/VideoCommonRequest"}]}
```
