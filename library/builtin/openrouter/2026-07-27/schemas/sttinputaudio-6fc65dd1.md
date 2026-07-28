---
title: STTInputAudio
page_id: schema-sttinputaudio-6fc65dd1
path: schemas
description: Base64-encoded audio to transcribe
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# STTInputAudio

Base64-encoded audio to transcribe

```yaml
{"description": "Base64-encoded audio to transcribe", "example": {"data": "UklGRiQA...", "format": "wav"}, "properties": {"data": {"description": "Base64-encoded audio data (raw bytes, not a data URI)", "type": "string"}, "format": {"description": "Audio format (e.g., wav, mp3, flac, m4a, ogg, webm, aac). Supported formats vary by provider.", "type": "string"}}, "required": ["data", "format"], "type": "object"}
```
