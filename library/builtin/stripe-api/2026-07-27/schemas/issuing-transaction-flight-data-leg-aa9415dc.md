---
title: issuing_transaction_flight_data_leg
page_id: schema-issuing-transaction-flight-data-leg-aa9415dc
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_flight_data_leg

```yaml
{"title": "IssuingTransactionFlightDataLeg", "type": "object", "properties": {"arrival_airport_code": {"maxLength": 5000, "type": "string", "description": "The three-letter IATA airport code of the flight's destination.", "nullable": true}, "carrier": {"maxLength": 5000, "type": "string", "description": "The airline carrier code.", "nullable": true}, "departure_airport_code": {"maxLength": 5000, "type": "string", "description": "The three-letter IATA airport code that the flight departed from.", "nullable": true}, "flight_number": {"maxLength": 5000, "type": "string", "description": "The flight number.", "nullable": true}, "service_class": {"maxLength": 5000, "type": "string", "description": "The flight's service class.", "nullable": true}, "stopover_allowed": {"type": "boolean", "description": "Whether a stopover is allowed on this flight.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
