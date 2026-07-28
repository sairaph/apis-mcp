---
title: Reject an account
page_id: operation-post-v1-accounts-account-reject-b18459cb
path: operations/untagged
description: |-
    <p>With <a href="/connect">Connect</a>, you can reject accounts that you have flagged as suspicious.</p>

    <p>Only accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be rejected. Test-mode accounts can be rejected at any time. Live-mode accounts can only be rejected after all balances are zero.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/accounts/{account}/reject
operation_ids:
    - PostAccountsAccountReject
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Reject an account

`POST /v1/accounts/{account}/reject`

Operation ID: `PostAccountsAccountReject`

<p>With <a href="/connect">Connect</a>, you can reject accounts that you have flagged as suspicious.</p>

<p>Only accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be rejected. Test-mode accounts can be rejected at any time. Live-mode accounts can only be rejected after all balances are zero.</p>

## Definition

```yaml
{"summary": "Reject an account", "description": "<p>With <a href=\"/connect\">Connect</a>, you can reject accounts that you have flagged as suspicious.</p>\n\n<p>Only accounts where your platform is liable for negative account balances, which includes Custom and Express accounts, can be rejected. Test-mode accounts can be rejected at any time. Live-mode accounts can only be rejected after all balances are zero.</p>", "operationId": "PostAccountsAccountReject", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["reason"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "reason": {"maxLength": 5000, "type": "string", "description": "The reason for rejecting the account. Can be `fraud`, `terms_of_service`, or `other`."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
