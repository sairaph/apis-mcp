---
title: treasury_outbound_transfers_resource_status_transitions
page_id: schema-treasury-outbound-transfers-resource-status-transitions-44e96bb9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_transfers_resource_status_transitions

```yaml
{"title": "TreasuryOutboundTransfersResourceStatusTransitions", "type": "object", "properties": {"canceled_at": {"type": "integer", "description": "Timestamp describing when an OutboundTransfer changed status to `canceled`", "format": "unix-time", "nullable": true}, "failed_at": {"type": "integer", "description": "Timestamp describing when an OutboundTransfer changed status to `failed`", "format": "unix-time", "nullable": true}, "posted_at": {"type": "integer", "description": "Timestamp describing when an OutboundTransfer changed status to `posted`", "format": "unix-time", "nullable": true}, "returned_at": {"type": "integer", "description": "Timestamp describing when an OutboundTransfer changed status to `returned`", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
