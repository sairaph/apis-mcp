---
title: payment_pages_checkout_session_custom_fields_numeric
page_id: schema-payment-pages-checkout-session-custom-fields-numeric-ab81c39f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_custom_fields_numeric

```yaml
{"title": "PaymentPagesCheckoutSessionCustomFieldsNumeric", "type": "object", "properties": {"default_value": {"maxLength": 5000, "type": "string", "description": "The value that pre-fills the field on the payment page.", "nullable": true}, "maximum_length": {"type": "integer", "description": "The maximum character length constraint for the customer's input.", "nullable": true}, "minimum_length": {"type": "integer", "description": "The minimum character length requirement for the customer's input.", "nullable": true}, "value": {"maxLength": 5000, "type": "string", "description": "The value entered by the customer, containing only digits.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
