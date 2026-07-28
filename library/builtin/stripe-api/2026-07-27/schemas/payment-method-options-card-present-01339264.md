---
title: payment_method_options_card_present
page_id: schema-payment-method-options-card-present-01339264
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_options_card_present

```yaml
{"title": "payment_method_options_card_present", "type": "object", "properties": {"capture_method": {"type": "string", "description": "Controls when the funds will be captured from the customer's account.", "enum": ["manual", "manual_preferred"], "x-stripeBypassValidation": true}, "request_extended_authorization": {"type": "boolean", "description": "Request ability to capture this payment beyond the standard [authorization validity window](https://docs.stripe.com/terminal/features/extended-authorizations#authorization-validity)", "nullable": true}, "request_incremental_authorization_support": {"type": "boolean", "description": "Request ability to [increment](https://docs.stripe.com/terminal/features/incremental-authorizations) this PaymentIntent if the combination of MCC and card brand is eligible. Check [incremental_authorization_supported](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-incremental_authorization_supported) in the [Confirm](https://docs.stripe.com/api/payment_intents/confirm) response to verify support.", "nullable": true}, "routing": {"$ref": "#/components/schemas/payment_method_options_card_present_routing"}}, "description": "", "x-expandableFields": ["routing"]}
```
