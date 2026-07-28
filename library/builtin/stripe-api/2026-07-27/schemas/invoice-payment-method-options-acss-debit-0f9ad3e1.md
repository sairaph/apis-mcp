---
title: invoice_payment_method_options_acss_debit
page_id: schema-invoice-payment-method-options-acss-debit-0f9ad3e1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment_method_options_acss_debit

```yaml
{"title": "invoice_payment_method_options_acss_debit", "type": "object", "properties": {"mandate_options": {"$ref": "#/components/schemas/invoice_payment_method_options_acss_debit_mandate_options"}, "verification_method": {"type": "string", "description": "Bank account verification method. The default value is `automatic`.", "enum": ["automatic", "instant", "microdeposits"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["mandate_options"]}
```
