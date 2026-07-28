---
title: shipping_rate_delivery_estimate
page_id: schema-shipping-rate-delivery-estimate-4f46f17b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# shipping_rate_delivery_estimate

```yaml
{"title": "ShippingRateDeliveryEstimate", "type": "object", "properties": {"maximum": {"description": "The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/shipping_rate_delivery_estimate_bound"}]}, "minimum": {"description": "The lower bound of the estimated range. If empty, represents no lower bound.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/shipping_rate_delivery_estimate_bound"}]}}, "description": "", "x-expandableFields": ["maximum", "minimum"]}
```
