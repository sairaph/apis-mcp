---
title: treasury_financial_accounts_resource_outbound_transfers
page_id: schema-treasury-financial-accounts-resource-outbound-transfers-6ce6bbce
path: schemas
description: OutboundTransfers contains outbound transfers features for a FinancialAccount.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_outbound_transfers

OutboundTransfers contains outbound transfers features for a FinancialAccount.

```yaml
{"title": "TreasuryFinancialAccountsResourceOutboundTransfers", "type": "object", "properties": {"ach": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_outbound_ach_toggle_settings"}, "us_domestic_wire": {"$ref": "#/components/schemas/treasury_financial_accounts_resource_toggle_settings"}}, "description": "OutboundTransfers contains outbound transfers features for a FinancialAccount.", "x-expandableFields": ["ach", "us_domestic_wire"]}
```
