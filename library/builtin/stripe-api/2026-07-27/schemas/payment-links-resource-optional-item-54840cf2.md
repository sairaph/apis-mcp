---
title: payment_links_resource_optional_item
page_id: schema-payment-links-resource-optional-item-54840cf2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_optional_item

```yaml
{"title": "PaymentLinksResourceOptionalItem", "required": ["price", "quantity"], "type": "object", "properties": {"adjustable_quantity": {"nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_optional_item_adjustable_quantity"}]}, "price": {"maxLength": 5000, "type": "string"}, "quantity": {"type": "integer"}}, "description": "", "x-expandableFields": ["adjustable_quantity"]}
```
