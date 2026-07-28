---
title: payment_pages_checkout_session_after_expiration_recovery
page_id: schema-payment-pages-checkout-session-after-expiration-recovery-27781e6a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_after_expiration_recovery

```yaml
{"title": "PaymentPagesCheckoutSessionAfterExpirationRecovery", "required": ["allow_promotion_codes", "enabled"], "type": "object", "properties": {"allow_promotion_codes": {"type": "boolean", "description": "Enables user redeemable promotion codes on the recovered Checkout Sessions. Defaults to `false`"}, "enabled": {"type": "boolean", "description": "If `true`, a recovery url will be generated to recover this Checkout Session if it\nexpires before a transaction is completed. It will be attached to the\nCheckout Session object upon expiration."}, "expires_at": {"type": "integer", "description": "The timestamp at which the recovery URL will expire.", "format": "unix-time", "nullable": true}, "url": {"maxLength": 5000, "type": "string", "description": "URL that creates a new Checkout Session when clicked that is a copy of this expired Checkout Session", "nullable": true}}, "description": "", "x-expandableFields": []}
```
