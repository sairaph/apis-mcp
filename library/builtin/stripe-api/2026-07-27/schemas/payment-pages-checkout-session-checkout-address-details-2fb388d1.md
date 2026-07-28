---
title: payment_pages_checkout_session_checkout_address_details
page_id: schema-payment-pages-checkout-session-checkout-address-details-2fb388d1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_checkout_address_details

```yaml
{"title": "PaymentPagesCheckoutSessionCheckoutAddressDetails", "required": ["address", "name"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/address"}, "name": {"maxLength": 5000, "type": "string", "description": "Customer name."}}, "description": "", "x-expandableFields": ["address"]}
```
