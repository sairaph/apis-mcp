---
title: treasury_outbound_transfers_resource_us_domestic_wire_tracking_details
page_id: schema-treasury-outbound-transfers-resource-us-domestic-wire-tracking-details-3f6bcea1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_transfers_resource_us_domestic_wire_tracking_details

```yaml
{"title": "TreasuryOutboundTransfersResourceUSDomesticWireTrackingDetails", "type": "object", "properties": {"chips": {"maxLength": 5000, "type": "string", "description": "CHIPS System Sequence Number (SSN) of the OutboundTransfer for transfers sent over the `us_domestic_wire` network.", "nullable": true}, "imad": {"maxLength": 5000, "type": "string", "description": "IMAD of the OutboundTransfer for transfers sent over the `us_domestic_wire` network.", "nullable": true}, "omad": {"maxLength": 5000, "type": "string", "description": "OMAD of the OutboundTransfer for transfers sent over the `us_domestic_wire` network.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
