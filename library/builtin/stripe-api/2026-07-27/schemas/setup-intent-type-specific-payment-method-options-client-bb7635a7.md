---
title: setup_intent_type_specific_payment_method_options_client
page_id: schema-setup-intent-type-specific-payment-method-options-client-bb7635a7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# setup_intent_type_specific_payment_method_options_client

```yaml
{"title": "SetupIntentTypeSpecificPaymentMethodOptionsClient", "type": "object", "properties": {"mandate_options": {"$ref": "#/components/schemas/setup_intent_payment_method_options_mandate_options_payto"}, "verification_method": {"type": "string", "description": "Bank account verification method. The default value is `automatic`.", "enum": ["automatic", "instant", "microdeposits"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["mandate_options"]}
```
