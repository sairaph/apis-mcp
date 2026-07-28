---
title: payment_links_resource_custom_fields_label
page_id: schema-payment-links-resource-custom-fields-label-fcb77285
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_custom_fields_label

```yaml
{"title": "PaymentLinksResourceCustomFieldsLabel", "required": ["type"], "type": "object", "properties": {"custom": {"maxLength": 5000, "type": "string", "description": "Custom text for the label, displayed to the customer. Up to 50 characters.", "nullable": true}, "type": {"type": "string", "description": "The type of the label.", "enum": ["custom"]}}, "description": "", "x-expandableFields": []}
```
