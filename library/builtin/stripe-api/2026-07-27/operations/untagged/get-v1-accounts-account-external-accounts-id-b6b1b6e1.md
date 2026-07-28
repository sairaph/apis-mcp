---
title: Retrieve an external account
page_id: operation-get-v1-accounts-account-external-accounts-id-de3bfb50
path: operations/untagged
description: <p>Retrieve a specified external account for a given account.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/accounts/{account}/external_accounts/{id}
operation_ids:
    - GetAccountsAccountExternalAccountsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an external account

`GET /v1/accounts/{account}/external_accounts/{id}`

Operation ID: `GetAccountsAccountExternalAccountsId`

<p>Retrieve a specified external account for a given account.</p>

## Definition

```yaml
{"summary": "Retrieve an external account", "description": "<p>Retrieve a specified external account for a given account.</p>", "operationId": "GetAccountsAccountExternalAccountsId", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "description": "Unique identifier for the external account to be retrieved.", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/external_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
