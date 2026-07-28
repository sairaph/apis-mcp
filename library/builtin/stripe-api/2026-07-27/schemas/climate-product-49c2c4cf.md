---
title: climate.product
page_id: schema-climate-product-49c2c4cf
path: schemas
description: |-
    A Climate product represents a type of carbon removal unit available for reservation.
    You can retrieve it to see the current price and availability.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# climate.product

A Climate product represents a type of carbon removal unit available for reservation.
You can retrieve it to see the current price and availability.

```yaml
{"title": "ClimateRemovalsProducts", "required": ["created", "current_prices_per_metric_ton", "id", "livemode", "metric_tons_available", "name", "object", "suppliers"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "current_prices_per_metric_ton": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/climate_removals_products_price"}, "description": "Current prices for a metric ton of carbon removal in a currency's smallest unit."}, "delivery_year": {"type": "integer", "description": "The year in which the carbon removal is expected to be delivered.", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object. For convenience, Climate product IDs are human-readable strings\nthat start with `climsku_`. See [carbon removal inventory](https://stripe.com/docs/climate/orders/carbon-removal-inventory)\nfor a list of available carbon removal products."}, "livemode": {"type": "boolean", "description": "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode."}, "metric_tons_available": {"type": "string", "description": "The quantity of metric tons available for reservation.", "format": "decimal"}, "name": {"maxLength": 5000, "type": "string", "description": "The Climate product's name."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["climate.product"]}, "suppliers": {"type": "array", "description": "The carbon removal suppliers that fulfill orders for this Climate product.", "items": {"$ref": "#/components/schemas/climate.supplier"}}}, "description": "A Climate product represents a type of carbon removal unit available for reservation.\nYou can retrieve it to see the current price and availability.", "x-expandableFields": ["current_prices_per_metric_ton", "suppliers"], "x-resourceId": "climate.product"}
```
