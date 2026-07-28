---
title: Create a test-mode settlement
page_id: operation-post-v1-test-helpers-issuing-settlements-4118afad
path: operations/untagged
description: <p>Allows the user to create an Issuing settlement.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/settlements
operation_ids:
    - PostTestHelpersIssuingSettlements
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a test-mode settlement

`POST /v1/test_helpers/issuing/settlements`

Operation ID: `PostTestHelpersIssuingSettlements`

<p>Allows the user to create an Issuing settlement.</p>

## Definition

```yaml
{"summary": "Create a test-mode settlement", "description": "<p>Allows the user to create an Issuing settlement.</p>", "operationId": "PostTestHelpersIssuingSettlements", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["bin", "clearing_date", "currency", "net_total_amount"], "type": "object", "properties": {"bin": {"maxLength": 5000, "type": "string", "description": "The Bank Identification Number reflecting this settlement record."}, "clearing_date": {"type": "integer", "description": "The date that the transactions are cleared and posted to user's accounts."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "interchange_fees_amount": {"type": "integer", "description": "The total interchange received as reimbursement for the transactions."}, "net_total_amount": {"type": "integer", "description": "The total net amount required to settle with the network."}, "network": {"type": "string", "description": "The card network for this settlement. One of [\"visa\", \"maestro\", \"mastercard\"]", "enum": ["maestro", "visa"], "x-stripeBypassValidation": true}, "network_settlement_identifier": {"maxLength": 5000, "type": "string", "description": "The Settlement Identification Number assigned by the network."}, "transaction_amount": {"type": "integer", "description": "The total transaction amount reflected in this settlement."}, "transaction_count": {"type": "integer", "description": "The total number of transactions reflected in this settlement."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.settlement"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
