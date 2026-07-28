---
title: Refund a test-mode transaction
page_id: operation-post-v1-test-helpers-issuing-transactions-transaction-refund-fed785bf
path: operations/untagged
description: <p>Refund a test-mode Transaction.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/transactions/{transaction}/refund
operation_ids:
    - PostTestHelpersIssuingTransactionsTransactionRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Refund a test-mode transaction

`POST /v1/test_helpers/issuing/transactions/{transaction}/refund`

Operation ID: `PostTestHelpersIssuingTransactionsTransactionRefund`

<p>Refund a test-mode Transaction.</p>

## Definition

```yaml
{"summary": "Refund a test-mode transaction", "description": "<p>Refund a test-mode Transaction.</p>", "operationId": "PostTestHelpersIssuingTransactionsTransactionRefund", "parameters": [{"name": "transaction", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "refund_amount": {"type": "integer", "description": "The total amount to attempt to refund. This amount is in the provided currency, or defaults to the cards currency, and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
