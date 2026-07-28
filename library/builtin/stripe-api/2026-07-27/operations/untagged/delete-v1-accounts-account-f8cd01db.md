---
title: Delete an account
page_id: operation-delete-v1-accounts-account-8610e6b8
path: operations/untagged
description: |-
    <p>With <a href="/connect">Connect</a>, you can delete accounts you manage.</p>

    <p>Test-mode accounts can be deleted at any time.</p>

    <p>Live-mode accounts that have access to the standard dashboard and Stripe is responsible for negative account balances cannot be deleted, which includes Standard accounts. All other Live-mode accounts, can be deleted when all <a href="/api/balance/balance_object">balances</a> are zero.</p>

    <p>If you want to delete your own account, use the <a href="https://dashboard.stripe.com/settings/account">account information tab in your account settings</a> instead.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/accounts/{account}
operation_ids:
    - DeleteAccountsAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete an account

`DELETE /v1/accounts/{account}`

Operation ID: `DeleteAccountsAccount`

<p>With <a href="/connect">Connect</a>, you can delete accounts you manage.</p>

<p>Test-mode accounts can be deleted at any time.</p>

<p>Live-mode accounts that have access to the standard dashboard and Stripe is responsible for negative account balances cannot be deleted, which includes Standard accounts. All other Live-mode accounts, can be deleted when all <a href="/api/balance/balance_object">balances</a> are zero.</p>

<p>If you want to delete your own account, use the <a href="https://dashboard.stripe.com/settings/account">account information tab in your account settings</a> instead.</p>

## Definition

```yaml
{"summary": "Delete an account", "description": "<p>With <a href=\"/connect\">Connect</a>, you can delete accounts you manage.</p>\n\n<p>Test-mode accounts can be deleted at any time.</p>\n\n<p>Live-mode accounts that have access to the standard dashboard and Stripe is responsible for negative account balances cannot be deleted, which includes Standard accounts. All other Live-mode accounts, can be deleted when all <a href=\"/api/balance/balance_object\">balances</a> are zero.</p>\n\n<p>If you want to delete your own account, use the <a href=\"https://dashboard.stripe.com/settings/account\">account information tab in your account settings</a> instead.</p>", "operationId": "DeleteAccountsAccount", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
