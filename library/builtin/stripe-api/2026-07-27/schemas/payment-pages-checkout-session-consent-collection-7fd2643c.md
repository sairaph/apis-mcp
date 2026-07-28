---
title: payment_pages_checkout_session_consent_collection
page_id: schema-payment-pages-checkout-session-consent-collection-7fd2643c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_consent_collection

```yaml
{"title": "PaymentPagesCheckoutSessionConsentCollection", "type": "object", "properties": {"payment_method_reuse_agreement": {"description": "If set to `hidden`, it will hide legal text related to the reuse of a payment method.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_pages_checkout_session_payment_method_reuse_agreement"}]}, "promotions": {"type": "string", "description": "If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout\nSession will determine whether to display an option to opt into promotional communication\nfrom the merchant depending on the customer's locale. Only available to US merchants and US customers.", "nullable": true, "enum": ["auto", "none"]}, "terms_of_service": {"type": "string", "description": "If set to `required`, it requires customers to accept the terms of service before being able to pay.", "nullable": true, "enum": ["none", "required"]}}, "description": "", "x-expandableFields": ["payment_method_reuse_agreement"]}
```
