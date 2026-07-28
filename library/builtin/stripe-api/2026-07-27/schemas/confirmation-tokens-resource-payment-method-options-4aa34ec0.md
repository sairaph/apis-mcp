---
title: confirmation_tokens_resource_payment_method_options
page_id: schema-confirmation-tokens-resource-payment-method-options-4aa34ec0
path: schemas
description: Payment-method-specific configuration
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# confirmation_tokens_resource_payment_method_options

Payment-method-specific configuration

```yaml
{"title": "ConfirmationTokensResourcePaymentMethodOptions", "type": "object", "properties": {"card": {"description": "This hash contains the card payment method options.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/confirmation_tokens_resource_payment_method_options_resource_card"}]}}, "description": "Payment-method-specific configuration", "x-expandableFields": ["card"]}
```
