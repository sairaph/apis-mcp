---
title: Simulate presenting a payment method
page_id: operation-post-v1-test-helpers-terminal-readers-reader-present-payment-method-019367af
path: operations/untagged
description: <p>Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/terminal/readers/{reader}/present_payment_method
operation_ids:
    - PostTestHelpersTerminalReadersReaderPresentPaymentMethod
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Simulate presenting a payment method

`POST /v1/test_helpers/terminal/readers/{reader}/present_payment_method`

Operation ID: `PostTestHelpersTerminalReadersReaderPresentPaymentMethod`

<p>Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.</p>

## Definition

```yaml
{"summary": "Simulate presenting a payment method", "description": "<p>Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.</p>", "operationId": "PostTestHelpersTerminalReadersReaderPresentPaymentMethod", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount_tip": {"type": "integer", "description": "Simulated on-reader tip amount."}, "card": {"title": "card", "required": ["exp_month", "exp_year", "number"], "type": "object", "properties": {"cvc": {"maxLength": 5000, "type": "string"}, "exp_month": {"type": "integer"}, "exp_year": {"type": "integer"}, "number": {"maxLength": 5000, "type": "string"}}, "description": "Simulated data for the card payment method."}, "card_present": {"title": "card_present", "type": "object", "properties": {"number": {"maxLength": 5000, "type": "string"}}, "description": "Simulated data for the card_present payment method."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "interac_present": {"title": "interac_present", "type": "object", "properties": {"number": {"maxLength": 5000, "type": "string"}}, "description": "Simulated data for the interac_present payment method."}, "type": {"type": "string", "description": "Simulated payment type.", "enum": ["card", "card_present", "interac_present"]}}, "additionalProperties": false}, "encoding": {"card": {"style": "deepObject", "explode": true}, "card_present": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "interac_present": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
