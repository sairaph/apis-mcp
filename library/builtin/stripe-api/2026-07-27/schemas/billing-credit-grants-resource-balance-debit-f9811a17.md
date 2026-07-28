---
title: billing_credit_grants_resource_balance_debit
page_id: schema-billing-credit-grants-resource-balance-debit-f9811a17
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_credit_grants_resource_balance_debit

```yaml
{"title": "BillingCreditGrantsResourceBalanceDebit", "required": ["amount", "type"], "type": "object", "properties": {"amount": {"$ref": "#/components/schemas/billing_credit_grants_resource_amount"}, "credits_applied": {"description": "Details of how the billing credits were applied to an invoice. Only present if `type` is `credits_applied`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_credit_grants_resource_balance_credits_applied"}]}, "type": {"type": "string", "description": "The type of debit transaction.", "enum": ["credits_applied", "credits_expired", "credits_voided"]}}, "description": "", "x-expandableFields": ["amount", "credits_applied"]}
```
