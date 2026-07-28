---
title: POST /v1/charges/{charge}/refunds/{refund}
page_id: operation-post-v1-charges-charge-refunds-refund-858eba99
path: operations/untagged
description: <p>Update a specified refund.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/charges/{charge}/refunds/{refund}
operation_ids:
    - PostChargesChargeRefundsRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# POST /v1/charges/{charge}/refunds/{refund}

`POST /v1/charges/{charge}/refunds/{refund}`

Operation ID: `PostChargesChargeRefundsRefund`

<p>Update a specified refund.</p>

## Definition

```yaml
{"description": "<p>Update a specified refund.</p>", "operationId": "PostChargesChargeRefundsRefund", "parameters": [{"name": "charge", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}, {"name": "refund", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
