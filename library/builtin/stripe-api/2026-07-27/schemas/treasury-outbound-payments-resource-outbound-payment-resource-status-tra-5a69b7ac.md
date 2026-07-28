---
title: treasury_outbound_payments_resource_outbound_payment_resource_status_transitions
page_id: schema-treasury-outbound-payments-resource-outbound-payment-resource-status-tra-5a69b7ac
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_outbound_payments_resource_outbound_payment_resource_status_transitions

```yaml
{"title": "TreasuryOutboundPaymentsResourceOutboundPaymentResourceStatusTransitions", "type": "object", "properties": {"canceled_at": {"type": "integer", "description": "Timestamp describing when an OutboundPayment changed status to `canceled`.", "format": "unix-time", "nullable": true}, "failed_at": {"type": "integer", "description": "Timestamp describing when an OutboundPayment changed status to `failed`.", "format": "unix-time", "nullable": true}, "posted_at": {"type": "integer", "description": "Timestamp describing when an OutboundPayment changed status to `posted`.", "format": "unix-time", "nullable": true}, "returned_at": {"type": "integer", "description": "Timestamp describing when an OutboundPayment changed status to `returned`.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
