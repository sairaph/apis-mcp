---
title: treasury_financial_accounts_resource_balance
page_id: schema-treasury-financial-accounts-resource-balance-6888f2d0
path: schemas
description: Balance information for the FinancialAccount
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_balance

Balance information for the FinancialAccount

```yaml
{"title": "TreasuryFinancialAccountsResourceBalance", "required": ["cash", "inbound_pending", "outbound_pending"], "type": "object", "properties": {"cash": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "Funds the user can spend right now."}, "inbound_pending": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "Funds not spendable yet, but will become available at a later time."}, "outbound_pending": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "Funds in the account, but not spendable because they are being held for pending outbound flows."}}, "description": "Balance information for the FinancialAccount", "x-expandableFields": []}
```
