---
title: confirmation_tokens_resource_payment_method_options_resource_card
page_id: schema-confirmation-tokens-resource-payment-method-options-resource-card-cd4f6194
path: schemas
description: This hash contains the card payment method options.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# confirmation_tokens_resource_payment_method_options_resource_card

This hash contains the card payment method options.

```yaml
{"title": "ConfirmationTokensResourcePaymentMethodOptionsResourceCard", "type": "object", "properties": {"cvc_token": {"maxLength": 5000, "type": "string", "description": "The `cvc_update` Token collected from the Payment Element.", "nullable": true}, "installments": {"$ref": "#/components/schemas/confirmation_tokens_resource_payment_method_options_resource_card_resource_installment"}}, "description": "This hash contains the card payment method options.", "x-expandableFields": ["installments"]}
```
