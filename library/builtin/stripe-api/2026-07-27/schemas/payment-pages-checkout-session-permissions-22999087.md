---
title: payment_pages_checkout_session_permissions
page_id: schema-payment-pages-checkout-session-permissions-22999087
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_permissions

```yaml
{"title": "PaymentPagesCheckoutSessionPermissions", "type": "object", "properties": {"update_shipping_details": {"type": "string", "description": "Determines which entity is allowed to update the shipping details.\n\nDefault is `client_only`. Stripe Checkout client will automatically update the shipping details. If set to `server_only`, only your server is allowed to update the shipping details.\n\nWhen set to `server_only`, you must add the onShippingDetailsChange event handler when initializing the Stripe Checkout client and manually update the shipping details from your server using the Stripe API.", "nullable": true, "enum": ["client_only", "server_only"]}}, "description": "", "x-expandableFields": []}
```
