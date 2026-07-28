---
title: Create an OutboundTransfer
page_id: operation-post-v1-treasury-outbound-transfers-24d37c7f
path: operations/untagged
description: <p>Creates an OutboundTransfer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/outbound_transfers
operation_ids:
    - PostTreasuryOutboundTransfers
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an OutboundTransfer

`POST /v1/treasury/outbound_transfers`

Operation ID: `PostTreasuryOutboundTransfers`

<p>Creates an OutboundTransfer.</p>

## Definition

```yaml
{"summary": "Create an OutboundTransfer", "description": "<p>Creates an OutboundTransfer.</p>", "operationId": "PostTreasuryOutboundTransfers", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "currency", "financial_account"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount (in cents) to be transferred."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "destination_payment_method": {"maxLength": 5000, "type": "string", "description": "The PaymentMethod to use as the payment instrument for the OutboundTransfer."}, "destination_payment_method_data": {"title": "payment_method_data", "required": ["type"], "type": "object", "properties": {"financial_account": {"type": "string"}, "type": {"type": "string", "enum": ["financial_account"]}}, "description": "Hash used to generate the PaymentMethod to be used for this OutboundTransfer. Exclusive with `destination_payment_method`."}, "destination_payment_method_options": {"title": "payment_method_options", "type": "object", "properties": {"us_bank_account": {"anyOf": [{"title": "payment_method_options_param", "type": "object", "properties": {"network": {"type": "string", "enum": ["ach", "us_domestic_wire"]}}}, {"type": "string", "enum": [""]}]}}, "description": "Hash describing payment method configuration details."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "financial_account": {"type": "string", "description": "The FinancialAccount to pull funds from."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "statement_descriptor": {"maxLength": 5000, "type": "string", "description": "Statement descriptor to be shown on the receiving end of an OutboundTransfer. Maximum 10 characters for `ach` transfers or 140 characters for `us_domestic_wire` transfers. The default value is \"transfer\". Can only include -#.$&*, spaces, and alphanumeric characters."}}, "additionalProperties": false}, "encoding": {"destination_payment_method_data": {"style": "deepObject", "explode": true}, "destination_payment_method_options": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.outbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
