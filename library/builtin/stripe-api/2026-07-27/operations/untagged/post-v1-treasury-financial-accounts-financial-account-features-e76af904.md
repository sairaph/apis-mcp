---
title: Update FinancialAccount Features
page_id: operation-post-v1-treasury-financial-accounts-financial-account-features-3f78c0e2
path: operations/untagged
description: <p>Updates the Features associated with a FinancialAccount.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/treasury/financial_accounts/{financial_account}/features
operation_ids:
    - PostTreasuryFinancialAccountsFinancialAccountFeatures
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update FinancialAccount Features

`POST /v1/treasury/financial_accounts/{financial_account}/features`

Operation ID: `PostTreasuryFinancialAccountsFinancialAccountFeatures`

<p>Updates the Features associated with a FinancialAccount.</p>

## Definition

```yaml
{"summary": "Update FinancialAccount Features", "description": "<p>Updates the Features associated with a FinancialAccount.</p>", "operationId": "PostTreasuryFinancialAccountsFinancialAccountFeatures", "parameters": [{"name": "financial_account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"card_issuing": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}, "description": "Encodes the FinancialAccount's ability to be used with the Issuing product, including attaching cards to and drawing funds from the FinancialAccount."}, "deposit_insurance": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}, "description": "Represents whether this FinancialAccount is eligible for deposit insurance. Various factors determine the insurance amount."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "financial_addresses": {"title": "financial_addresses", "type": "object", "properties": {"aba": {"title": "aba_access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}, "description": "Contains Features that add FinancialAddresses to the FinancialAccount."}, "inbound_transfers": {"title": "inbound_transfers", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_inbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}, "description": "Contains settings related to adding funds to a FinancialAccount from another Account with the same owner."}, "intra_stripe_flows": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}, "description": "Represents the ability for the FinancialAccount to send money to, or receive money from other FinancialAccounts (for example, via OutboundPayment)."}, "outbound_payments": {"title": "outbound_payments", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_outbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "us_domestic_wire": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}, "description": "Includes Features related to initiating money movement out of the FinancialAccount to someone else's bucket of money."}, "outbound_transfers": {"title": "outbound_transfers", "type": "object", "properties": {"ach": {"title": "access_with_ach_details_outbound", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}, "us_domestic_wire": {"title": "access", "required": ["requested"], "type": "object", "properties": {"requested": {"type": "boolean"}}}}, "description": "Contains a Feature and settings related to moving money out of the FinancialAccount into another Account with the same owner."}}, "additionalProperties": false}, "encoding": {"card_issuing": {"style": "deepObject", "explode": true}, "deposit_insurance": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "financial_addresses": {"style": "deepObject", "explode": true}, "inbound_transfers": {"style": "deepObject", "explode": true}, "intra_stripe_flows": {"style": "deepObject", "explode": true}, "outbound_payments": {"style": "deepObject", "explode": true}, "outbound_transfers": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.financial_account_features"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
