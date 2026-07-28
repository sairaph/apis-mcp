---
title: payment_pages_checkout_session_consent
page_id: schema-payment-pages-checkout-session-consent-12785356
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_consent

```yaml
{"title": "PaymentPagesCheckoutSessionConsent", "type": "object", "properties": {"promotions": {"type": "string", "description": "If `opt_in`, the customer consents to receiving promotional communications\nfrom the merchant about this Checkout Session.", "nullable": true, "enum": ["opt_in", "opt_out"]}, "terms_of_service": {"type": "string", "description": "If `accepted`, the customer in this Checkout Session has agreed to the merchant's terms of service.", "nullable": true, "enum": ["accepted"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
