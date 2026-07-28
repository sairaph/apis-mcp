---
title: Create a Transaction from a Calculation
page_id: operation-post-v1-tax-transactions-create-from-calculation-a589d642
path: operations/untagged
description: <p>Creates a Tax Transaction from a calculation, if that calculation hasn’t expired. Calculations expire after 90 days.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/tax/transactions/create_from_calculation
operation_ids:
    - PostTaxTransactionsCreateFromCalculation
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Transaction from a Calculation

`POST /v1/tax/transactions/create_from_calculation`

Operation ID: `PostTaxTransactionsCreateFromCalculation`

<p>Creates a Tax Transaction from a calculation, if that calculation hasn’t expired. Calculations expire after 90 days.</p>

## Definition

```yaml
{"summary": "Create a Transaction from a Calculation", "description": "<p>Creates a Tax Transaction from a calculation, if that calculation hasn’t expired. Calculations expire after 90 days.</p>", "operationId": "PostTaxTransactionsCreateFromCalculation", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["calculation", "reference"], "type": "object", "properties": {"calculation": {"maxLength": 5000, "type": "string", "description": "Tax Calculation ID to be used as input when creating the transaction."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "posted_at": {"type": "integer", "description": "The Unix timestamp representing when the tax liability is assumed or reduced, which determines the liability posting period and handling in tax liability reports. The timestamp must fall within the `tax_date` and the current time, unless the `tax_date` is scheduled in advance. Defaults to the current time.", "format": "unix-time"}, "reference": {"maxLength": 500, "type": "string", "description": "A custom order or sale identifier, such as 'myOrder_123'. Must be unique across all transactions, including reversals."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tax.transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
