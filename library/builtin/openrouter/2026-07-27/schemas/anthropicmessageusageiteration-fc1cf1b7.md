---
title: AnthropicMessageUsageIteration
page_id: schema-anthropicmessageusageiteration-fc1cf1b7
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicMessageUsageIteration

```yaml
{"allOf": [{"$ref": "#/components/schemas/AnthropicBaseUsageIteration"}, {"properties": {"model": {"type": "string"}, "type": {"enum": ["message"], "type": "string"}}, "required": ["type"], "type": "object"}], "example": {"cache_creation": null, "cache_creation_input_tokens": 0, "cache_read_input_tokens": 0, "input_tokens": 100, "output_tokens": 50, "type": "message"}}
```
