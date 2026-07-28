---
title: Verify microdeposits on a PaymentIntent
page_id: operation-post-v1-payment-intents-intent-verify-microdeposits-34c8c293
path: operations/untagged
description: <p>Verifies microdeposits on a PaymentIntent object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_intents/{intent}/verify_microdeposits
operation_ids:
    - PostPaymentIntentsIntentVerifyMicrodeposits
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Verify microdeposits on a PaymentIntent

`POST /v1/payment_intents/{intent}/verify_microdeposits`

Operation ID: `PostPaymentIntentsIntentVerifyMicrodeposits`

<p>Verifies microdeposits on a PaymentIntent object.</p>

## Definition

```yaml
{"summary": "Verify microdeposits on a PaymentIntent", "description": "<p>Verifies microdeposits on a PaymentIntent object.</p>", "operationId": "PostPaymentIntentsIntentVerifyMicrodeposits", "parameters": [{"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amounts": {"type": "array", "description": "Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.", "items": {"type": "integer"}}, "client_secret": {"maxLength": 5000, "type": "string", "description": "The client secret of the PaymentIntent."}, "descriptor_code": {"maxLength": 5000, "type": "string", "description": "A six-character code starting with SM present in the microdeposit sent to the bank account."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"amounts": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
