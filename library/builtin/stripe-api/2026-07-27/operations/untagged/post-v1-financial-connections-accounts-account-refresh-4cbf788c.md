---
title: Refresh Account data
page_id: operation-post-v1-financial-connections-accounts-account-refresh-0ecea2d0
path: operations/untagged
description: <p>Refreshes the data associated with a Financial Connections <code>Account</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/financial_connections/accounts/{account}/refresh
operation_ids:
    - PostFinancialConnectionsAccountsAccountRefresh
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Refresh Account data

`POST /v1/financial_connections/accounts/{account}/refresh`

Operation ID: `PostFinancialConnectionsAccountsAccountRefresh`

<p>Refreshes the data associated with a Financial Connections <code>Account</code>.</p>

## Definition

```yaml
{"summary": "Refresh Account data", "description": "<p>Refreshes the data associated with a Financial Connections <code>Account</code>.</p>", "operationId": "PostFinancialConnectionsAccountsAccountRefresh", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["features"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "features": {"type": "array", "description": "The list of account features that you would like to refresh.", "items": {"type": "string", "enum": ["balance", "ownership", "transactions"], "x-stripeBypassValidation": true}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "features": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/financial_connections.account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
