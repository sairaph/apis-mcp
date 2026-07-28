---
title: issuing_transaction_flight_data
page_id: schema-issuing-transaction-flight-data-ac122d08
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_flight_data

```yaml
{"title": "IssuingTransactionFlightData", "type": "object", "properties": {"departure_at": {"type": "integer", "description": "The time that the flight departed.", "nullable": true}, "passenger_name": {"maxLength": 5000, "type": "string", "description": "The name of the passenger.", "nullable": true}, "refundable": {"type": "boolean", "description": "Whether the ticket is refundable.", "nullable": true}, "segments": {"type": "array", "description": "The legs of the trip.", "nullable": true, "items": {"$ref": "#/components/schemas/issuing_transaction_flight_data_leg"}}, "travel_agency": {"maxLength": 5000, "type": "string", "description": "The travel agency that issued the ticket.", "nullable": true}}, "description": "", "x-expandableFields": ["segments"]}
```
