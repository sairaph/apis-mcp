---
title: STTRequest
page_id: schema-sttrequest-c504b7d5
path: schemas
description: Speech-to-text request input. Accepts a JSON body with input_audio containing base64-encoded audio.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# STTRequest

Speech-to-text request input. Accepts a JSON body with input_audio containing base64-encoded audio.

```yaml
{"description": "Speech-to-text request input. Accepts a JSON body with input_audio containing base64-encoded audio.", "example": {"input_audio": {"data": "UklGRiQA...", "format": "wav"}, "language": "en", "model": "openai/whisper-large-v3"}, "properties": {"input_audio": {"$ref": "#/components/schemas/STTInputAudio"}, "language": {"description": "ISO-639-1 language code (e.g., \"en\", \"ja\"). Auto-detected if omitted.", "example": "en", "type": "string"}, "model": {"description": "STT model identifier", "example": "openai/whisper-large-v3", "type": "string"}, "provider": {"description": "Provider-specific passthrough configuration", "properties": {"options": {"$ref": "#/components/schemas/ProviderOptions"}}, "type": "object"}, "response_format": {"description": "Output format. \"json\" (default) returns { text, usage }. \"verbose_json\" additionally returns task, language, duration, and segment-level timestamps; only supported by OpenAI-compatible providers.", "enum": ["json", "verbose_json"], "example": "json", "type": "string", "x-speakeasy-unknown-values": "allow"}, "temperature": {"description": "Sampling temperature for transcription", "example": 0, "format": "double", "type": "number"}, "timestamp_granularities": {"description": "Timestamp detail levels to include when response_format is \"verbose_json\". \"segment\" returns segment-level timestamps; \"word\" additionally returns word-level timestamps in the words array. Ignored unless response_format is \"verbose_json\".", "example": ["segment"], "items": {"$ref": "#/components/schemas/STTTimestampGranularity"}, "type": "array"}}, "required": ["model", "input_audio"], "type": "object"}
```
