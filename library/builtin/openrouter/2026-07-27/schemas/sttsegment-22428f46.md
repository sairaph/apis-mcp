---
title: STTSegment
page_id: schema-sttsegment-22428f46
path: schemas
description: A timestamped transcript segment, returned when response_format is verbose_json
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# STTSegment

A timestamped transcript segment, returned when response_format is verbose_json

```yaml
{"description": "A timestamped transcript segment, returned when response_format is verbose_json", "example": {"avg_logprob": -0.28, "compression_ratio": 1.13, "end": 3.2, "id": 0, "no_speech_prob": 0.01, "seek": 0, "speaker": 0, "start": 0, "temperature": 0, "text": "Hello there.", "tokens": [50364, 2425, 456]}, "properties": {"avg_logprob": {"description": "Average log probability of the segment", "format": "double", "type": "number"}, "compression_ratio": {"description": "Compression ratio of the segment", "format": "double", "type": "number"}, "end": {"description": "Segment end time in seconds", "example": 3.2, "format": "double", "type": "number"}, "id": {"description": "Segment index within the transcript", "example": 0, "type": "integer"}, "no_speech_prob": {"description": "Probability the segment contains no speech", "format": "double", "type": "number"}, "seek": {"description": "Seek offset of the segment", "example": 0, "type": "integer"}, "speaker": {"description": "Speaker index for the segment, present when the provider returns diarization data", "example": 0, "type": "integer"}, "start": {"description": "Segment start time in seconds", "example": 0, "format": "double", "type": "number"}, "temperature": {"description": "Temperature used for the segment", "format": "double", "type": "number"}, "text": {"description": "Transcribed text of the segment", "example": "Hello there.", "type": "string"}, "tokens": {"description": "Token IDs of the segment", "items": {"type": "integer"}, "type": "array"}}, "required": ["id", "start", "end", "text"], "type": "object"}
```
