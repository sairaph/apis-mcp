---
title: terminal_reader_reader_resource_collect_payment_method_action
page_id: schema-terminal-reader-reader-resource-collect-payment-method-action-de2819c2
path: schemas
description: Represents a reader action to collect a payment method
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_collect_payment_method_action

Represents a reader action to collect a payment method

```yaml
{"title": "TerminalReaderReaderResourceCollectPaymentMethodAction", "required": ["payment_intent"], "type": "object", "properties": {"collect_config": {"$ref": "#/components/schemas/terminal_reader_reader_resource_collect_config"}, "payment_intent": {"description": "Most recent PaymentIntent processed by the reader.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}, "payment_method": {"$ref": "#/components/schemas/payment_method"}}, "description": "Represents a reader action to collect a payment method", "x-expandableFields": ["collect_config", "payment_intent", "payment_method"]}
```
