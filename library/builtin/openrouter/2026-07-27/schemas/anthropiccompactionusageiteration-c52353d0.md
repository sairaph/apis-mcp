---
title: AnthropicCompactionUsageIteration
page_id: schema-anthropiccompactionusageiteration-c52353d0
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# AnthropicCompactionUsageIteration

```yaml
{"allOf": [{"$ref": "#/components/schemas/AnthropicBaseUsageIteration"}, {"properties": {"type": {"enum": ["compaction"], "type": "string"}}, "required": ["type"], "type": "object"}], "example": {"cache_creation": null, "cache_creation_input_tokens": 0, "cache_read_input_tokens": 0, "input_tokens": 50, "output_tokens": 25, "type": "compaction"}}
```
