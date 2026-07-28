---
title: stream_direct_upload_request
page_id: schema-stream-direct-upload-request-716a616e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_direct_upload_request

```yaml
{"type": "object", "properties": {"allowedOrigins": {"$ref": "#/components/schemas/stream_allowedOrigins"}, "creator": {"$ref": "#/components/schemas/stream_creator"}, "expiry": {"description": "The date and time after upload when videos will not be accepted.", "type": "string", "format": "date-time", "example": "2021-01-02T02:20:00Z", "default": "Now + 30 minutes", "x-auditable": true}, "maxDurationSeconds": {"$ref": "#/components/schemas/stream_maxDurationSeconds"}, "meta": {"$ref": "#/components/schemas/stream_media_metadata"}, "requireSignedURLs": {"$ref": "#/components/schemas/stream_requireSignedURLs"}, "scheduledDeletion": {"$ref": "#/components/schemas/stream_scheduledDeletion"}, "thumbnailTimestampPct": {"$ref": "#/components/schemas/stream_thumbnailTimestampPct"}, "watermark": {"$ref": "#/components/schemas/stream_watermark_at_upload"}}, "required": ["maxDurationSeconds"]}
```
