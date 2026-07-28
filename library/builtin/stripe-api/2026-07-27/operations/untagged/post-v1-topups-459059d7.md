---
title: Create a top-up
page_id: operation-post-v1-topups-175a53a9
path: operations/untagged
description: <p>Top up the balance of an account</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/topups
operation_ids:
    - PostTopups
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a top-up

`POST /v1/topups`

Operation ID: `PostTopups`

<p>Top up the balance of an account</p>

## Definition

```yaml
{"summary": "Create a top-up", "description": "<p>Top up the balance of an account</p>", "operationId": "PostTopups", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer representing how much to transfer."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies)."}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "payment_method": {"maxLength": 5000, "type": "string", "description": "The ID of a PaymentMethod representing the payment method to be used for the top-up. A PaymentMethod of type `us_bank_account` can be used."}, "payment_method_options": {"title": "payment_method_options_param", "type": "object", "properties": {"us_bank_account": {"title": "us_bank_account_param", "required": ["network"], "type": "object", "properties": {"network": {"type": "string", "enum": ["ach"], "x-stripeBypassValidation": true}}}}, "description": "Payment method-specific configuration for this top-up."}, "source": {"maxLength": 5000, "type": "string", "description": "The ID of a source to transfer funds from. For most users, this should be left unspecified which will use the bank account that was set up in the dashboard for the specified currency. In test mode, this can be a test bank token (see [Testing Top-ups](https://docs.stripe.com/connect/testing#testing-top-ups))."}, "statement_descriptor": {"maxLength": 15, "type": "string", "description": "Extra information about a top-up for the source's bank statement. Limited to 15 ASCII characters."}, "transfer_group": {"type": "string", "description": "A string that identifies this top-up as part of a group."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "payment_method_options": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/topup"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
