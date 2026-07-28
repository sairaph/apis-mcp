---
title: ChatTokenLogprobs
page_id: schema-chattokenlogprobs-49dd7221
path: schemas
description: Log probabilities for the completion
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatTokenLogprobs

Log probabilities for the completion

```yaml
{"description": "Log probabilities for the completion", "example": {"content": [{"bytes": null, "logprob": -0.612345, "token": " Hello", "top_logprobs": []}], "refusal": null}, "properties": {"content": {"description": "Log probabilities for content tokens", "items": {"$ref": "#/components/schemas/ChatTokenLogprob"}, "type": ["array", "null"]}, "refusal": {"description": "Log probabilities for refusal tokens", "items": {"$ref": "#/components/schemas/ChatTokenLogprob"}, "type": ["array", "null"]}}, "required": ["content"], "type": ["object", "null"]}
```
