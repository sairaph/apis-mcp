---
title: issuing_transaction_purchase_details
page_id: schema-issuing-transaction-purchase-details-27c1dae9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_purchase_details

```yaml
{"title": "IssuingTransactionPurchaseDetails", "type": "object", "properties": {"fleet": {"description": "Fleet-specific information for transactions using Fleet cards.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_transaction_fleet_data"}]}, "flight": {"description": "Information about the flight that was purchased with this transaction.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_transaction_flight_data"}]}, "fuel": {"description": "Information about fuel that was purchased with this transaction.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_transaction_fuel_data"}]}, "lodging": {"description": "Information about lodging that was purchased with this transaction.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_transaction_lodging_data"}]}, "receipt": {"type": "array", "description": "The line items in the purchase.", "nullable": true, "items": {"$ref": "#/components/schemas/issuing_transaction_receipt_data"}}, "reference": {"maxLength": 5000, "type": "string", "description": "A merchant-specific order number.", "nullable": true}}, "description": "", "x-expandableFields": ["fleet", "flight", "fuel", "lodging", "receipt"]}
```
