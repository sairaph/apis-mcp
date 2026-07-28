---
title: treasury_shared_resource_initiating_payment_method_details_initiating_payment_method_details
page_id: schema-treasury-shared-resource-initiating-payment-method-details-initiating-pa-e9210b14
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_shared_resource_initiating_payment_method_details_initiating_payment_method_details

```yaml
{"title": "TreasurySharedResourceInitiatingPaymentMethodDetailsInitiatingPaymentMethodDetails", "required": ["billing_details", "type"], "type": "object", "properties": {"balance": {"type": "string", "description": "Set when `type` is `balance`.", "enum": ["payments"]}, "billing_details": {"$ref": "#/components/schemas/treasury_shared_resource_billing_details"}, "financial_account": {"$ref": "#/components/schemas/received_payment_method_details_financial_account"}, "issuing_card": {"maxLength": 5000, "type": "string", "description": "Set when `type` is `issuing_card`. This is an [Issuing Card](https://api.stripe.com#issuing_cards) ID."}, "type": {"type": "string", "description": "Polymorphic type matching the originating money movement's source. This can be an external account, a Stripe balance, or a FinancialAccount.", "enum": ["balance", "financial_account", "issuing_card", "stripe", "us_bank_account"], "x-stripeBypassValidation": true}, "us_bank_account": {"$ref": "#/components/schemas/treasury_shared_resource_initiating_payment_method_details_us_bank_account"}}, "description": "", "x-expandableFields": ["billing_details", "financial_account", "us_bank_account"]}
```
