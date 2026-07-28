---
title: treasury_financial_accounts_resource_financial_address
page_id: schema-treasury-financial-accounts-resource-financial-address-8b04ebb3
path: schemas
description: FinancialAddresses contain identifying information that resolves to a FinancialAccount.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_financial_address

FinancialAddresses contain identifying information that resolves to a FinancialAccount.

```yaml
{"title": "TreasuryFinancialAccountsResourceFinancialAddress", "required": ["type"], "type": "object", "properties": {"aba": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_aba_record"}, "supported_networks": {"type": "array", "description": "The list of networks that the address supports", "items": {"type": "string", "enum": ["ach", "us_domestic_wire"], "x-stripeBypassValidation": true}}, "type": {"type": "string", "description": "The type of financial address", "enum": ["aba"], "x-stripeBypassValidation": true}}, "description": "FinancialAddresses contain identifying information that resolves to a FinancialAccount.", "x-expandableFields": ["aba"]}
```
