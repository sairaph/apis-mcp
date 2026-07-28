---
title: stream_video_copy_request
page_id: schema-stream-video-copy-request-e316561c
path: schemas
description: Copy upload request. Provide `input` (preferred) or `url` (deprecated).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_video_copy_request

Copy upload request. Provide `input` (preferred) or `url` (deprecated).

```yaml
{"description": "Copy upload request. Provide `input` (preferred) or `url` (deprecated).", "type": "object", "properties": {"allowedOrigins": {"$ref": "#/components/schemas/stream_allowedOrigins"}, "creator": {"$ref": "#/components/schemas/stream_creator"}, "input": {"description": "A video's URL. The server must be publicly routable and support `HTTP HEAD` requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD` requests with a `content-range` header that includes the size of the file. This is the preferred field over `url`.", "type": "string", "format": "uri", "example": "https://example.com/myvideo.mp4"}, "meta": {"$ref": "#/components/schemas/stream_media_metadata"}, "name": {"description": "A video's name. Used for legacy compatibility.", "type": "string", "example": "myvideo.mp4"}, "requireSignedURLs": {"$ref": "#/components/schemas/stream_requireSignedURLs"}, "scheduledDeletion": {"$ref": "#/components/schemas/stream_scheduledDeletion"}, "thumbnailTimestampPct": {"$ref": "#/components/schemas/stream_thumbnailTimestampPct"}, "url": {"description": "A video's URL. The server must be publicly routable and support `HTTP HEAD` requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD` requests with a `content-range` header that includes the size of the file. This field is deprecated in favor of `input`.", "type": "string", "format": "uri", "example": "https://example.com/myvideo.mp4"}, "watermark": {"$ref": "#/components/schemas/stream_watermark_at_upload"}}, "anyOf": [{"required": ["input"]}, {"required": ["url"]}]}
```
