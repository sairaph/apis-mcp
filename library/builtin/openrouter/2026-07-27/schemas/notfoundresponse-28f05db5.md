---
title: NotFoundResponse
page_id: schema-notfoundresponse-28f05db5
path: schemas
description: Not Found - Resource does not exist
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# NotFoundResponse

Not Found - Resource does not exist

```yaml
{"description": "Not Found - Resource does not exist", "example": {"error": {"code": 404, "message": "Resource not found"}}, "properties": {"error": {"$ref": "#/components/schemas/NotFoundResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
