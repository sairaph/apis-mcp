---
title: treasury_inbound_transfers_resource_inbound_transfer_resource_status_transitions
page_id: schema-treasury-inbound-transfers-resource-inbound-transfer-resource-status-tra-a9fc7898
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_inbound_transfers_resource_inbound_transfer_resource_status_transitions

```yaml
{"title": "TreasuryInboundTransfersResourceInboundTransferResourceStatusTransitions", "type": "object", "properties": {"canceled_at": {"type": "integer", "description": "Timestamp describing when an InboundTransfer changed status to `canceled`.", "format": "unix-time", "nullable": true}, "failed_at": {"type": "integer", "description": "Timestamp describing when an InboundTransfer changed status to `failed`.", "format": "unix-time", "nullable": true}, "succeeded_at": {"type": "integer", "description": "Timestamp describing when an InboundTransfer changed status to `succeeded`.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
