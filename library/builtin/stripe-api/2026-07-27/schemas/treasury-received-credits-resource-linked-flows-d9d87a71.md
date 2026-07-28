---
title: treasury_received_credits_resource_linked_flows
page_id: schema-treasury-received-credits-resource-linked-flows-d9d87a71
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_received_credits_resource_linked_flows

```yaml
{"title": "TreasuryReceivedCreditsResourceLinkedFlows", "type": "object", "properties": {"credit_reversal": {"maxLength": 5000, "type": "string", "description": "The CreditReversal created as a result of this ReceivedCredit being reversed.", "nullable": true}, "issuing_authorization": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedCredit was created due to an [Issuing Authorization](https://api.stripe.com#issuing_authorizations) object.", "nullable": true}, "issuing_transaction": {"maxLength": 5000, "type": "string", "description": "Set if the ReceivedCredit is also viewable as an [Issuing transaction](https://api.stripe.com#issuing_transactions) object.", "nullable": true}, "source_flow": {"maxLength": 5000, "type": "string", "description": "ID of the source flow. Set if `network` is `stripe` and the source flow is visible to the user. Examples of source flows include OutboundPayments, payouts, or CreditReversals.", "nullable": true}, "source_flow_details": {"description": "The expandable object of the source flow.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/treasury_received_credits_resource_source_flows_details"}]}, "source_flow_type": {"maxLength": 5000, "type": "string", "description": "The type of flow that originated the ReceivedCredit (for example, `outbound_payment`).", "nullable": true}}, "description": "", "x-expandableFields": ["source_flow_details"]}
```
