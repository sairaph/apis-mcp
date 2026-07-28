---
title: treasury_transactions_resource_flow_details
page_id: schema-treasury-transactions-resource-flow-details-84d0cc21
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_transactions_resource_flow_details

```yaml
{"title": "TreasuryTransactionsResourceFlowDetails", "required": ["type"], "type": "object", "properties": {"credit_reversal": {"$ref": "#/components/schemas/treasury.credit_reversal"}, "debit_reversal": {"$ref": "#/components/schemas/treasury.debit_reversal"}, "inbound_transfer": {"$ref": "#/components/schemas/treasury.inbound_transfer"}, "issuing_authorization": {"$ref": "#/components/schemas/issuing.authorization"}, "outbound_payment": {"$ref": "#/components/schemas/treasury.outbound_payment"}, "outbound_transfer": {"$ref": "#/components/schemas/treasury.outbound_transfer"}, "received_credit": {"$ref": "#/components/schemas/treasury.received_credit"}, "received_debit": {"$ref": "#/components/schemas/treasury.received_debit"}, "type": {"type": "string", "description": "Type of the flow that created the Transaction. Set to the same value as `flow_type`.", "enum": ["credit_reversal", "debit_reversal", "inbound_transfer", "issuing_authorization", "other", "outbound_payment", "outbound_transfer", "received_credit", "received_debit"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["credit_reversal", "debit_reversal", "inbound_transfer", "issuing_authorization", "outbound_payment", "outbound_transfer", "received_credit", "received_debit"]}
```
