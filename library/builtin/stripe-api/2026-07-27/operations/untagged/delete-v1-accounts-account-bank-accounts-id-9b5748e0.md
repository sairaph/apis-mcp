---
title: Delete an external account
page_id: operation-delete-v1-accounts-account-bank-accounts-id-87ee1f63
path: operations/untagged
description: <p>Delete a specified external account for a given account.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/accounts/{account}/bank_accounts/{id}
operation_ids:
    - DeleteAccountsAccountBankAccountsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete an external account

`DELETE /v1/accounts/{account}/bank_accounts/{id}`

Operation ID: `DeleteAccountsAccountBankAccountsId`

<p>Delete a specified external account for a given account.</p>

## Definition

```yaml
{"summary": "Delete an external account", "description": "<p>Delete a specified external account for a given account.</p>", "operationId": "DeleteAccountsAccountBankAccountsId", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "id", "in": "path", "description": "Unique identifier for the external account to be deleted.", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_external_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
