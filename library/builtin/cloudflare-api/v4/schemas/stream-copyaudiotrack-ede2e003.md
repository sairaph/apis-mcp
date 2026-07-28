---
title: stream_copyAudioTrack
page_id: schema-stream-copyaudiotrack-ede2e003
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_copyAudioTrack

```yaml
{"type": "object", "properties": {"label": {"$ref": "#/components/schemas/stream_audio_label"}, "url": {"description": "An audio track URL. The server must be publicly routable and support `HTTP HEAD` requests and `HTTP GET` range requests. The server should respond to `HTTP HEAD` requests with a `content-range` header that includes the size of the file.", "type": "string", "format": "uri", "example": "https://www.examplestorage.com/audio_file.mp3", "x-auditable": true}}, "required": ["label"]}
```
