---
title: portal_subscription_update_product
page_id: schema-portal-subscription-update-product-4c5c2c52
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_subscription_update_product

```yaml
{"title": "PortalSubscriptionUpdateProduct", "required": ["adjustable_quantity", "prices", "product"], "type": "object", "properties": {"adjustable_quantity": {"$ref": "#/components/schemas/portal_subscription_update_product_adjustable_quantity"}, "prices": {"type": "array", "description": "The list of price IDs which, when subscribed to, a subscription can be updated.", "items": {"maxLength": 5000, "type": "string"}}, "product": {"maxLength": 5000, "type": "string", "description": "The product ID."}}, "description": "", "x-expandableFields": ["adjustable_quantity"]}
```
