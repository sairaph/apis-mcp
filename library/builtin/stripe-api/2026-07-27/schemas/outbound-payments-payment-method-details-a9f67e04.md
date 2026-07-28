---
title: outbound_payments_payment_method_details
page_id: schema-outbound-payments-payment-method-details-a9f67e04
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# outbound_payments_payment_method_details

```yaml
{"title": "OutboundPaymentsPaymentMethodDetails", "required": ["billing_details", "type"], "type": "object", "properties": {"billing_details": {"$ref": "#/components/schemas/treasury_shared_resource_billing_details"}, "financial_account": {"$ref": "#/components/schemas/outbound_payments_payment_method_details_financial_account"}, "type": {"type": "string", "description": "The type of the payment method used in the OutboundPayment.", "enum": ["financial_account", "us_bank_account"], "x-stripeBypassValidation": true}, "us_bank_account": {"$ref": "#/components/schemas/outbound_payments_payment_method_details_us_bank_account"}}, "description": "", "x-expandableFields": ["billing_details", "financial_account", "us_bank_account"]}
```
