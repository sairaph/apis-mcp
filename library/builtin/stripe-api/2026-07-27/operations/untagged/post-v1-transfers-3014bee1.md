---
title: Create a transfer
page_id: operation-post-v1-transfers-6f4692d3
path: operations/untagged
description: <p>To send funds from your Stripe account to a connected account, you create a new transfer object. Your <a href="#balance">Stripe balance</a> must be able to cover the transfer amount, or you’ll receive an “Insufficient Funds” error.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/transfers
operation_ids:
    - PostTransfers
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a transfer

`POST /v1/transfers`

Operation ID: `PostTransfers`

<p>To send funds from your Stripe account to a connected account, you create a new transfer object. Your <a href="#balance">Stripe balance</a> must be able to cover the transfer amount, or you’ll receive an “Insufficient Funds” error.</p>

## Definition

```yaml
{"summary": "Create a transfer", "description": "<p>To send funds from your Stripe account to a connected account, you create a new transfer object. Your <a href=\"#balance\">Stripe balance</a> must be able to cover the transfer amount, or you’ll receive an “Insufficient Funds” error.</p>", "operationId": "PostTransfers", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["currency", "destination"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer in cents (or local equivalent) representing how much to transfer."}, "currency": {"type": "string", "description": "Three-letter [ISO code for currency](https://www.iso.org/iso-4217-currency-codes.html) in lowercase. Must be a [supported currency](https://docs.stripe.com/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "destination": {"type": "string", "description": "The ID of a connected Stripe account. <a href=\"/docs/connect/separate-charges-and-transfers\">See the Connect documentation</a> for details."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "source_transaction": {"type": "string", "description": "You can use this parameter to transfer funds from a charge before they are added to your available balance. A pending balance will transfer immediately but the funds will not become available until the original charge becomes available. [See the Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-availability) for details."}, "source_type": {"maxLength": 5000, "type": "string", "description": "The source balance to use for this transfer. One of `bank_account`, `card`, or `fpx`. For most users, this will default to `card`.", "enum": ["bank_account", "card", "fpx"], "x-stripeBypassValidation": true}, "transfer_group": {"type": "string", "description": "A string that identifies this transaction as part of a group. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
