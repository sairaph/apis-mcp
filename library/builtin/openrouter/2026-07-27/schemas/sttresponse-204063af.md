---
title: STTResponse
page_id: schema-sttresponse-204063af
path: schemas
description: STT response containing transcribed text and optional usage statistics
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# STTResponse

STT response containing transcribed text and optional usage statistics

```yaml
{"description": "STT response containing transcribed text and optional usage statistics", "example": {"text": "Hello, this is a test of OpenAI speech-to-text transcription.", "usage": {"cost": 0.000508, "input_tokens": 83, "output_tokens": 30, "seconds": 9.2, "total_tokens": 113}}, "properties": {"duration": {"description": "Duration of the input audio in seconds, present when response_format is verbose_json", "example": 9.2, "format": "double", "type": "number"}, "language": {"description": "Detected or forced language, present when response_format is verbose_json", "example": "english", "type": "string"}, "segments": {"description": "Timestamped transcript segments, present when response_format is verbose_json", "items": {"$ref": "#/components/schemas/STTSegment"}, "type": "array"}, "task": {"description": "The task performed, present when response_format is verbose_json", "example": "transcribe", "type": "string"}, "text": {"description": "The transcribed text", "example": "Hello, this is a test of OpenAI speech-to-text transcription. The weather is sunny today and the temperature is around 72 degrees.", "type": "string"}, "usage": {"$ref": "#/components/schemas/STTUsage"}, "words": {"description": "Timestamped words, present when the provider returns word-level timestamps", "items": {"$ref": "#/components/schemas/STTWord"}, "type": "array"}}, "required": ["text"], "type": "object"}
```
