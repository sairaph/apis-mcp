---
title: shipping_rate_delivery_estimate_bound
page_id: schema-shipping-rate-delivery-estimate-bound-75b1e6c3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# shipping_rate_delivery_estimate_bound

```yaml
{"title": "ShippingRateDeliveryEstimateBound", "required": ["unit", "value"], "type": "object", "properties": {"unit": {"type": "string", "description": "A unit of time.", "enum": ["business_day", "day", "hour", "month", "week"]}, "value": {"type": "integer", "description": "Must be greater than 0."}}, "description": "", "x-expandableFields": []}
```
