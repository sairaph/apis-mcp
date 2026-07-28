---
title: CogVideoX3Request
page_id: schema-cogvideox3request-2173736f
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# CogVideoX3Request

```yaml
{"allOf": [{"type": "object", "properties": {"model": {"type": "string", "description": "The model code to be called.", "enum": ["cogvideox-3"]}, "prompt": {"type": "string", "description": "Text description of the video, maximum input length of 512 characters. Either image_url or prompt must be provided, or both."}, "quality": {"type": "string", "description": "Output mode, default is `speed`.\n- `quality`: Prioritizes quality, higher generation quality. \n- `speed`: Prioritizes speed, faster generation time, relatively lower quality.", "example": "speed", "enum": ["speed", "quality"]}, "with_audio": {"type": "boolean", "description": "Whether to generate AI sound effects. Default: `false` (no sound effects).", "example": false}, "image_url": {"type": "array", "description": "Provide an image based on which content will be generated. If this parameter is passed, the system will operate based on this image. Supports passing images via URL or Base64 encoding. Image requirements: images support `.png`, `.jpeg`, `.jpg` formats; image size: no more than `5M`. Either image_url and prompt can be used, or both can be passed simultaneously.\nFirst and last frames: supports inputting two images. The first uploaded image is regarded as the first frame, and the second image is regarded as the last frame. The model will generate the video based on the images passed in this parameter.\nFirst and last frame mode only supports `speed` mode", "items": {"oneOf": [{"title": "Image URL", "type": "string", "format": "uri", "example": "https://example.com/image.png"}, {"title": "Base64 Encoded Image", "type": "string", "format": "byte", "example": "data:image/png;base64, XXX"}]}}, "size": {"type": "string", "description": "Default value: if not specified, the short side of the generated video is 1080 by default, and the long side is determined according to the original image ratio. Maximum support for 4K resolution. Resolution options: \"1280x720\", \"720x1280\", \"1024x1024\", \"1080x1920\", \"2048x1080\", \"3840x2160\"", "example": "1920x1080", "enum": ["1280x720", "720x1280", "1024x1024", "1920x1080", "1080x1920", "2048x1080", "3840x2160"]}, "fps": {"type": "integer", "description": "Video frame rate (FPS), optional values are `30` or `60`. Default: `30`.", "example": 30, "enum": [30, 60]}, "duration": {"type": "integer", "description": "Video duration, default is 5 seconds, supports `5` and `10` seconds.", "example": 5, "enum": [5, 10]}}, "required": ["model"]}, {"$ref": "#/components/schemas/VideoCommonRequest"}]}
```
