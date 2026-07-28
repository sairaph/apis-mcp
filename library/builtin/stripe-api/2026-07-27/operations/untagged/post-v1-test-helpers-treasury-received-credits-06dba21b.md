---
title: 'Test mode: Create a ReceivedCredit'
page_id: operation-post-v1-test-helpers-treasury-received-credits-e022c3fe
path: operations/untagged
description: <p>Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can’t directly create ReceivedCredits initiated by third parties.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/treasury/received_credits
operation_ids:
    - PostTestHelpersTreasuryReceivedCredits
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Test mode: Create a ReceivedCredit

`POST /v1/test_helpers/treasury/received_credits`

Operation ID: `PostTestHelpersTreasuryReceivedCredits`

<p>Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can’t directly create ReceivedCredits initiated by third parties.</p>

## Definition

```yaml
{"summary": "Test mode: Create a ReceivedCredit", "description": "<p>Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can’t directly create ReceivedCredits initiated by third parties.</p>", "operationId": "PostTestHelpersTreasuryReceivedCredits", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "currency", "financial_account", "network"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount (in cents) to be transferred."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "financial_account": {"type": "string", "description": "The FinancialAccount to send funds to."}, "initiating_payment_method_details": {"title": "source_params", "required": ["type"], "type": "object", "properties": {"type": {"type": "string", "enum": ["us_bank_account"]}, "us_bank_account": {"title": "us_bank_account_source_params", "type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string"}, "account_number": {"maxLength": 5000, "type": "string"}, "routing_number": {"maxLength": 5000, "type": "string"}}}}, "description": "Initiating payment method details for the object."}, "network": {"type": "string", "description": "Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.", "enum": ["ach", "us_domestic_wire"], "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "initiating_payment_method_details": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.received_credit"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
