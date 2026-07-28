---
title: Retrieve a transaction
page_id: operation-get-v1-issuing-transactions-transaction-1035be7d
path: operations/untagged
description: <p>Retrieves an Issuing <code>Transaction</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/issuing/transactions/{transaction}
operation_ids:
    - GetIssuingTransactionsTransaction
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a transaction

`GET /v1/issuing/transactions/{transaction}`

Operation ID: `GetIssuingTransactionsTransaction`

<p>Retrieves an Issuing <code>Transaction</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a transaction", "description": "<p>Retrieves an Issuing <code>Transaction</code> object.</p>", "operationId": "GetIssuingTransactionsTransaction", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "transaction", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
