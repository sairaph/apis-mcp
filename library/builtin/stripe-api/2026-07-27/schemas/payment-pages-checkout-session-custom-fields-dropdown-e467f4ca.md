---
title: payment_pages_checkout_session_custom_fields_dropdown
page_id: schema-payment-pages-checkout-session-custom-fields-dropdown-e467f4ca
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_custom_fields_dropdown

```yaml
{"title": "PaymentPagesCheckoutSessionCustomFieldsDropdown", "required": ["options"], "type": "object", "properties": {"default_value": {"maxLength": 5000, "type": "string", "description": "The value that pre-fills on the payment page.", "nullable": true}, "options": {"type": "array", "description": "The options available for the customer to select. Up to 200 options allowed.", "items": {"$ref": "#/components/schemas/payment_pages_checkout_session_custom_fields_option"}}, "value": {"maxLength": 5000, "type": "string", "description": "The option selected by the customer. This will be the `value` for the option.", "nullable": true}}, "description": "", "x-expandableFields": ["options"]}
```
