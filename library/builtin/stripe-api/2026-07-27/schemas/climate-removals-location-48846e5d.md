---
title: climate_removals_location
page_id: schema-climate-removals-location-48846e5d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# climate_removals_location

```yaml
{"title": "ClimateRemovalsLocation", "required": ["country"], "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string", "description": "The city where the supplier is located.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country where the supplier is located."}, "latitude": {"type": "number", "description": "The geographic latitude where the supplier is located.", "nullable": true}, "longitude": {"type": "number", "description": "The geographic longitude where the supplier is located.", "nullable": true}, "region": {"maxLength": 5000, "type": "string", "description": "The state/county/province/region where the supplier is located.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
