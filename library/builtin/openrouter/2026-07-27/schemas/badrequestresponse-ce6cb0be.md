---
title: BadRequestResponse
page_id: schema-badrequestresponse-ce6cb0be
path: schemas
description: Bad Request - Invalid request parameters or malformed input
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BadRequestResponse

Bad Request - Invalid request parameters or malformed input

```yaml
{"description": "Bad Request - Invalid request parameters or malformed input", "example": {"error": {"code": 400, "message": "Invalid request parameters"}}, "properties": {"error": {"$ref": "#/components/schemas/BadRequestResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
