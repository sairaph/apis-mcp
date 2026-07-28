---
title: Create a FinancialAccount
page_id: operation-post-v1-treasury-financial-accounts-6975326c
path: operations/untagged
description: <p>Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/financial_accounts
operation_ids:
    - PostTreasuryFinancialAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a FinancialAccount

`POST /v1/treasury/financial_accounts`

Operation ID: `PostTreasuryFinancialAccounts`

<p>Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.</p>

## Definition

```yaml
{"summary": "Create a FinancialAccount", "description": "<p>Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.</p>", "operationId": "PostTreasuryFinancialAccounts", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["supported_currencies"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "features": {"title": "feature_access", "type": "object", "properties": {"card_issuing": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "deposit_insurance": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "financial_addresses": {"title": "financial_addresses", "type": "object", "properties": {"aba": {"title": "aba_access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}, "inbound_transfers": {"title": "inbound_transfers", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_inbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}, "intra_stripe_flows": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "outbound_payments": {"title": "outbound_payments", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_outbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "us_domestic_wire": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}, "outbound_transfers": {"title": "outbound_transfers", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_outbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "us_domestic_wire": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}}}, "description": "Encodes whether a FinancialAccount has access to a particular feature. Stripe or the platform can control features via the requested field."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "nickname": {"description": "The nickname for the FinancialAccount.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "platform_restrictions": {"title": "platform_restrictions", "type": "object", "properties": {"inbound_flows": {"type": "string", "enum": ["restricted", "unrestricted"]}, "outbound_flows": {"type": "string", "enum": ["restricted", "unrestricted"]}}, "description": "The set of functionalities that the platform can restrict on the FinancialAccount."}, "supported_currencies": {"type": "array", "description": "The currencies the FinancialAccount can hold a balance in.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "features": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "nickname": {"style": "deepObject", "explode": true}, "platform_restrictions": {"style": "deepObject", "explode": true}, "supported_currencies": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.financial_account"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
