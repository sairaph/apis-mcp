---
title: stream_listAudioTrackResponse
page_id: schema-stream-listaudiotrackresponse-5023493a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_listAudioTrackResponse

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"audio": {"description": "Array of audio tracks for the video.", "type": "array", "items": {"$ref": "#/components/schemas/stream_additionalAudio"}}}}}, "type": "object"}]}
```
