---
title: Retrieve a Transaction
page_id: operation-get-v1-financial-connections-transactions-transaction-c337a323
path: operations/untagged
description: <p>Retrieves the details of a Financial Connections <code>Transaction</code></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/financial_connections/transactions/{transaction}
operation_ids:
    - GetFinancialConnectionsTransactionsTransaction
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Transaction

`GET /v1/financial_connections/transactions/{transaction}`

Operation ID: `GetFinancialConnectionsTransactionsTransaction`

<p>Retrieves the details of a Financial Connections <code>Transaction</code></p>

## Definition

```yaml
{"summary": "Retrieve a Transaction", "description": "<p>Retrieves the details of a Financial Connections <code>Transaction</code></p>", "operationId": "GetFinancialConnectionsTransactionsTransaction", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "transaction", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/financial_connections.transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
