---
title: payment_pages_checkout_session_invoice_creation
page_id: schema-payment-pages-checkout-session-invoice-creation-9a7f8f1f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_invoice_creation

```yaml
{"title": "PaymentPagesCheckoutSessionInvoiceCreation", "required": ["enabled", "invoice_data"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Indicates whether invoice creation is enabled for the Checkout Session."}, "invoice_data": {"$ref": "#/components/schemas/payment_pages_checkout_session_invoice_settings"}}, "description": "", "x-expandableFields": ["invoice_data"]}
```
