---
title: climate.supplier
page_id: schema-climate-supplier-bb22087d
path: schemas
description: A supplier of carbon removal.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# climate.supplier

A supplier of carbon removal.

```yaml
{"title": "ClimateRemovalsSuppliers", "required": ["id", "info_url", "livemode", "locations", "name", "object", "removal_pathway"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "info_url": {"maxLength": 5000, "type": "string", "description": "Link to a webpage to learn more about the supplier."}, "livemode": {"type": "boolean", "description": "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode."}, "locations": {"type": "array", "description": "The locations in which this supplier operates.", "items": {"$ref": "#/components/schemas/climate_removals_location"}}, "name": {"maxLength": 5000, "type": "string", "description": "Name of this carbon removal supplier."}, "object": {"type": "string", "description": "String representing the object’s type. Objects of the same type share the same value.", "enum": ["climate.supplier"]}, "removal_pathway": {"type": "string", "description": "The scientific pathway used for carbon removal.", "enum": ["biomass_carbon_removal_and_storage", "direct_air_capture", "enhanced_weathering", "marine_carbon_removal"], "x-stripeBypassValidation": true}}, "description": "A supplier of carbon removal.", "x-expandableFields": ["locations"], "x-resourceId": "climate.supplier"}
```
