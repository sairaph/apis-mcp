---
title: OpenResponsesLogProbs
page_id: schema-openresponseslogprobs-f54cb630
path: schemas
description: Log probability information for a token
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OpenResponsesLogProbs

Log probability information for a token

```yaml
{"description": "Log probability information for a token", "example": {"logprob": -0.1, "token": "world", "top_logprobs": [{"logprob": -0.5, "token": "hello"}]}, "properties": {"bytes": {"items": {"type": "integer"}, "type": "array"}, "logprob": {"format": "double", "type": "number"}, "token": {"type": "string"}, "top_logprobs": {"items": {"$ref": "#/components/schemas/OpenResponsesTopLogprobs"}, "type": "array"}}, "required": ["logprob", "token"], "type": "object"}
```
