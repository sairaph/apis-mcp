---
title: treasury_outbound_transfers_resource_returned_details
page_id: schema-treasury-outbound-transfers-resource-returned-details-b423169d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_transfers_resource_returned_details

```yaml
{"title": "TreasuryOutboundTransfersResourceReturnedDetails", "required": ["code", "transaction"], "type": "object", "properties": {"code": {"type": "string", "description": "Reason for the return.", "enum": ["account_closed", "account_frozen", "bank_account_restricted", "bank_ownership_changed", "declined", "incorrect_account_holder_name", "invalid_account_number", "invalid_currency", "no_account", "other"]}, "transaction": {"description": "The Transaction associated with this object.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/treasury.transaction"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/treasury.transaction"}]}}}, "description": "", "x-expandableFields": ["transaction"]}
```
