---
title: inbound_transfers
page_id: schema-inbound-transfers-5d7fabbd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# inbound_transfers

```yaml
{"title": "InboundTransfers", "required": ["billing_details", "type"], "type": "object", "properties": {"billing_details": {"$ref": "#/components/schemas/treasury_shared_resource_billing_details"}, "type": {"type": "string", "description": "The type of the payment method used in the InboundTransfer.", "enum": ["us_bank_account"], "x-stripeBypassValidation": true}, "us_bank_account": {"$ref": "#/components/schemas/inbound_transfers_payment_method_details_us_bank_account"}}, "description": "", "x-expandableFields": ["billing_details", "us_bank_account"]}
```
