---
title: climate_removals_order_deliveries
page_id: schema-climate-removals-order-deliveries-e2f8e3f3
path: schemas
description: The delivery of a specified quantity of carbon for an order.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# climate_removals_order_deliveries

The delivery of a specified quantity of carbon for an order.

```yaml
{"title": "ClimateRemovalsOrderDeliveries", "required": ["delivered_at", "metric_tons", "supplier"], "type": "object", "properties": {"delivered_at": {"type": "integer", "description": "Time at which the delivery occurred. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "location": {"description": "Specific location of this delivery.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/climate_removals_location"}]}, "metric_tons": {"maxLength": 5000, "type": "string", "description": "Quantity of carbon removal supplied by this delivery."}, "registry_url": {"maxLength": 5000, "type": "string", "description": "Once retired, a URL to the registry entry for the tons from this delivery.", "nullable": true}, "supplier": {"$ref": "#/components/schemas/climate.supplier"}}, "description": "The delivery of a specified quantity of carbon for an order.", "x-expandableFields": ["location", "supplier"]}
```
