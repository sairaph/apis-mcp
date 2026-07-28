---
title: terminal_reader_reader_resource_confirm_payment_intent_action
page_id: schema-terminal-reader-reader-resource-confirm-payment-intent-action-dd5b1b5d
path: schemas
description: Represents a reader action to confirm a payment
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_confirm_payment_intent_action

Represents a reader action to confirm a payment

```yaml
{"title": "TerminalReaderReaderResourceConfirmPaymentIntentAction", "required": ["payment_intent"], "type": "object", "properties": {"confirm_config": {"$ref": "#/components/schemas/terminal_reader_reader_resource_confirm_config"}, "payment_intent": {"description": "Most recent PaymentIntent processed by the reader.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}}, "description": "Represents a reader action to confirm a payment", "x-expandableFields": ["confirm_config", "payment_intent"]}
```
