---
title: stream_watermark_basic_upload
page_id: schema-stream-watermark-basic-upload-789f6bc0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_watermark_basic_upload

```yaml
{"type": "object", "properties": {"file": {"description": "The image file to upload.", "type": "string", "example": "@/Users/rchen/Downloads/watermark.png", "x-auditable": true}, "name": {"$ref": "#/components/schemas/stream_name"}, "opacity": {"$ref": "#/components/schemas/stream_opacity"}, "padding": {"$ref": "#/components/schemas/stream_padding"}, "position": {"$ref": "#/components/schemas/stream_position"}, "scale": {"$ref": "#/components/schemas/stream_scale"}}, "required": ["file"]}
```
