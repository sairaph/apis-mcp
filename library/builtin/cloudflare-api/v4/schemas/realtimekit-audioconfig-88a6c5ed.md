---
title: realtimekit_AudioConfig
page_id: schema-realtimekit-audioconfig-88a6c5ed
path: schemas
description: Object containing configuration regarding the audio that is being recorded.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_AudioConfig

Object containing configuration regarding the audio that is being recorded.

```yaml
{"description": "Object containing configuration regarding the audio that is being recorded.", "type": "object", "properties": {"channel": {"description": "Audio signal pathway within an audio file that carries a specific sound source.", "type": "string", "default": "stereo", "enum": ["mono", "stereo"]}, "codec": {"description": "Codec using which the recording will be encoded. If VP8/VP9 is selected for videoConfig, changing audioConfig is not allowed. In this case, the codec in the audioConfig is automatically set to vorbis.", "type": "string", "default": "AAC", "enum": ["MP3", "AAC"]}, "export_file": {"description": "Controls whether to export audio file seperately", "type": "boolean", "default": true}}, "title": "AudioConfig"}
```
