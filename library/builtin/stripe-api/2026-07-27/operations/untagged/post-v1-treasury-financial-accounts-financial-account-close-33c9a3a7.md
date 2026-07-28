---
title: Close a FinancialAccount
page_id: operation-post-v1-treasury-financial-accounts-financial-account-close-d17b2805
path: operations/untagged
description: <p>Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/financial_accounts/{financial_account}/close
operation_ids:
    - PostTreasuryFinancialAccountsFinancialAccountClose
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Close a FinancialAccount

`POST /v1/treasury/financial_accounts/{financial_account}/close`

Operation ID: `PostTreasuryFinancialAccountsFinancialAccountClose`

<p>Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.</p>

## Definition

```yaml
{"summary": "Close a FinancialAccount", "description": "<p>Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.</p>", "operationId": "PostTreasuryFinancialAccountsFinancialAccountClose", "parameters": [{"name": "financial_account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "forwarding_settings": {"title": "forwarding_settings", "required": ["type"], "type": "object", "properties": {"financial_account": {"type": "string"}, "payment_method": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["financial_account", "payment_method"]}}, "description": "A different bank account where funds can be deposited/debited in order to get the closing FA's balance to $0"}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "forwarding_settings": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.financial_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
