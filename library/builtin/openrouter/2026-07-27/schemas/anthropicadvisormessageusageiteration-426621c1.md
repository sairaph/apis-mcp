---
title: AnthropicAdvisorMessageUsageIteration
page_id: schema-anthropicadvisormessageusageiteration-426621c1
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicAdvisorMessageUsageIteration

```yaml
{"allOf": [{"$ref": "#/components/schemas/AnthropicBaseUsageIteration"}, {"properties": {"model": {"type": "string"}, "type": {"enum": ["advisor_message"], "type": "string"}}, "required": ["type", "model"], "type": "object"}], "example": {"cache_creation": null, "cache_creation_input_tokens": 0, "cache_read_input_tokens": 0, "input_tokens": 823, "model": "claude-opus-4-6", "output_tokens": 1612, "type": "advisor_message"}}
```
