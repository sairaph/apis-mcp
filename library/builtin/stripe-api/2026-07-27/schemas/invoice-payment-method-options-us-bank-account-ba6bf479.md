---
title: invoice_payment_method_options_us_bank_account
page_id: schema-invoice-payment-method-options-us-bank-account-ba6bf479
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment_method_options_us_bank_account

```yaml
{"title": "invoice_payment_method_options_us_bank_account", "type": "object", "properties": {"financial_connections": {"$ref": "#/components/schemas/invoice_payment_method_options_us_bank_account_linked_account_options"}, "verification_method": {"type": "string", "description": "Bank account verification method. The default value is `automatic`.", "enum": ["automatic", "instant", "microdeposits"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["financial_connections"]}
```
