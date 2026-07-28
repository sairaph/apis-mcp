---
title: outbound_transfers_payment_method_details
page_id: schema-outbound-transfers-payment-method-details-c2b545a3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# outbound_transfers_payment_method_details

```yaml
{"title": "OutboundTransfersPaymentMethodDetails", "required": ["billing_details", "type"], "type": "object", "properties": {"billing_details": {"$ref": "#/components/schemas/treasury_shared_resource_billing_details"}, "financial_account": {"$ref": "#/components/schemas/outbound_transfers_payment_method_details_financial_account"}, "type": {"type": "string", "description": "The type of the payment method used in the OutboundTransfer.", "enum": ["financial_account", "us_bank_account"]}, "us_bank_account": {"$ref": "#/components/schemas/outbound_transfers_payment_method_details_us_bank_account"}}, "description": "", "x-expandableFields": ["billing_details", "financial_account", "us_bank_account"]}
```
