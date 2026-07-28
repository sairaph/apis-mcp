---
title: Hand-off a PaymentIntent to a Reader
page_id: operation-post-v1-terminal-readers-reader-process-payment-intent-2729671b
path: operations/untagged
description: <p>Initiates a payment flow on a Reader. See <a href="/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment">process the payment</a> for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/process_payment_intent
operation_ids:
    - PostTerminalReadersReaderProcessPaymentIntent
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Hand-off a PaymentIntent to a Reader

`POST /v1/terminal/readers/{reader}/process_payment_intent`

Operation ID: `PostTerminalReadersReaderProcessPaymentIntent`

<p>Initiates a payment flow on a Reader. See <a href="/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment">process the payment</a> for more details.</p>

## Definition

```yaml
{"summary": "Hand-off a PaymentIntent to a Reader", "description": "<p>Initiates a payment flow on a Reader. See <a href=\"/docs/terminal/payments/collect-card-payment?terminal-sdk-platform=server-driven&process=immediately#process-payment\">process the payment</a> for more details.</p>", "operationId": "PostTerminalReadersReaderProcessPaymentIntent", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["payment_intent"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The ID of the PaymentIntent to process on the reader."}, "process_config": {"title": "process_config", "type": "object", "properties": {"allow_redisplay": {"type": "string", "enum": ["always", "limited", "unspecified"]}, "enable_customer_cancellation": {"type": "boolean"}, "return_url": {"type": "string"}, "skip_tipping": {"type": "boolean"}, "tipping": {"title": "tipping_config", "type": "object", "properties": {"amount_eligible": {"type": "integer"}}}}, "description": "Configuration overrides for this transaction, such as tipping and customer cancellation settings."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "process_config": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
