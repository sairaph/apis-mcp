---
title: Hand-off a SetupIntent to a Reader
page_id: operation-post-v1-terminal-readers-reader-process-setup-intent-ab453ce8
path: operations/untagged
description: <p>Initiates a SetupIntent flow on a Reader. See <a href="/docs/terminal/features/saving-payment-details/save-directly">Save directly without charging</a> for more details.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/process_setup_intent
operation_ids:
    - PostTerminalReadersReaderProcessSetupIntent
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Hand-off a SetupIntent to a Reader

`POST /v1/terminal/readers/{reader}/process_setup_intent`

Operation ID: `PostTerminalReadersReaderProcessSetupIntent`

<p>Initiates a SetupIntent flow on a Reader. See <a href="/docs/terminal/features/saving-payment-details/save-directly">Save directly without charging</a> for more details.</p>

## Definition

```yaml
{"summary": "Hand-off a SetupIntent to a Reader", "description": "<p>Initiates a SetupIntent flow on a Reader. See <a href=\"/docs/terminal/features/saving-payment-details/save-directly\">Save directly without charging</a> for more details.</p>", "operationId": "PostTerminalReadersReaderProcessSetupIntent", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["allow_redisplay", "setup_intent"], "type": "object", "properties": {"allow_redisplay": {"type": "string", "description": "This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow.", "enum": ["always", "limited", "unspecified"]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "process_config": {"title": "process_setup_config", "type": "object", "properties": {"enable_customer_cancellation": {"type": "boolean"}}, "description": "Configuration overrides for this setup, such as MOTO and customer cancellation settings."}, "setup_intent": {"maxLength": 5000, "type": "string", "description": "The ID of the SetupIntent to process on the reader."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "process_config": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
