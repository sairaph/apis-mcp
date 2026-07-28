---
title: payment_pages_checkout_session_tax_id_collection
page_id: schema-payment-pages-checkout-session-tax-id-collection-0279f5d3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_tax_id_collection

```yaml
{"title": "PaymentPagesCheckoutSessionTaxIDCollection", "required": ["enabled", "required"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Indicates whether tax ID collection is enabled for the session"}, "required": {"type": "string", "description": "Indicates whether a tax ID is required on the payment page", "enum": ["if_supported", "never"]}}, "description": "", "x-expandableFields": []}
```
