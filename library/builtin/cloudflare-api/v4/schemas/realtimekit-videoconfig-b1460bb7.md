---
title: realtimekit_VideoConfig
page_id: schema-realtimekit-videoconfig-b1460bb7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_VideoConfig

```yaml
{"type": "object", "properties": {"codec": {"description": "Codec using which the recording will be encoded.", "type": "string", "default": "H264", "enum": ["H264", "VP8"]}, "export_file": {"description": "Controls whether to export video file seperately", "type": "boolean", "default": true}, "height": {"description": "Height of the recording video in pixels", "type": "integer", "example": 720, "default": 720, "maximum": 1920, "minimum": 1}, "watermark": {"description": "Watermark to be added to the recording", "type": "object", "properties": {"position": {"description": "Position of the watermark", "type": "string", "default": "left top", "enum": ["left top", "right top", "left bottom", "right bottom"]}, "size": {"description": "Size of the watermark", "type": "object", "properties": {"height": {"description": "Height of the watermark in px", "type": "integer", "minimum": 1}, "width": {"description": "Width of the watermark in px", "type": "integer", "minimum": 1}}}, "url": {"description": "URL of the watermark image", "type": "string", "format": "uri"}}}, "width": {"description": "Width of the recording video in pixels", "type": "integer", "example": 1280, "default": 1280, "maximum": 1920, "minimum": 1}}, "title": "VideoConfig"}
```
