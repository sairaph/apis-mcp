---
title: billing_credit_grants_resource_balance_credit
page_id: schema-billing-credit-grants-resource-balance-credit-6f0863b7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_credit_grants_resource_balance_credit

```yaml
{"title": "BillingCreditGrantsResourceBalanceCredit", "required": ["amount", "type"], "type": "object", "properties": {"amount": {"$ref": "#/components/schemas/billing_credit_grants_resource_amount"}, "credits_application_invoice_voided": {"description": "Details of the invoice to which the reinstated credits were originally applied. Only present if `type` is `credits_application_invoice_voided`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_credit_grants_resource_balance_credits_application_invoice_voided"}]}, "type": {"type": "string", "description": "The type of credit transaction.", "enum": ["credits_application_invoice_voided", "credits_granted"]}}, "description": "", "x-expandableFields": ["amount", "credits_application_invoice_voided"]}
```
