---
title: payment_pages_checkout_session_custom_fields_text
page_id: schema-payment-pages-checkout-session-custom-fields-text-f2822ad6
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_custom_fields_text

```yaml
{"title": "PaymentPagesCheckoutSessionCustomFieldsText", "type": "object", "properties": {"default_value": {"maxLength": 5000, "type": "string", "description": "The value that pre-fills the field on the payment page.", "nullable": true}, "maximum_length": {"type": "integer", "description": "The maximum character length constraint for the customer's input.", "nullable": true}, "minimum_length": {"type": "integer", "description": "The minimum character length requirement for the customer's input.", "nullable": true}, "value": {"maxLength": 5000, "type": "string", "description": "The value entered by the customer.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
