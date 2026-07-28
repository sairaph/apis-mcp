---
title: ForbiddenResponse
page_id: schema-forbiddenresponse-04aaa5df
path: schemas
description: Forbidden - Authentication successful but insufficient permissions
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ForbiddenResponse

Forbidden - Authentication successful but insufficient permissions

```yaml
{"description": "Forbidden - Authentication successful but insufficient permissions", "example": {"error": {"code": 403, "message": "Only management keys can perform this operation"}}, "properties": {"error": {"$ref": "#/components/schemas/ForbiddenResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
