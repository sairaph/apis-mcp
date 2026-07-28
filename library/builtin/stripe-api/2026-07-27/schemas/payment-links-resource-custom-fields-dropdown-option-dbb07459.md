---
title: payment_links_resource_custom_fields_dropdown_option
page_id: schema-payment-links-resource-custom-fields-dropdown-option-dbb07459
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_custom_fields_dropdown_option

```yaml
{"title": "PaymentLinksResourceCustomFieldsDropdownOption", "required": ["label", "value"], "type": "object", "properties": {"label": {"maxLength": 5000, "type": "string", "description": "The label for the option, displayed to the customer. Up to 100 characters."}, "value": {"maxLength": 5000, "type": "string", "description": "The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters."}}, "description": "", "x-expandableFields": []}
```
