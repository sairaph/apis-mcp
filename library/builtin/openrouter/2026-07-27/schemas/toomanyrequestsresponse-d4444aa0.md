---
title: TooManyRequestsResponse
page_id: schema-toomanyrequestsresponse-d4444aa0
path: schemas
description: Too Many Requests - Rate limit exceeded
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TooManyRequestsResponse

Too Many Requests - Rate limit exceeded

```yaml
{"description": "Too Many Requests - Rate limit exceeded", "example": {"error": {"code": 429, "message": "Rate limit exceeded"}}, "properties": {"error": {"$ref": "#/components/schemas/TooManyRequestsResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
