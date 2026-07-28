---
title: stream_video_update
page_id: schema-stream-video-update-ab76a6ff
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_video_update

```yaml
{"type": "object", "properties": {"allowedOrigins": {"$ref": "#/components/schemas/stream_allowedOrigins"}, "creator": {"$ref": "#/components/schemas/stream_creator"}, "maxDurationSeconds": {"$ref": "#/components/schemas/stream_maxDurationSeconds"}, "meta": {"$ref": "#/components/schemas/stream_media_metadata"}, "publicDetails": {"description": "Public details for the video including title, share link, channel link, and logo.", "type": "object", "properties": {"channel_link": {"type": "string", "nullable": true}, "logo": {"type": "string", "nullable": true}, "share_link": {"type": "string", "nullable": true}, "title": {"type": "string", "nullable": true}}}, "requireSignedURLs": {"$ref": "#/components/schemas/stream_requireSignedURLs"}, "scheduledDeletion": {"$ref": "#/components/schemas/stream_scheduledDeletion"}, "thumbnailTimestampPct": {"$ref": "#/components/schemas/stream_thumbnailTimestampPct"}, "uid": {"description": "The unique identifier for the video. Can be used to verify the video being updated.", "type": "string", "example": "ea95132c15732412d22c1476fa83f27a"}, "uploadExpiry": {"$ref": "#/components/schemas/stream_oneTimeUploadExpiry"}}}
```
