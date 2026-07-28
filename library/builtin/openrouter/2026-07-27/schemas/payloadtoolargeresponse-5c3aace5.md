---
title: PayloadTooLargeResponse
page_id: schema-payloadtoolargeresponse-5c3aace5
path: schemas
description: Payload Too Large - Request payload exceeds size limits
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PayloadTooLargeResponse

Payload Too Large - Request payload exceeds size limits

```yaml
{"description": "Payload Too Large - Request payload exceeds size limits", "example": {"error": {"code": 413, "message": "Request payload too large"}}, "properties": {"error": {"$ref": "#/components/schemas/PayloadTooLargeResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
