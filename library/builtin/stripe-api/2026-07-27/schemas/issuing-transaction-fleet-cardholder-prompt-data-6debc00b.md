---
title: issuing_transaction_fleet_cardholder_prompt_data
page_id: schema-issuing-transaction-fleet-cardholder-prompt-data-6debc00b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_fleet_cardholder_prompt_data

```yaml
{"title": "IssuingTransactionFleetCardholderPromptData", "type": "object", "properties": {"driver_id": {"maxLength": 5000, "type": "string", "description": "Driver ID.", "nullable": true}, "odometer": {"type": "integer", "description": "Odometer reading.", "nullable": true}, "unspecified_id": {"maxLength": 5000, "type": "string", "description": "An alphanumeric ID. This field is used when a vehicle ID, driver ID, or generic ID is entered by the cardholder, but the merchant or card network did not specify the prompt type.", "nullable": true}, "user_id": {"maxLength": 5000, "type": "string", "description": "User ID.", "nullable": true}, "vehicle_number": {"maxLength": 5000, "type": "string", "description": "Vehicle number.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
