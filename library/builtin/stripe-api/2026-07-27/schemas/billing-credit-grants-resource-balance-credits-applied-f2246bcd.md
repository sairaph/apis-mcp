---
title: billing_credit_grants_resource_balance_credits_applied
page_id: schema-billing-credit-grants-resource-balance-credits-applied-f2246bcd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_credit_grants_resource_balance_credits_applied

```yaml
{"title": "BillingCreditGrantsResourceBalanceCreditsApplied", "required": ["invoice", "invoice_line_item"], "type": "object", "properties": {"invoice": {"description": "The invoice to which the billing credits were applied.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/invoice"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/invoice"}]}}, "invoice_line_item": {"maxLength": 5000, "type": "string", "description": "The invoice line item to which the billing credits were applied."}}, "description": "", "x-expandableFields": ["invoice"]}
```
