---
title: BadGatewayResponse
page_id: schema-badgatewayresponse-73a1369b
path: schemas
description: Bad Gateway - Provider/upstream API failure
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BadGatewayResponse

Bad Gateway - Provider/upstream API failure

```yaml
{"description": "Bad Gateway - Provider/upstream API failure", "example": {"error": {"code": 502, "message": "Provider returned error"}}, "properties": {"error": {"$ref": "#/components/schemas/BadGatewayResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
