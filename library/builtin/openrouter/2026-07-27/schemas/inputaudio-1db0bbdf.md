---
title: InputAudio
page_id: schema-inputaudio-1db0bbdf
path: schemas
description: Audio input content item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InputAudio

Audio input content item

```yaml
{"description": "Audio input content item", "example": {"input_audio": {"data": "SGVsbG8gV29ybGQ=", "format": "mp3"}, "type": "input_audio"}, "properties": {"input_audio": {"properties": {"data": {"type": "string"}, "format": {"enum": ["mp3", "wav"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["data", "format"], "type": "object"}, "type": {"enum": ["input_audio"], "type": "string"}}, "required": ["type", "input_audio"], "type": "object"}
```
