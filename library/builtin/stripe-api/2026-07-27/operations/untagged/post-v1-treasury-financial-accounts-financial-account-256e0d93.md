---
title: Update a FinancialAccount
page_id: operation-post-v1-treasury-financial-accounts-financial-account-14f97065
path: operations/untagged
description: <p>Updates the details of a FinancialAccount.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/financial_accounts/{financial_account}
operation_ids:
    - PostTreasuryFinancialAccountsFinancialAccount
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a FinancialAccount

`POST /v1/treasury/financial_accounts/{financial_account}`

Operation ID: `PostTreasuryFinancialAccountsFinancialAccount`

<p>Updates the details of a FinancialAccount.</p>

## Definition

```yaml
{"summary": "Update a FinancialAccount", "description": "<p>Updates the details of a FinancialAccount.</p>", "operationId": "PostTreasuryFinancialAccountsFinancialAccount", "parameters": [{"name": "financial_account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "features": {"title": "feature_access", "type": "object", "properties": {"card_issuing": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "deposit_insurance": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "financial_addresses": {"title": "financial_addresses", "type": "object", "properties": {"aba": {"title": "aba_access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}, "inbound_transfers": {"title": "inbound_transfers", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_inbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}, "intra_stripe_flows": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "outbound_payments": {"title": "outbound_payments", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_outbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "us_domestic_wire": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}, "outbound_transfers": {"title": "outbound_transfers", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_outbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "us_domestic_wire": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}}, "description": "Encodes whether a FinancialAccount has access to a particular feature, with a status enum and associated `status_details`. Stripe or the platform may control features via the requested field."}, "forwarding_settings": {"title": "forwarding_settings", "required": ["type"], "type": "object", "properties": {"financial_account": {"type": "string"}, "payment_method": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["financial_account", "payment_method"]}}, "description": "A different bank account where funds can be deposited/debited in order to get the closing FA's balance to $0"}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "nickname": {"description": "The nickname for the FinancialAccount.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "platform_restrictions": {"title": "platform_restrictions", "type": "object", "properties": {"inbound_flows": {"type": "string", "enum": ["restricted", "unrestricted"]}, "outbound_flows": {"type": "string", "enum": ["restricted", "unrestricted"]}}, "description": "The set of functionalities that the platform can restrict on the FinancialAccount."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "features": {"style": "deepObject", "explode": true}, "forwarding_settings": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "nickname": {"style": "deepObject", "explode": true}, "platform_restrictions": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.financial_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
