---
title: Speech to Text
page_id: operation-post-paas-v4-audio-transcriptions-6c8b2f5b
path: operations/untagged
description: Use the [GLM-ASR-2512](/guides/audio/glm-asr-2512) model to transcribe audio files into text, supporting multiple languages and real-time streaming transcription.
source: https://docs.z.ai/openapi.json
http_methods:
    - POST
api_endpoints:
    - /paas/v4/audio/transcriptions
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# Speech to Text

`POST /paas/v4/audio/transcriptions`

Use the [GLM-ASR-2512](/guides/audio/glm-asr-2512) model to transcribe audio files into text, supporting multiple languages and real-time streaming transcription.

## Definition

```yaml
{"summary": "Speech to Text", "description": "Use the [GLM-ASR-2512](/guides/audio/glm-asr-2512) model to transcribe audio files into text, supporting multiple languages and real-time streaming transcription.", "requestBody": {"content": {"multipart/form-data": {"schema": {"$ref": "#/components/schemas/AudioTranscriptionRequest"}, "example": {"model": "glm-asr-2512", "stream": false}}}, "required": true}, "responses": {"200": {"description": "Request processed successfully", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/AudioTranscriptionResponse"}}, "text/event-stream": {"schema": {"$ref": "#/components/schemas/AudioTranscriptionStreamResponse"}}}}, "default": {"description": "Request failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/Error"}}}}}}
```
