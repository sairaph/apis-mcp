---
title: Retrieve an Account
page_id: operation-get-v1-linked-accounts-account-83242b26
path: operations/untagged
description: <p>Retrieves the details of an Financial Connections <code>Account</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/linked_accounts/{account}
operation_ids:
    - GetLinkedAccountsAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an Account

`GET /v1/linked_accounts/{account}`

Operation ID: `GetLinkedAccountsAccount`

<p>Retrieves the details of an Financial Connections <code>Account</code>.</p>

## Definition

```yaml
{"summary": "Retrieve an Account", "description": "<p>Retrieves the details of an Financial Connections <code>Account</code>.</p>", "operationId": "GetLinkedAccountsAccount", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/financial_connections.account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
