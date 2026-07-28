---
title: Create a reversal Transaction
page_id: operation-post-v1-tax-transactions-create-reversal-53a2bc39
path: operations/untagged
description: <p>Partially or fully reverses a previously created <code>Transaction</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/tax/transactions/create_reversal
operation_ids:
    - PostTaxTransactionsCreateReversal
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a reversal Transaction

`POST /v1/tax/transactions/create_reversal`

Operation ID: `PostTaxTransactionsCreateReversal`

<p>Partially or fully reverses a previously created <code>Transaction</code>.</p>

## Definition

```yaml
{"summary": "Create a reversal Transaction", "description": "<p>Partially or fully reverses a previously created <code>Transaction</code>.</p>", "operationId": "PostTaxTransactionsCreateReversal", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["mode", "original_transaction", "reference"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "flat_amount": {"type": "integer", "description": "A flat amount to reverse across the entire transaction, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units) in negative. This value represents the total amount to refund from the transaction, including taxes."}, "line_items": {"type": "array", "description": "The line item amounts to reverse.", "items": {"title": "transaction_line_item_reversal", "required": ["amount", "amount_tax", "original_line_item", "reference"], "type": "object", "properties": {"amount": {"type": "integer"}, "amount_tax": {"type": "integer"}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}}, "original_line_item": {"maxLength": 5000, "type": "string"}, "quantity": {"type": "integer"}, "reference": {"maxLength": 500, "type": "string"}}}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "mode": {"type": "string", "description": "If `partial`, the provided line item or shipping cost amounts are reversed. If `full`, the original transaction is fully reversed.", "enum": ["full", "partial"]}, "original_transaction": {"maxLength": 5000, "type": "string", "description": "The ID of the Transaction to partially or fully reverse."}, "reference": {"maxLength": 500, "type": "string", "description": "A custom identifier for this reversal, such as `myOrder_123-refund_1`, which must be unique across all transactions. The reference helps identify this reversal transaction in exported [tax reports](https://docs.stripe.com/tax/reports)."}, "shipping_cost": {"title": "transaction_shipping_cost_reversal", "required": ["amount", "amount_tax"], "type": "object", "properties": {"amount": {"type": "integer"}, "amount_tax": {"type": "integer"}}, "description": "The shipping cost to reverse."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "line_items": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "shipping_cost": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
