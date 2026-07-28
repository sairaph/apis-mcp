---
title: Subscribe to data refreshes for an Account
page_id: operation-post-v1-financial-connections-accounts-account-subscribe-3b0c9415
path: operations/untagged
description: <p>Subscribes to periodic refreshes of data associated with a Financial Connections <code>Account</code>. When the account status is active, data is typically refreshed once a day.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/financial_connections/accounts/{account}/subscribe
operation_ids:
    - PostFinancialConnectionsAccountsAccountSubscribe
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Subscribe to data refreshes for an Account

`POST /v1/financial_connections/accounts/{account}/subscribe`

Operation ID: `PostFinancialConnectionsAccountsAccountSubscribe`

<p>Subscribes to periodic refreshes of data associated with a Financial Connections <code>Account</code>. When the account status is active, data is typically refreshed once a day.</p>

## Definition

```yaml
{"summary": "Subscribe to data refreshes for an Account", "description": "<p>Subscribes to periodic refreshes of data associated with a Financial Connections <code>Account</code>. When the account status is active, data is typically refreshed once a day.</p>", "operationId": "PostFinancialConnectionsAccountsAccountSubscribe", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["features"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "features": {"type": "array", "description": "The list of account features to which you would like to subscribe.", "items": {"type": "string", "enum": ["transactions"], "x-stripeBypassValidation": true}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "features": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/financial_connections.account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
