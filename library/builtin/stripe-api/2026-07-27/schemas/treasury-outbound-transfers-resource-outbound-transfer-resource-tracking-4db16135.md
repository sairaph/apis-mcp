---
title: treasury_outbound_transfers_resource_outbound_transfer_resource_tracking_details
page_id: schema-treasury-outbound-transfers-resource-outbound-transfer-resource-tracking-4db16135
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_transfers_resource_outbound_transfer_resource_tracking_details

```yaml
{"title": "TreasuryOutboundTransfersResourceOutboundTransferResourceTrackingDetails", "required": ["type"], "type": "object", "properties": {"ach": {"$ref": "#/components/schemas/treasury_outbound_transfers_resource_ach_tracking_details"}, "type": {"type": "string", "description": "The US bank account network used to send funds.", "enum": ["ach", "us_domestic_wire"]}, "us_domestic_wire": {"$ref": "#/components/schemas/treasury_outbound_transfers_resource_us_domestic_wire_tracking_details"}}, "description": "", "x-expandableFields": ["ach", "us_domestic_wire"]}
```
