---
title: payment_links_resource_custom_fields_dropdown
page_id: schema-payment-links-resource-custom-fields-dropdown-6c324708
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_custom_fields_dropdown

```yaml
{"title": "PaymentLinksResourceCustomFieldsDropdown", "required": ["options"], "type": "object", "properties": {"default_value": {"maxLength": 5000, "type": "string", "description": "The value that pre-fills on the payment page.", "nullable": true}, "options": {"type": "array", "description": "The options available for the customer to select. Up to 200 options allowed.", "items": {"$ref": "#/components/schemas/payment_links_resource_custom_fields_dropdown_option"}}}, "description": "", "x-expandableFields": ["options"]}
```
