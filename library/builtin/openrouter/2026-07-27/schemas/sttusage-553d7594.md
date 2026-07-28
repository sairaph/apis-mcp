---
title: STTUsage
page_id: schema-sttusage-553d7594
path: schemas
description: Aggregated usage statistics for the request
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# STTUsage

Aggregated usage statistics for the request

```yaml
{"description": "Aggregated usage statistics for the request", "example": {"cost": 0.000508, "input_tokens": 83, "output_tokens": 30, "seconds": 9.2, "total_tokens": 113}, "properties": {"cost": {"description": "Total cost of the request in USD", "example": 0.000508, "format": "double", "type": "number"}, "input_tokens": {"description": "Number of input tokens billed for this request", "example": 83, "type": "integer"}, "output_tokens": {"description": "Number of output tokens generated", "example": 30, "type": "integer"}, "seconds": {"description": "Duration of the input audio in seconds", "example": 9.2, "format": "double", "type": "number"}, "total_tokens": {"description": "Total number of tokens used (input + output)", "example": 113, "type": "integer"}}, "type": "object"}
```
