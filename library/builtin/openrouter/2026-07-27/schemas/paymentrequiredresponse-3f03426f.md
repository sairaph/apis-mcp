---
title: PaymentRequiredResponse
page_id: schema-paymentrequiredresponse-3f03426f
path: schemas
description: Payment Required - Insufficient credits or quota to complete request
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PaymentRequiredResponse

Payment Required - Insufficient credits or quota to complete request

```yaml
{"description": "Payment Required - Insufficient credits or quota to complete request", "example": {"error": {"code": 402, "message": "Insufficient credits. Add more using https://openrouter.ai/credits"}}, "properties": {"error": {"$ref": "#/components/schemas/PaymentRequiredResponseErrorData"}, "openrouter_metadata": {"additionalProperties": {}, "type": ["object", "null"]}, "user_id": {"type": ["string", "null"]}}, "required": ["error"], "type": "object"}
```
