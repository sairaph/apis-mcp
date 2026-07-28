---
title: StreamLogprob
page_id: schema-streamlogprob-36ce234f
path: schemas
description: Log probability information for a token
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# StreamLogprob

Log probability information for a token

```yaml
{"allOf": [{"$ref": "#/components/schemas/OpenResponsesLogProbs"}, {"properties": {"top_logprobs": {"items": {"$ref": "#/components/schemas/StreamLogprobTopLogprob"}, "type": "array"}}, "type": "object"}], "description": "Log probability information for a token", "example": {"bytes": [72, 101, 108, 108, 111], "logprob": -0.5, "token": "Hello", "top_logprobs": []}}
```
