---
title: Hand off a PaymentIntent to a Reader and collect card details
page_id: operation-post-v1-terminal-readers-reader-collect-payment-method-02e8197a
path: operations/untagged
description: <p>Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See <a href="/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod">Collecting a Payment method</a> for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/collect_payment_method
operation_ids:
    - PostTerminalReadersReaderCollectPaymentMethod
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Hand off a PaymentIntent to a Reader and collect card details

`POST /v1/terminal/readers/{reader}/collect_payment_method`

Operation ID: `PostTerminalReadersReaderCollectPaymentMethod`

<p>Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See <a href="/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod">Collecting a Payment method</a> for more details.</p>

## Definition

```yaml
{"summary": "Hand off a PaymentIntent to a Reader and collect card details", "description": "<p>Initiates a payment flow on a Reader and updates the PaymentIntent with card details before manual confirmation. See <a href=\"/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=inspect#collect-a-paymentmethod\">Collecting a Payment method</a> for more details.</p>", "operationId": "PostTerminalReadersReaderCollectPaymentMethod", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["payment_intent"], "type": "object", "properties": {"collect_config": {"title": "collect_config", "type": "object", "properties": {"allow_redisplay": {"type": "string", "enum": ["always", "limited", "unspecified"]}, "enable_customer_cancellation": {"type": "boolean"}, "skip_tipping": {"type": "boolean"}, "tipping": {"title": "tipping_config", "type": "object", "properties": {"amount_eligible": {"type": "integer"}}}}, "description": "Configuration overrides for this collection, such as tipping, surcharging, and customer cancellation settings."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The ID of the PaymentIntent to collect a payment method for."}}, "additionalProperties": false}, "encoding": {"collect_config": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
