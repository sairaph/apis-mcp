---
title: treasury_financial_accounts_resource_closed_status_details
page_id: schema-treasury-financial-accounts-resource-closed-status-details-cba09c4d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_closed_status_details

```yaml
{"title": "TreasuryFinancialAccountsResourceClosedStatusDetails", "required": ["reasons"], "type": "object", "properties": {"reasons": {"type": "array", "description": "The array that contains reasons for a FinancialAccount closure.", "items": {"type": "string", "enum": ["account_rejected", "closed_by_platform", "other"]}}}, "description": "", "x-expandableFields": []}
```
