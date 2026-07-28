---
title: issuing_authorization_fleet_data
page_id: schema-issuing-authorization-fleet-data-93137657
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_fleet_data

```yaml
{"title": "IssuingAuthorizationFleetData", "type": "object", "properties": {"cardholder_prompt_data": {"description": "Answers to prompts presented to the cardholder at the point of sale. Prompted fields vary depending on the configuration of your physical fleet cards. Typical points of sale support only numeric entry.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_fleet_cardholder_prompt_data"}]}, "purchase_type": {"type": "string", "description": "The type of purchase.", "nullable": true, "enum": ["fuel_and_non_fuel_purchase", "fuel_purchase", "non_fuel_purchase"]}, "reported_breakdown": {"description": "More information about the total amount. Typically this information is received from the merchant after the authorization has been approved and the fuel dispensed. This information is not guaranteed to be accurate as some merchants may provide unreliable data.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_fleet_reported_breakdown"}]}, "service_type": {"type": "string", "description": "The type of fuel service.", "nullable": true, "enum": ["full_service", "non_fuel_transaction", "self_service"]}}, "description": "", "x-expandableFields": ["cardholder_prompt_data", "reported_breakdown"]}
```
