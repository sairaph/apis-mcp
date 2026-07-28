---
title: invoice_payment_method_options_us_bank_account_linked_account_options
page_id: schema-invoice-payment-method-options-us-bank-account-linked-account-options-75245deb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment_method_options_us_bank_account_linked_account_options

```yaml
{"title": "invoice_payment_method_options_us_bank_account_linked_account_options", "type": "object", "properties": {"filters": {"$ref": "#/components/schemas/invoice_payment_method_options_us_bank_account_linked_account_options_filters"}, "permissions": {"type": "array", "description": "The list of permissions to request. The `payment_method` permission must be included.", "items": {"type": "string", "enum": ["balances", "ownership", "payment_method", "transactions"], "x-stripeBypassValidation": true}}, "prefetch": {"type": "array", "description": "Data features requested to be retrieved upon account creation.", "nullable": true, "items": {"type": "string", "enum": ["balances", "ownership", "transactions"], "x-stripeBypassValidation": true}}}, "description": "", "x-expandableFields": ["filters"]}
```
