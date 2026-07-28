---
title: issuing_authorization_fleet_reported_breakdown
page_id: schema-issuing-authorization-fleet-reported-breakdown-fd8c8af1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_fleet_reported_breakdown

```yaml
{"title": "IssuingAuthorizationFleetReportedBreakdown", "type": "object", "properties": {"fuel": {"description": "Breakdown of fuel portion of the purchase.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_fleet_fuel_price_data"}]}, "non_fuel": {"description": "Breakdown of non-fuel portion of the purchase.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_fleet_non_fuel_price_data"}]}, "tax": {"description": "Information about tax included in this transaction.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_fleet_tax_data"}]}}, "description": "", "x-expandableFields": ["fuel", "non_fuel", "tax"]}
```
