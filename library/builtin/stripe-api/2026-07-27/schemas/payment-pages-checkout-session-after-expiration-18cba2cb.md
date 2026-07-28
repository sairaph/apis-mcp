---
title: payment_pages_checkout_session_after_expiration
page_id: schema-payment-pages-checkout-session-after-expiration-18cba2cb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_after_expiration

```yaml
{"title": "PaymentPagesCheckoutSessionAfterExpiration", "type": "object", "properties": {"recovery": {"description": "When set, configuration used to recover the Checkout Session on expiry.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_pages_checkout_session_after_expiration_recovery"}]}}, "description": "", "x-expandableFields": ["recovery"]}
```
