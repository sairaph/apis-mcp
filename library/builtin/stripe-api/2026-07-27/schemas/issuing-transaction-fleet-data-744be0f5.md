---
title: issuing_transaction_fleet_data
page_id: schema-issuing-transaction-fleet-data-744be0f5
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_fleet_data

```yaml
{"title": "IssuingTransactionFleetData", "type": "object", "properties": {"cardholder_prompt_data": {"description": "Answers to prompts presented to cardholder at point of sale.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_transaction_fleet_cardholder_prompt_data"}]}, "purchase_type": {"maxLength": 5000, "type": "string", "description": "The type of purchase. One of `fuel_purchase`, `non_fuel_purchase`, or `fuel_and_non_fuel_purchase`.", "nullable": true}, "reported_breakdown": {"description": "More information about the total amount. This information is not guaranteed to be accurate as some merchants may provide unreliable data.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_transaction_fleet_reported_breakdown"}]}, "service_type": {"maxLength": 5000, "type": "string", "description": "The type of fuel service. One of `non_fuel_transaction`, `full_service`, or `self_service`.", "nullable": true}}, "description": "", "x-expandableFields": ["cardholder_prompt_data", "reported_breakdown"]}
```
