---
title: UnauthorizedResponse
page_id: schema-unauthorizedresponse-6cfcf4c5
path: schemas
description: Unauthorized - Authentication required or invalid credentials
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnauthorizedResponse

Unauthorized - Authentication required or invalid credentials

```yaml
{"description": "Unauthorized - Authentication required or invalid credentials", "example": {"error": {"code": 401, "message": "Missing Authentication header"}}, "properties": {"error": {"$ref": "#/components/schemas/UnauthorizedResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
