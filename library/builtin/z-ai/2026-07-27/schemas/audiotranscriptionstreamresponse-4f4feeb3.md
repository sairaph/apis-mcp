---
title: AudioTranscriptionStreamResponse
page_id: schema-audiotranscriptionstreamresponse-4f4feeb3
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# AudioTranscriptionStreamResponse

```yaml
{"type": "object", "properties": {"id": {"type": "string", "description": "Task ID"}, "created": {"type": "integer", "format": "int64", "description": "Request creation time, as a `Unix` timestamp in seconds."}, "model": {"type": "string", "description": "Model name"}, "type": {"type": "string", "description": "Audio transcription event type. `transcript.text.delta` indicates transcription in progress, `transcript.text.done` indicates transcription completed."}, "delta": {"type": "string", "description": "Incremental audio transcription information returned by the model."}}}
```
