---
title: Refund a Charge or a PaymentIntent in-person
page_id: operation-post-v1-terminal-readers-reader-refund-payment-7109d847
path: operations/untagged
description: <p>Initiates an in-person refund on a Reader. See <a href="/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment">Refund an Interac Payment</a> for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/refund_payment
operation_ids:
    - PostTerminalReadersReaderRefundPayment
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Refund a Charge or a PaymentIntent in-person

`POST /v1/terminal/readers/{reader}/refund_payment`

Operation ID: `PostTerminalReadersReaderRefundPayment`

<p>Initiates an in-person refund on a Reader. See <a href="/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment">Refund an Interac Payment</a> for more details.</p>

## Definition

```yaml
{"summary": "Refund a Charge or a PaymentIntent in-person", "description": "<p>Initiates an in-person refund on a Reader. See <a href=\"/docs/terminal/payments/regional?integration-country=CA#refund-an-interac-payment\">Refund an Interac Payment</a> for more details.</p>", "operationId": "PostTerminalReadersReaderRefundPayment", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer in __cents__ representing how much of this charge to refund."}, "charge": {"maxLength": 5000, "type": "string", "description": "ID of the Charge to refund."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "ID of the PaymentIntent to refund."}, "refund_application_fee": {"type": "boolean", "description": "Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge."}, "refund_payment_config": {"title": "refund_payment_config", "type": "object", "properties": {"enable_customer_cancellation": {"type": "boolean"}}, "description": "Configuration overrides for this refund, such as customer cancellation settings."}, "reverse_transfer": {"type": "boolean", "description": "Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount). A transfer can be reversed only by the application that created the charge."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "refund_payment_config": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
