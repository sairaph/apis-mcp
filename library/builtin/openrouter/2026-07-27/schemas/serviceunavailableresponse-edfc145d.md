---
title: ServiceUnavailableResponse
page_id: schema-serviceunavailableresponse-edfc145d
path: schemas
description: Service Unavailable - Service temporarily unavailable
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ServiceUnavailableResponse

Service Unavailable - Service temporarily unavailable

```yaml
{"description": "Service Unavailable - Service temporarily unavailable", "example": {"error": {"code": 503, "message": "Service temporarily unavailable"}}, "properties": {"error": {"$ref": "#/components/schemas/ServiceUnavailableResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
