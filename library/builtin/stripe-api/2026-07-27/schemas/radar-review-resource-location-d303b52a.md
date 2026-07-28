---
title: radar_review_resource_location
page_id: schema-radar-review-resource-location-d303b52a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# radar_review_resource_location

```yaml
{"title": "RadarReviewResourceLocation", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string", "description": "The city where the payment originated.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country where the payment originated.", "nullable": true}, "latitude": {"type": "number", "description": "The geographic latitude where the payment originated.", "nullable": true}, "longitude": {"type": "number", "description": "The geographic longitude where the payment originated.", "nullable": true}, "region": {"maxLength": 5000, "type": "string", "description": "The state/county/province/region where the payment originated.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
