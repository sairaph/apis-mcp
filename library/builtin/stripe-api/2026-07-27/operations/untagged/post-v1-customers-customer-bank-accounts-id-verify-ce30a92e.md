---
title: Verify a bank account
page_id: operation-post-v1-customers-customer-bank-accounts-id-verify-e43628a3
path: operations/untagged
description: <p>Verify a specified bank account for a given customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customers/{customer}/bank_accounts/{id}/verify
operation_ids:
    - PostCustomersCustomerBankAccountsIdVerify
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Verify a bank account

`POST /v1/customers/{customer}/bank_accounts/{id}/verify`

Operation ID: `PostCustomersCustomerBankAccountsIdVerify`

<p>Verify a specified bank account for a given customer.</p>

## Definition

```yaml
{"summary": "Verify a bank account", "description": "<p>Verify a specified bank account for a given customer.</p>", "operationId": "PostCustomersCustomerBankAccountsIdVerify", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amounts": {"type": "array", "description": "Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.", "items": {"type": "integer"}}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"amounts": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bank_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
