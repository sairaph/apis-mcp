---
title: Fund a test mode cash balance
page_id: operation-post-v1-test-helpers-customers-customer-fund-cash-balance-a7c56ab5
path: operations/untagged
description: <p>Create an incoming testmode bank transfer</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/customers/{customer}/fund_cash_balance
operation_ids:
    - PostTestHelpersCustomersCustomerFundCashBalance
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Fund a test mode cash balance

`POST /v1/test_helpers/customers/{customer}/fund_cash_balance`

Operation ID: `PostTestHelpersCustomersCustomerFundCashBalance`

<p>Create an incoming testmode bank transfer</p>

## Definition

```yaml
{"summary": "Fund a test mode cash balance", "description": "<p>Create an incoming testmode bank transfer</p>", "operationId": "PostTestHelpersCustomersCustomerFundCashBalance", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to be used for this test cash balance transaction. A positive integer representing how much to fund in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to fund $1.00 or 100 to fund ¥100, a zero-decimal currency)."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "reference": {"maxLength": 5000, "type": "string", "description": "A description of the test funding. This simulates free-text references supplied by customers when making bank transfers to their cash balance. You can use this to test how Stripe's [reconciliation algorithm](https://docs.stripe.com/payments/customer-balance/reconciliation) applies to different user inputs."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/customer_cash_balance_transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
