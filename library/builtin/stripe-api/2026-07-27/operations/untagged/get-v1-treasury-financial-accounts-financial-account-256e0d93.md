---
title: Retrieve a FinancialAccount
page_id: operation-get-v1-treasury-financial-accounts-financial-account-22d09b5d
path: operations/untagged
description: <p>Retrieves the details of a FinancialAccount.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/treasury/financial_accounts/{financial_account}
operation_ids:
    - GetTreasuryFinancialAccountsFinancialAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a FinancialAccount

`GET /v1/treasury/financial_accounts/{financial_account}`

Operation ID: `GetTreasuryFinancialAccountsFinancialAccount`

<p>Retrieves the details of a FinancialAccount.</p>

## Definition

```yaml
{"summary": "Retrieve a FinancialAccount", "description": "<p>Retrieves the details of a FinancialAccount.</p>", "operationId": "GetTreasuryFinancialAccountsFinancialAccount", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "financial_account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.financial_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
