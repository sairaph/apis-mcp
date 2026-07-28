---
title: treasury.financial_account_features
page_id: schema-treasury-financial-account-features-10d48ccc
path: schemas
description: |-
    Encodes whether a FinancialAccount has access to a particular Feature, with a `status` enum and associated `status_details`.
    Stripe or the platform can control Features via the requested field.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury.financial_account_features

Encodes whether a FinancialAccount has access to a particular Feature, with a `status` enum and associated `status_details`.
Stripe or the platform can control Features via the requested field.

```yaml
{"title": "TreasuryFinancialAccountsResourceFinancialAccountFeatures", "required": ["object"], "type": "object", "properties": {"card_issuing": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggle_settings"}, "deposit_insurance": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggle_settings"}, "financial_addresses": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_financial_addresses_features"}, "inbound_transfers": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_inbound_transfers"}, "intra_stripe_flows": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggle_settings"}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["treasury.financial_account_features"]}, "outbound_payments": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_outbound_payments"}, "outbound_transfers": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_outbound_transfers"}}, "description": "Encodes whether a FinancialAccount has access to a particular Feature, with a `status` enum and associated `status_details`.\nStripe or the platform can control Features via the requested field.", "x-expandableFields": ["card_issuing", "deposit_insurance", "financial_addresses", "inbound_transfers", "intra_stripe_flows", "outbound_payments", "outbound_transfers"], "x-resourceId": "treasury.financial_account_features"}
```
