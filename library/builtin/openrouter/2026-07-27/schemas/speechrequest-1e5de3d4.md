---
title: SpeechRequest
page_id: schema-speechrequest-1e5de3d4
path: schemas
description: Text-to-speech request input
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SpeechRequest

Text-to-speech request input

```yaml
{"description": "Text-to-speech request input", "example": {"input": "Hello world", "model": "mistralai/voxtral-mini-tts-2603", "response_format": "pcm", "speed": 1, "voice": "en_paul_neutral"}, "properties": {"input": {"description": "Text to synthesize", "example": "Hello world", "type": "string"}, "model": {"description": "TTS model identifier", "example": "mistralai/voxtral-mini-tts-2603", "type": "string"}, "provider": {"description": "Provider-specific passthrough configuration", "properties": {"options": {"$ref": "#/components/schemas/ProviderOptions"}}, "type": "object"}, "response_format": {"default": "pcm", "description": "Audio output format", "enum": ["mp3", "pcm"], "example": "pcm", "type": "string", "x-speakeasy-unknown-values": "allow"}, "speed": {"description": "Playback speed multiplier. Only used by models that support it (e.g. OpenAI TTS). Ignored by other providers.", "example": 1, "format": "double", "type": "number"}, "voice": {"description": "Voice identifier (provider-specific).", "example": "en_paul_neutral", "type": "string"}}, "required": ["model", "input", "voice"], "type": "object"}
```
