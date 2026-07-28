---
title: UnprocessableEntityResponse
page_id: schema-unprocessableentityresponse-bf07efe2
path: schemas
description: Unprocessable Entity - Semantic validation failure
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnprocessableEntityResponse

Unprocessable Entity - Semantic validation failure

```yaml
{"description": "Unprocessable Entity - Semantic validation failure", "example": {"error": {"code": 422, "message": "Invalid argument"}}, "properties": {"error": {"$ref": "#/components/schemas/UnprocessableEntityResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
