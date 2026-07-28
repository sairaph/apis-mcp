---
title: OpenAIResponsesUsage
page_id: schema-openairesponsesusage-e5390826
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenAIResponsesUsage

```yaml
{"example": {"input_tokens": 100, "input_tokens_details": {"cached_tokens": 0}, "output_tokens": 50, "output_tokens_details": {"reasoning_tokens": 0}, "total_tokens": 150}, "properties": {"input_tokens": {"type": "integer"}, "input_tokens_details": {"properties": {"cache_write_tokens": {"type": ["integer", "null"]}, "cached_tokens": {"type": "integer"}}, "required": ["cached_tokens"], "type": "object"}, "output_tokens": {"type": "integer"}, "output_tokens_details": {"properties": {"reasoning_tokens": {"type": "integer"}}, "required": ["reasoning_tokens"], "type": "object"}, "total_tokens": {"type": "integer"}}, "required": ["input_tokens", "input_tokens_details", "output_tokens", "output_tokens_details", "total_tokens"], "type": "object"}
```
