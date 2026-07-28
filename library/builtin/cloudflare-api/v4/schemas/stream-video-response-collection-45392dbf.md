---
title: stream_video_response_collection
page_id: schema-stream-video-response-collection-45392dbf
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_video_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/stream_videos"}}}, "type": "object"}, {"properties": {"range": {"description": "The total number of remaining videos based on cursor position.", "type": "integer", "example": 1000}, "total": {"description": "The total number of videos that match the provided filters.", "type": "integer", "example": 35586}}, "type": "object"}]}
```
