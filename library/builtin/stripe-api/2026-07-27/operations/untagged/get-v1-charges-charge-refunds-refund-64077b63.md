---
title: GET /v1/charges/{charge}/refunds/{refund}
page_id: operation-get-v1-charges-charge-refunds-refund-ea6863e8
path: operations/untagged
description: <p>Retrieves the details of an existing refund.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/charges/{charge}/refunds/{refund}
operation_ids:
    - GetChargesChargeRefundsRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# GET /v1/charges/{charge}/refunds/{refund}

`GET /v1/charges/{charge}/refunds/{refund}`

Operation ID: `GetChargesChargeRefundsRefund`

<p>Retrieves the details of an existing refund.</p>

## Definition

```yaml
{"description": "<p>Retrieves the details of an existing refund.</p>", "operationId": "GetChargesChargeRefundsRefund", "parameters": [{"name": "charge", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "refund", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
