---
title: payment_pages_checkout_session_shipping_option
page_id: schema-payment-pages-checkout-session-shipping-option-efbc72c2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_shipping_option

```yaml
{"title": "PaymentPagesCheckoutSessionShippingOption", "required": ["shipping_amount", "shipping_rate"], "type": "object", "properties": {"shipping_amount": {"type": "integer", "description": "A non-negative integer in cents representing how much to charge."}, "shipping_rate": {"description": "The shipping rate.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/shipping_rate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/shipping_rate"}]}}}, "description": "", "x-expandableFields": ["shipping_rate"]}
```
