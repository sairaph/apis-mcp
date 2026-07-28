---
title: Create an InboundTransfer
page_id: operation-post-v1-treasury-inbound-transfers-b7d82c10
path: operations/untagged
description: <p>Creates an InboundTransfer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/inbound_transfers
operation_ids:
    - PostTreasuryInboundTransfers
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an InboundTransfer

`POST /v1/treasury/inbound_transfers`

Operation ID: `PostTreasuryInboundTransfers`

<p>Creates an InboundTransfer.</p>

## Definition

```yaml
{"summary": "Create an InboundTransfer", "description": "<p>Creates an InboundTransfer.</p>", "operationId": "PostTreasuryInboundTransfers", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount", "currency", "financial_account", "origin_payment_method"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount (in cents) to be transferred."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "financial_account": {"type": "string", "description": "The FinancialAccount to send funds to."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "origin_payment_method": {"maxLength": 5000, "type": "string", "description": "The origin payment method to be debited for the InboundTransfer."}, "statement_descriptor": {"maxLength": 10, "type": "string", "description": "The complete description that appears on your customers' statements. Maximum 10 characters. Can only include -#.$&*, spaces, and alphanumeric characters."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.inbound_transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
