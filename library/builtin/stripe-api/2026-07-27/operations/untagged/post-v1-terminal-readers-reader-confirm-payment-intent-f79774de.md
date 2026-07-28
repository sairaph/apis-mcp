---
title: Confirm a PaymentIntent on the Reader
page_id: operation-post-v1-terminal-readers-reader-confirm-payment-intent-4da403cf
path: operations/untagged
description: <p>Finalizes a payment on a Reader. See <a href="/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent">Confirming a Payment</a> for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/confirm_payment_intent
operation_ids:
    - PostTerminalReadersReaderConfirmPaymentIntent
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Confirm a PaymentIntent on the Reader

`POST /v1/terminal/readers/{reader}/confirm_payment_intent`

Operation ID: `PostTerminalReadersReaderConfirmPaymentIntent`

<p>Finalizes a payment on a Reader. See <a href="/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent">Confirming a Payment</a> for more details.</p>

## Definition

```yaml
{"summary": "Confirm a PaymentIntent on the Reader", "description": "<p>Finalizes a payment on a Reader. See <a href=\"/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#confirm-the-paymentintent\">Confirming a Payment</a> for more details.</p>", "operationId": "PostTerminalReadersReaderConfirmPaymentIntent", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["payment_intent"], "type": "object", "properties": {"confirm_config": {"title": "confirm_config", "type": "object", "properties": {"return_url": {"type": "string"}}, "description": "Configuration overrides for this confirmation, such as surcharge settings and return URL."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The ID of the PaymentIntent to confirm."}}, "additionalProperties": false}, "encoding": {"confirm_config": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
