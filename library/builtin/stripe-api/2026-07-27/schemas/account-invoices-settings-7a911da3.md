---
title: account_invoices_settings
page_id: schema-account-invoices-settings-7a911da3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_invoices_settings

```yaml
{"title": "AccountInvoicesSettings", "type": "object", "properties": {"default_account_tax_ids": {"type": "array", "description": "The list of default Account Tax IDs to automatically include on invoices. Account Tax IDs get added when an invoice is finalized.", "nullable": true, "items": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/tax_id"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/tax_id"}]}}}, "hosted_payment_method_save": {"type": "string", "description": "Whether to save the payment method after a payment is completed for a one-time invoice or a subscription invoice when the customer already has a default payment method on the hosted invoice page.", "nullable": true, "enum": ["always", "never", "offer"]}}, "description": "", "x-expandableFields": ["default_account_tax_ids"]}
```
