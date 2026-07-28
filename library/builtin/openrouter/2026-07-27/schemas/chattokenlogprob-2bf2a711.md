---
title: ChatTokenLogprob
page_id: schema-chattokenlogprob-2bf2a711
path: schemas
description: Token log probability information
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatTokenLogprob

Token log probability information

```yaml
{"description": "Token log probability information", "example": {"bytes": null, "logprob": -0.612345, "token": " Hello", "top_logprobs": [{"bytes": null, "logprob": -0.612345, "token": " Hello"}]}, "properties": {"bytes": {"description": "UTF-8 bytes of the token", "items": {"type": "integer"}, "type": ["array", "null"]}, "logprob": {"description": "Log probability of the token", "format": "double", "type": "number"}, "token": {"description": "The token", "type": "string"}, "top_logprobs": {"description": "Top alternative tokens with probabilities", "items": {"properties": {"bytes": {"items": {"type": "integer"}, "type": ["array", "null"]}, "logprob": {"format": "double", "type": "number"}, "token": {"type": "string"}}, "required": ["token", "logprob", "bytes"], "type": "object"}, "type": "array"}}, "required": ["token", "logprob", "bytes", "top_logprobs"], "type": "object"}
```
