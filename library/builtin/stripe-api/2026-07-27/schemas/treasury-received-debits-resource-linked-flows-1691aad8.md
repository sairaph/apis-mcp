---
title: treasury_received_debits_resource_linked_flows
page_id: schema-treasury-received-debits-resource-linked-flows-1691aad8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_received_debits_resource_linked_flows

```yaml
{"title": "TreasuryReceivedDebitsResourceLinkedFlows", "type": "object", "properties": {"debit_reversal": {"maxLength": 5000, "type": "string", "description": "The DebitReversal created as a result of this ReceivedDebit being reversed.", "nullable": true}, "inbound_transfer": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedDebit is associated with an InboundTransfer's return of funds.", "nullable": true}, "issuing_authorization": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedDebit was created due to an [Issuing Authorization](https://api.stripe.com#issuing_authorizations) object.", "nullable": true}, "issuing_transaction": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedDebit is also viewable as an [Issuing Dispute](https://api.stripe.com#issuing_disputes) object.", "nullable": true}, "payout": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedDebit was created due to a [Payout](https://api.stripe.com#payouts) object.", "nullable": true}, "topup": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedDebit was created due to a [Topup](https://api.stripe.com#topups) object.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
