---
title: Cancel an OutboundPayment
page_id: operation-post-v1-treasury-outbound-payments-id-cancel-fd88316f
path: operations/untagged
description: <p>Cancel an OutboundPayment.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/outbound_payments/{id}/cancel
operation_ids:
    - PostTreasuryOutboundPaymentsIdCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel an OutboundPayment

`POST /v1/treasury/outbound_payments/{id}/cancel`

Operation ID: `PostTreasuryOutboundPaymentsIdCancel`

<p>Cancel an OutboundPayment.</p>

## Definition

```yaml
{"summary": "Cancel an OutboundPayment", "description": "<p>Cancel an OutboundPayment.</p>", "operationId": "PostTreasuryOutboundPaymentsIdCancel", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.outbound_payment"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
