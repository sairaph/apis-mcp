---
title: treasury_outbound_payments_resource_us_domestic_wire_tracking_details
page_id: schema-treasury-outbound-payments-resource-us-domestic-wire-tracking-details-657e557f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_payments_resource_us_domestic_wire_tracking_details

```yaml
{"title": "TreasuryOutboundPaymentsResourceUSDomesticWireTrackingDetails", "type": "object", "properties": {"chips": {"maxLength": 5000, "type": "string", "description": "CHIPS System Sequence Number (SSN) of the OutboundPayment for payments sent over the `us_domestic_wire` network.", "nullable": true}, "imad": {"maxLength": 5000, "type": "string", "description": "IMAD of the OutboundPayment for payments sent over the `us_domestic_wire` network.", "nullable": true}, "omad": {"maxLength": 5000, "type": "string", "description": "OMAD of the OutboundPayment for payments sent over the `us_domestic_wire` network.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
