---
title: stream_videoClipStandard
page_id: schema-stream-videoclipstandard-3b093773
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_videoClipStandard

```yaml
{"type": "object", "properties": {"allowedOrigins": {"$ref": "#/components/schemas/stream_allowedOrigins"}, "clippedFromVideoUID": {"$ref": "#/components/schemas/stream_clipped_from_video_uid"}, "creator": {"$ref": "#/components/schemas/stream_creator"}, "endTimeSeconds": {"$ref": "#/components/schemas/stream_end_time_seconds"}, "input": {"description": "A video's URL. Preferred over 'url'.", "type": "string", "format": "uri", "example": "https://example.com/myvideo.mp4"}, "meta": {"$ref": "#/components/schemas/stream_media_metadata"}, "name": {"description": "A name for the video.", "type": "string", "example": "myvideo.mp4"}, "requireSignedURLs": {"$ref": "#/components/schemas/stream_requireSignedURLs"}, "scheduledDeletion": {"$ref": "#/components/schemas/stream_scheduledDeletion"}, "startTimeSeconds": {"$ref": "#/components/schemas/stream_start_time_seconds"}, "thumbnailTimestampPct": {"$ref": "#/components/schemas/stream_thumbnailTimestampPct"}, "url": {"description": "A video's URL (legacy field, use 'input' instead).", "type": "string", "format": "uri", "example": "https://example.com/myvideo.mp4"}, "watermark": {"$ref": "#/components/schemas/stream_watermarkAtUpload"}}, "required": ["clippedFromVideoUID", "startTimeSeconds", "endTimeSeconds"]}
```
