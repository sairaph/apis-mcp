---
title: InternalServerResponse
page_id: schema-internalserverresponse-749b1978
path: schemas
description: Internal Server Error - Unexpected server error
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# InternalServerResponse

Internal Server Error - Unexpected server error

```yaml
{"description": "Internal Server Error - Unexpected server error", "example": {"error": {"code": 500, "message": "Internal Server Error"}}, "properties": {"error": {"$ref": "#/components/schemas/InternalServerResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
