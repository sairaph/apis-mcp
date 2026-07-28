---
title: payment_links_resource_custom_fields
page_id: schema-payment-links-resource-custom-fields-a38160be
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_custom_fields

```yaml
{"title": "PaymentLinksResourceCustomFields", "required": ["key", "label", "optional", "type"], "type": "object", "properties": {"dropdown": {"$ref": "#/components/schemas/payment_links_resource_custom_fields_dropdown"}, "key": {"maxLength": 5000, "type": "string", "description": "String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters."}, "label": {"$ref": "#/components/schemas/payment_links_resource_custom_fields_label"}, "numeric": {"$ref": "#/components/schemas/payment_links_resource_custom_fields_numeric"}, "optional": {"type": "boolean", "description": "Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`."}, "text": {"$ref": "#/components/schemas/payment_links_resource_custom_fields_text"}, "type": {"type": "string", "description": "The type of the field.", "enum": ["dropdown", "numeric", "text"]}}, "description": "", "x-expandableFields": ["dropdown", "label", "numeric", "text"]}
```
