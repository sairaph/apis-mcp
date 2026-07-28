---
title: ReasoningConfig
page_id: schema-reasoningconfig-dc042603
path: schemas
description: Configuration for reasoning mode in the response
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningConfig

Configuration for reasoning mode in the response

```yaml
{"anyOf": [{"allOf": [{"$ref": "#/components/schemas/BaseReasoningConfig"}, {"properties": {"enabled": {"type": ["boolean", "null"]}, "max_tokens": {"type": ["integer", "null"]}}, "type": "object"}]}, {"type": "null"}], "description": "Configuration for reasoning mode in the response", "example": {"enabled": true, "summary": "auto"}}
```
