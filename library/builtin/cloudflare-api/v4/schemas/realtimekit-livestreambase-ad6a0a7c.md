---
title: realtimekit_LivestreamBase
page_id: schema-realtimekit-livestreambase-ad6a0a7c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_LivestreamBase

```yaml
{"type": "object", "properties": {"created_at": {"description": "The timestamp at which the livestream was created. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "disabled": {"description": "Specifies if the livestream was disabled.", "type": "boolean"}, "id": {"description": "The livestream ID.", "type": "string"}, "ingest_server": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string"}, "meeting_id": {"description": "ID of the meeting.", "type": "string", "nullable": true}, "name": {"description": "Name of the livestream.", "type": "string", "nullable": true}, "org_id": {"type": "string"}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "status": {"description": "The status of the livestream.", "type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string"}, "updated_at": {"description": "The timestamp at which the livestream was updated. The time is returned in ISO format.", "type": "string", "format": "date-time"}}}
```
