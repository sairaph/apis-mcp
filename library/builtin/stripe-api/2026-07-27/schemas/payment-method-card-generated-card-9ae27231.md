---
title: payment_method_card_generated_card
page_id: schema-payment-method-card-generated-card-9ae27231
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_card_generated_card

```yaml
{"title": "payment_method_card_generated_card", "type": "object", "properties": {"charge": {"maxLength": 5000, "type": "string", "description": "The charge that created this object.", "nullable": true}, "payment_method_details": {"description": "Transaction-specific details of the payment method used in the payment.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/card_generated_from_payment_method_details"}]}, "setup_attempt": {"description": "The ID of the SetupAttempt that generated this PaymentMethod, if any.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/setup_attempt"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/setup_attempt"}]}}}, "description": "", "x-expandableFields": ["payment_method_details", "setup_attempt"]}
```
