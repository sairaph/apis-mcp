---
title: AudioTranscriptionResponse
page_id: schema-audiotranscriptionresponse-88e80b91
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# AudioTranscriptionResponse

```yaml
{"type": "object", "properties": {"id": {"type": "string", "description": "Task ID"}, "created": {"type": "integer", "format": "int64", "description": "Request creation time, as a `Unix` timestamp in seconds."}, "request_id": {"type": "string", "description": "Passed by the client, must be unique. A unique identifier to distinguish each request. If not provided by the client, the platform will generate one by default."}, "model": {"type": "string", "description": "Model name"}, "text": {"type": "string", "description": "The complete transcribed content of the audio."}}}
```
