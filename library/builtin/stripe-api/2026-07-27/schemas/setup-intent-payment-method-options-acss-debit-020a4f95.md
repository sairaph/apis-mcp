---
title: setup_intent_payment_method_options_acss_debit
page_id: schema-setup-intent-payment-method-options-acss-debit-020a4f95
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# setup_intent_payment_method_options_acss_debit

```yaml
{"title": "setup_intent_payment_method_options_acss_debit", "type": "object", "properties": {"currency": {"type": "string", "description": "Currency supported by the bank account", "nullable": true, "enum": ["cad", "usd"]}, "mandate_options": {"$ref": "#/components/schemas/setup_intent_payment_method_options_mandate_options_acss_debit"}, "verification_method": {"type": "string", "description": "Bank account verification method. The default value is `automatic`.", "enum": ["automatic", "instant", "microdeposits"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["mandate_options"]}
```
