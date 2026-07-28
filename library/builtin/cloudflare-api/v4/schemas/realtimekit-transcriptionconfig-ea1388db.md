---
title: realtimekit_TranscriptionConfig
page_id: schema-realtimekit-transcriptionconfig-ea1388db
path: schemas
description: Transcription Configurations
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_TranscriptionConfig

Transcription Configurations

```yaml
{"description": "Transcription Configurations", "type": "object", "properties": {"keywords": {"description": "Adds specific terms to improve accurate detection during transcription.", "type": "array", "items": {"type": "string"}}, "language": {"description": "Specifies the language code for transcription to ensure accurate results.", "type": "string", "default": "en-US", "enum": ["en-US", "en-IN", "de", "hi", "sv", "ru", "pl", "el", "fr", "nl"]}, "profanity_filter": {"description": "Control the inclusion of offensive language in transcriptions.", "type": "boolean", "default": false}}, "title": "TranscriptionConfig"}
```
