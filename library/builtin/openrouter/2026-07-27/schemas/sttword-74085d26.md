---
title: STTWord
page_id: schema-sttword-74085d26
path: schemas
description: A timestamped word, returned when the provider includes word-level timestamps
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# STTWord

A timestamped word, returned when the provider includes word-level timestamps

```yaml
{"description": "A timestamped word, returned when the provider includes word-level timestamps", "example": {"end": 0.4, "speaker": 0, "start": 0, "word": "Hello"}, "properties": {"end": {"description": "Word end time in seconds", "example": 0.4, "format": "double", "type": "number"}, "speaker": {"description": "Speaker index for the word, present when the provider returns diarization data", "example": 0, "type": "integer"}, "start": {"description": "Word start time in seconds", "example": 0, "format": "double", "type": "number"}, "word": {"description": "The transcribed word", "example": "Hello", "type": "string"}}, "required": ["word", "start", "end"], "type": "object"}
```
