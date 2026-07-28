---
title: PerRequestLimits
page_id: schema-perrequestlimits-34892b22
path: schemas
description: Per-request token limits
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PerRequestLimits

Per-request token limits

```yaml
{"description": "Per-request token limits", "example": {"completion_tokens": 1000, "prompt_tokens": 1000}, "properties": {"completion_tokens": {"description": "Maximum completion tokens per request", "example": 1000, "type": "number"}, "prompt_tokens": {"description": "Maximum prompt tokens per request", "example": 1000, "type": "number"}}, "required": ["prompt_tokens", "completion_tokens"], "type": ["object", "null"]}
```
