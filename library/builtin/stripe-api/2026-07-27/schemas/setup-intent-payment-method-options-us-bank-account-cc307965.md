---
title: setup_intent_payment_method_options_us_bank_account
page_id: schema-setup-intent-payment-method-options-us-bank-account-cc307965
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# setup_intent_payment_method_options_us_bank_account

```yaml
{"title": "setup_intent_payment_method_options_us_bank_account", "type": "object", "properties": {"financial_connections": {"$ref": "#/components/schemas/linked_account_options_common"}, "mandate_options": {"$ref": "#/components/schemas/payment_method_options_us_bank_account_mandate_options"}, "verification_method": {"type": "string", "description": "Bank account verification method. The default value is `automatic`.", "enum": ["automatic", "instant", "microdeposits"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["financial_connections", "mandate_options"]}
```
