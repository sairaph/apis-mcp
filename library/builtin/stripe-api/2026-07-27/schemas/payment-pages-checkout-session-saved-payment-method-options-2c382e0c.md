---
title: payment_pages_checkout_session_saved_payment_method_options
page_id: schema-payment-pages-checkout-session-saved-payment-method-options-2c382e0c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_saved_payment_method_options

```yaml
{"title": "PaymentPagesCheckoutSessionSavedPaymentMethodOptions", "type": "object", "properties": {"allow_redisplay_filters": {"type": "array", "description": "Uses the `allow_redisplay` value of each saved payment method to filter the set presented to a returning customer. By default, only saved payment methods with ’allow_redisplay: ‘always’ are shown in Checkout.", "nullable": true, "items": {"type": "string", "enum": ["always", "limited", "unspecified"]}}, "payment_method_remove": {"type": "string", "description": "Enable customers to choose if they wish to remove their saved payment methods. Disabled by default.", "nullable": true, "enum": ["disabled", "enabled"]}, "payment_method_save": {"type": "string", "description": "Enable customers to choose if they wish to save their payment method for future use. Disabled by default.", "nullable": true, "enum": ["disabled", "enabled"]}}, "description": "", "x-expandableFields": []}
```
