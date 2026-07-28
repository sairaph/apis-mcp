---
title: stream_videos
page_id: schema-stream-videos-2cce8c6a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_videos

```yaml
{"type": "object", "properties": {"allowedOrigins": {"$ref": "#/components/schemas/stream_allowedOrigins"}, "clippedFrom": {"description": "The unique identifier of the source video this video was clipped from.", "type": "string", "example": "ea95132c15732412d22c1476fa83f27a"}, "created": {"$ref": "#/components/schemas/stream_created"}, "creator": {"$ref": "#/components/schemas/stream_creator"}, "duration": {"$ref": "#/components/schemas/stream_duration"}, "input": {"$ref": "#/components/schemas/stream_input"}, "liveInput": {"$ref": "#/components/schemas/stream_liveInput"}, "maxDurationSeconds": {"$ref": "#/components/schemas/stream_maxDurationSeconds"}, "maxSizeBytes": {"description": "The maximum size in bytes for the video upload.", "type": "integer", "format": "int64"}, "meta": {"$ref": "#/components/schemas/stream_media_metadata"}, "modified": {"$ref": "#/components/schemas/stream_modified"}, "playback": {"$ref": "#/components/schemas/stream_playback"}, "preview": {"$ref": "#/components/schemas/stream_preview"}, "publicDetails": {"description": "Public details for the video including title, share link, channel link, and logo.", "type": "object", "properties": {"channel_link": {"type": "string", "nullable": true}, "logo": {"type": "string", "nullable": true}, "media_id": {"type": "integer"}, "share_link": {"type": "string", "nullable": true}, "title": {"type": "string", "nullable": true}}}, "readyToStream": {"$ref": "#/components/schemas/stream_readyToStream"}, "readyToStreamAt": {"$ref": "#/components/schemas/stream_readyToStreamAt"}, "requireSignedURLs": {"$ref": "#/components/schemas/stream_requireSignedURLs"}, "scheduledDeletion": {"$ref": "#/components/schemas/stream_scheduledDeletion"}, "size": {"$ref": "#/components/schemas/stream_size"}, "status": {"$ref": "#/components/schemas/stream_media_status"}, "thumbnail": {"$ref": "#/components/schemas/stream_thumbnail_url"}, "thumbnailTimestampPct": {"$ref": "#/components/schemas/stream_thumbnailTimestampPct"}, "uid": {"$ref": "#/components/schemas/stream_identifier"}, "uploadExpiry": {"$ref": "#/components/schemas/stream_oneTimeUploadExpiry"}, "uploaded": {"$ref": "#/components/schemas/stream_uploaded"}, "watermark": {"$ref": "#/components/schemas/stream_watermarks"}}}
```
