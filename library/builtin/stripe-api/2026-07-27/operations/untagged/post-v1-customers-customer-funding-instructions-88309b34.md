---
title: Create or retrieve funding instructions for a customer cash balance
page_id: operation-post-v1-customers-customer-funding-instructions-a7bd2502
path: operations/untagged
description: |-
    <p>Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
    funding instructions will be created. If funding instructions have already been created for a given customer, the same
    funding instructions will be retrieved. In other words, we will return the same funding instructions each time.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customers/{customer}/funding_instructions
operation_ids:
    - PostCustomersCustomerFundingInstructions
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create or retrieve funding instructions for a customer cash balance

`POST /v1/customers/{customer}/funding_instructions`

Operation ID: `PostCustomersCustomerFundingInstructions`

<p>Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new
funding instructions will be created. If funding instructions have already been created for a given customer, the same
funding instructions will be retrieved. In other words, we will return the same funding instructions each time.</p>

## Definition

```yaml
{"summary": "Create or retrieve funding instructions for a customer cash balance", "description": "<p>Retrieve funding instructions for a customer cash balance. If funding instructions do not yet exist for the customer, new\nfunding instructions will be created. If funding instructions have already been created for a given customer, the same\nfunding instructions will be retrieved. In other words, we will return the same funding instructions each time.</p>", "operationId": "PostCustomersCustomerFundingInstructions", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["bank_transfer", "currency", "funding_type"], "type": "object", "properties": {"bank_transfer": {"title": "bank_transfer_params", "required": ["type"], "type": "object", "properties": {"eu_bank_transfer": {"title": "eu_bank_account_params", "required": ["country"], "type": "object", "properties": {"country": {"maxLength": 5000, "type": "string"}}}, "requested_address_types": {"type": "array", "items": {"type": "string", "enum": ["iban", "sort_code", "spei", "zengin"], "x-stripeBypassValidation": true}}, "type": {"type": "string", "enum": ["eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer"], "x-stripeBypassValidation": true}}, "description": "Additional parameters for `bank_transfer` funding types"}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "funding_type": {"type": "string", "description": "The `funding_type` to get the instructions for.", "enum": ["bank_transfer"]}}, "additionalProperties": false}, "encoding": {"bank_transfer": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/funding_instructions"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
