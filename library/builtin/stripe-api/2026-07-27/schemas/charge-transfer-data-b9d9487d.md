---
title: charge_transfer_data
page_id: schema-charge-transfer-data-b9d9487d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# charge_transfer_data

```yaml
{"title": "ChargeTransferData", "required": ["destination"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.", "nullable": true}, "destination": {"description": "ID of an existing, connected Stripe account to transfer funds to if `transfer_data` was specified in the charge request.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}}, "description": "", "x-expandableFields": ["destination"]}
```
