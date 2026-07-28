---
title: ChatAudioOutput
page_id: schema-chataudiooutput-d541d29e
path: schemas
description: Audio output data or reference
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatAudioOutput

Audio output data or reference

```yaml
{"description": "Audio output data or reference", "example": {"data": "UklGRnoGAABXQVZFZm10IBAAAAABAAEAQB8AAEAfAAABAAgAZGF0YQoGAACBhYqFbF1f", "expires_at": 1677652400, "id": "audio_abc123", "transcript": "Hello! How can I help you today?"}, "properties": {"data": {"description": "Base64 encoded audio data", "type": "string"}, "expires_at": {"description": "Audio expiration timestamp", "type": "integer"}, "id": {"description": "Audio output identifier", "type": "string"}, "transcript": {"description": "Audio transcript", "type": "string"}}, "type": "object"}
```
