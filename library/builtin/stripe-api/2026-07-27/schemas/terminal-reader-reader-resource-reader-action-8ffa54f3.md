---
title: terminal_reader_reader_resource_reader_action
page_id: schema-terminal-reader-reader-resource-reader-action-8ffa54f3
path: schemas
description: Represents an action performed by the reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_reader_action

Represents an action performed by the reader

```yaml
{"title": "TerminalReaderReaderResourceReaderAction", "required": ["status", "type"], "type": "object", "properties": {"api_error": {"description": "The reader action failed due to an [API error](https://docs.stripe.com/api/errors). Only present when `status` is `failed` and the underlying failure was an API error. Avoid parsing the `message` field for programmatic logic; use `type` or `code` instead. The `message` field is for display to humans only and may be updated at anytime. Requires [reader version](https://docs.stripe.com/terminal/readers/stripe-reader-s700-s710#reader-software-version) 2.42 or later. Readers on older versions always return null.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/api_errors"}]}, "collect_inputs": {"$ref": "#/components/schemas/terminal_reader_reader_resource_collect_inputs_action"}, "collect_payment_method": {"$ref": "#/components/schemas/terminal_reader_reader_resource_collect_payment_method_action"}, "confirm_payment_intent": {"$ref": "#/components/schemas/terminal_reader_reader_resource_confirm_payment_intent_action"}, "failure_code": {"maxLength": 5000, "type": "string", "description": "Failure code, only set if status is `failed`.", "nullable": true}, "failure_message": {"maxLength": 5000, "type": "string", "description": "Detailed failure message, only set if status is `failed`.", "nullable": true}, "print_content": {"$ref": "#/components/schemas/terminal_reader_reader_resource_print_content"}, "process_payment_intent": {"$ref": "#/components/schemas/terminal_reader_reader_resource_process_payment_intent_action"}, "process_setup_intent": {"$ref": "#/components/schemas/terminal_reader_reader_resource_process_setup_intent_action"}, "refund_payment": {"$ref": "#/components/schemas/terminal_reader_reader_resource_refund_payment_action"}, "set_reader_display": {"$ref": "#/components/schemas/terminal_reader_reader_resource_set_reader_display_action"}, "status": {"type": "string", "description": "Status of the action performed by the reader.", "enum": ["failed", "in_progress", "succeeded"]}, "type": {"type": "string", "description": "Type of action performed by the reader.", "enum": ["collect_inputs", "collect_payment_method", "confirm_payment_intent", "print_content", "process_payment_intent", "process_setup_intent", "refund_payment", "set_reader_display"], "x-stripeBypassValidation": true}}, "description": "Represents an action performed by the reader", "x-expandableFields": ["api_error", "collect_inputs", "collect_payment_method", "confirm_payment_intent", "print_content", "process_payment_intent", "process_setup_intent", "refund_payment", "set_reader_display"]}
```
