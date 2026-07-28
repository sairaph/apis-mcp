---
title: ViduFrames2VideoRequest
page_id: schema-viduframes2videorequest-a1ae4cf2
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ViduFrames2VideoRequest

```yaml
{"allOf": [{"type": "object", "properties": {"model": {"type": "string", "description": "The model code to be called.", "enum": ["viduq1-start-end", "vidu2-start-end"]}, "prompt": {"type": "string", "description": "Text description of the video, maximum input length of 512 characters. Either image_url or prompt must be provided, or both."}, "image_url": {"type": "array", "description": "Images\nSupports input of two images: the first uploaded image will be treated as the first frame, and the second image as the last frame. The model will use the images provided in this parameter to generate a video.\nThe resolutions of the two input images (first and last frame) must be similar, with the ratio between the resolution of the first frame and the resolution of the last frame falling within `0.8–1.25`. Additionally, the image aspect ratio must be less than `1:4` or `4:1`.\nSupports image URLs or images encoded in Base64 (ensure accessibility; using image URLs is recommended).\nSupported formats: `png`, `jpeg`, `.jpg`, `webp`.\nImage file size must not exceed `50 MB`.\nNote: After Base64 decoding, the byte length must be less than 50 MB, and the encoding must include the appropriate content type string, such as `data:image/png;base64,{base64_encode}`.", "items": {"type": "string", "minLength": 1}, "minItems": 1, "maxItems": 2}, "duration": {"oneOf": [{"title": "viduq1-start-end", "type": "integer", "description": "Video duration parameter.\nDefault: `5` , Optional: `5`.", "example": 5, "enum": [5]}, {"title": "vidu2-start-end", "type": "integer", "description": "Video duration parameter.\nDefault: `4` , Optional: `4`.", "example": 4, "enum": [4]}]}, "size": {"oneOf": [{"title": "viduq1-start-end", "type": "string", "description": "Resolution parameter\nDefault: `1920x1080`, Optional: `1920x1080`", "example": "1920x1080", "enum": ["1920x1080"]}, {"title": "vidu2-start-end", "type": "string", "description": "Resolution parameter\nDefault: `1280x720`, Optional: `1280x720`", "example": "1280x720", "default": "1280x720", "enum": ["1280x720"]}]}, "movement_amplitude": {"type": "string", "description": "Motion amplitude\nDefault: `auto` , Optional values:  `auto` ,`small` ,`medium` ,`large`", "example": "auto", "enum": ["auto", "small", "medium", "large"]}, "with_audio": {"type": "boolean", "description": "Add background music to the generated video."}}, "required": ["model"]}, {"$ref": "#/components/schemas/VideoCommonRequest"}]}
```
