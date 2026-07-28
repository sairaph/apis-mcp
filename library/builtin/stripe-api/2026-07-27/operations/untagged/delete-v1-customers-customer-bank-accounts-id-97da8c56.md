---
title: Delete a customer source
page_id: operation-delete-v1-customers-customer-bank-accounts-id-d53ea7c6
path: operations/untagged
description: <p>Delete a specified source for a given customer.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/customers/{customer}/bank_accounts/{id}
operation_ids:
    - DeleteCustomersCustomerBankAccountsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a customer source

`DELETE /v1/customers/{customer}/bank_accounts/{id}`

Operation ID: `DeleteCustomersCustomerBankAccountsId`

<p>Delete a specified source for a given customer.</p>

## Definition

```yaml
{"summary": "Delete a customer source", "description": "<p>Delete a specified source for a given customer.</p>", "operationId": "DeleteCustomersCustomerBankAccountsId", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/payment_source"}, {"$ref": "#/components/schemas/deleted_payment_source"}]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
