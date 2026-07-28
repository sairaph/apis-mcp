---
title: payment_pages_checkout_session_collected_information
page_id: schema-payment-pages-checkout-session-collected-information-066e9288
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_collected_information

```yaml
{"title": "PaymentPagesCheckoutSessionCollectedInformation", "type": "object", "properties": {"business_name": {"maxLength": 5000, "type": "string", "description": "Customer’s business name for this Checkout Session", "nullable": true}, "individual_name": {"maxLength": 5000, "type": "string", "description": "Customer’s individual name for this Checkout Session", "nullable": true}, "shipping_details": {"description": "Shipping information for this Checkout Session.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_pages_checkout_session_checkout_address_details"}]}}, "description": "", "x-expandableFields": ["shipping_details"]}
```
