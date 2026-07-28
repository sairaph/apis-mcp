---
title: ChatContentAudio
page_id: schema-chatcontentaudio-799ed022
path: schemas
description: Audio input content part. Supported audio formats vary by provider.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatContentAudio

Audio input content part. Supported audio formats vary by provider.

```yaml
{"description": "Audio input content part. Supported audio formats vary by provider.", "example": {"input_audio": {"data": "SGVsbG8gV29ybGQ=", "format": "wav"}, "type": "input_audio"}, "properties": {"input_audio": {"properties": {"data": {"description": "Base64 encoded audio data", "type": "string"}, "format": {"description": "Audio format (e.g., wav, mp3, flac, m4a, ogg, aiff, aac, pcm16, pcm24). Supported formats vary by provider.", "type": "string"}}, "required": ["data", "format"], "type": "object"}, "type": {"enum": ["input_audio"], "type": "string"}}, "required": ["type", "input_audio"], "type": "object"}
```
