---
title: payment_pages_checkout_session_discount
page_id: schema-payment-pages-checkout-session-discount-a4f09374
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_discount

```yaml
{"title": "PaymentPagesCheckoutSessionDiscount", "type": "object", "properties": {"coupon": {"description": "Coupon attached to the Checkout Session.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/coupon"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/coupon"}]}}, "promotion_code": {"description": "Promotion code attached to the Checkout Session.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/promotion_code"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/promotion_code"}]}}}, "description": "", "x-expandableFields": ["coupon", "promotion_code"]}
```
