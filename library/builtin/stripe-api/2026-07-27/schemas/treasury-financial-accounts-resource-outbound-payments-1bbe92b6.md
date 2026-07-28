---
title: treasury_financial_accounts_resource_outbound_payments
page_id: schema-treasury-financial-accounts-resource-outbound-payments-1bbe92b6
path: schemas
description: Settings related to Outbound Payments features on a Financial Account
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_outbound_payments

Settings related to Outbound Payments features on a Financial Account

```yaml
{"title": "TreasuryFinancialAccountsResourceOutboundPayments", "type": "object", "properties": {"ach": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_outbound_ach_toggle_settings"}, "us_domestic_wire": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggle_settings"}}, "description": "Settings related to Outbound Payments features on a Financial Account", "x-expandableFields": ["ach", "us_domestic_wire"]}
```
