---
title: RequestTimeoutResponse
page_id: schema-requesttimeoutresponse-64097e4f
path: schemas
description: Request Timeout - Operation exceeded time limit
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RequestTimeoutResponse

Request Timeout - Operation exceeded time limit

```yaml
{"description": "Request Timeout - Operation exceeded time limit", "example": {"error": {"code": 408, "message": "Operation timed out. Please try again later."}}, "properties": {"error": {"$ref": "#/components/schemas/RequestTimeoutResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
