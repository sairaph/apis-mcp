---
title: terminal_reader_reader_resource_process_payment_intent_action
page_id: schema-terminal-reader-reader-resource-process-payment-intent-action-21191522
path: schemas
description: Represents a reader action to process a payment intent
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_process_payment_intent_action

Represents a reader action to process a payment intent

```yaml
{"title": "TerminalReaderReaderResourceProcessPaymentIntentAction", "required": ["payment_intent"], "type": "object", "properties": {"payment_intent": {"description": "Most recent PaymentIntent processed by the reader.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}, "process_config": {"$ref": "#/components/schemas/terminal_reader_reader_resource_process_config"}}, "description": "Represents a reader action to process a payment intent", "x-expandableFields": ["payment_intent", "process_config"]}
```
