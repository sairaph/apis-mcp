---
title: treasury_received_credits_resource_source_flows_details
page_id: schema-treasury-received-credits-resource-source-flows-details-6b043662
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_received_credits_resource_source_flows_details

```yaml
{"title": "TreasuryReceivedCreditsResourceSourceFlowsDetails", "required": ["type"], "type": "object", "properties": {"credit_reversal": {"$ref": "#/components/schemas/treasury.credit_reversal"}, "outbound_payment": {"$ref": "#/components/schemas/treasury.outbound_payment"}, "outbound_transfer": {"$ref": "#/components/schemas/treasury.outbound_transfer"}, "payout": {"$ref": "#/components/schemas/payout"}, "type": {"type": "string", "description": "The type of the source flow that originated the ReceivedCredit.", "enum": ["credit_reversal", "other", "outbound_payment", "outbound_transfer", "payout"]}}, "description": "", "x-expandableFields": ["credit_reversal", "outbound_payment", "outbound_transfer", "payout"]}
```
