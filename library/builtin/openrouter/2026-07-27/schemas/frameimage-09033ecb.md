---
title: FrameImage
page_id: schema-frameimage-09033ecb
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FrameImage

```yaml
{"allOf": [{"$ref": "#/components/schemas/ContentPartImage"}, {"properties": {"frame_type": {"description": "Whether this image represents the first or last frame of the video", "enum": ["first_frame", "last_frame"], "example": "first_frame", "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["frame_type"], "type": "object"}], "example": {"frame_type": "first_frame", "image_url": {"url": "https://example.com/image.png"}, "type": "image_url"}}
```
