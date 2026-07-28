---
title: Retrieve a customer balance transaction
page_id: operation-get-v1-customers-customer-balance-transactions-transaction-45dea455
path: operations/untagged
description: <p>Retrieves a specific customer balance transaction that updated the customer’s <a href="/docs/billing/customer/balance">balances</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/customers/{customer}/balance_transactions/{transaction}
operation_ids:
    - GetCustomersCustomerBalanceTransactionsTransaction
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a customer balance transaction

`GET /v1/customers/{customer}/balance_transactions/{transaction}`

Operation ID: `GetCustomersCustomerBalanceTransactionsTransaction`

<p>Retrieves a specific customer balance transaction that updated the customer’s <a href="/docs/billing/customer/balance">balances</a>.</p>

## Definition

```yaml
{"summary": "Retrieve a customer balance transaction", "description": "<p>Retrieves a specific customer balance transaction that updated the customer’s <a href=\"/docs/billing/customer/balance\">balances</a>.</p>", "operationId": "GetCustomersCustomerBalanceTransactionsTransaction", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "transaction", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/customer_balance_transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
