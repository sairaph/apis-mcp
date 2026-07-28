---
title: treasury_transactions_resource_balance_impact
page_id: schema-treasury-transactions-resource-balance-impact-1ea17054
path: schemas
description: Change to a FinancialAccount's balance
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_transactions_resource_balance_impact

Change to a FinancialAccount's balance

```yaml
{"title": "TreasuryTransactionsResourceBalanceImpact", "required": ["cash", "inbound_pending", "outbound_pending"], "type": "object", "properties": {"cash": {"type": "integer", "description": "The change made to funds the user can spend right now."}, "inbound_pending": {"type": "integer", "description": "The change made to funds that are not spendable yet, but will become available at a later time."}, "outbound_pending": {"type": "integer", "description": "The change made to funds in the account, but not spendable because they are being held for pending outbound flows."}}, "description": "Change to a FinancialAccount's balance", "x-expandableFields": []}
```
