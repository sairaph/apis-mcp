---
title: portal_subscription_update_product_adjustable_quantity
page_id: schema-portal-subscription-update-product-adjustable-quantity-0d6ff986
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_subscription_update_product_adjustable_quantity

```yaml
{"title": "PortalSubscriptionUpdateProductAdjustableQuantity", "required": ["enabled", "minimum"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "If true, the quantity can be adjusted to any non-negative integer."}, "maximum": {"type": "integer", "description": "The maximum quantity that can be set for the product.", "nullable": true}, "minimum": {"type": "integer", "description": "The minimum quantity that can be set for the product."}}, "description": "", "x-expandableFields": []}
```
