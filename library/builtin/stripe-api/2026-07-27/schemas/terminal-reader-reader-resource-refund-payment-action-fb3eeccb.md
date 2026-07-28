---
title: terminal_reader_reader_resource_refund_payment_action
page_id: schema-terminal-reader-reader-resource-refund-payment-action-fb3eeccb
path: schemas
description: Represents a reader action to refund a payment
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_refund_payment_action

Represents a reader action to refund a payment

```yaml
{"title": "TerminalReaderReaderResourceRefundPaymentAction", "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount being refunded."}, "charge": {"description": "Charge that is being refunded.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "payment_intent": {"description": "Payment intent that is being refunded.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}, "reason": {"type": "string", "description": "The reason for the refund.", "enum": ["duplicate", "fraudulent", "requested_by_customer"]}, "refund": {"description": "Unique identifier for the refund object.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/refund"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/refund"}]}}, "refund_application_fee": {"type": "boolean", "description": "Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge."}, "refund_payment_config": {"$ref": "#/components/schemas/terminal_reader_reader_resource_refund_payment_config"}, "reverse_transfer": {"type": "boolean", "description": "Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount). A transfer can be reversed only by the application that created the charge."}}, "description": "Represents a reader action to refund a payment", "x-expandableFields": ["charge", "payment_intent", "refund", "refund_payment_config"]}
```
