---
title: terminal_reader_reader_resource_process_setup_intent_action
page_id: schema-terminal-reader-reader-resource-process-setup-intent-action-1f41754c
path: schemas
description: Represents a reader action to process a setup intent
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_process_setup_intent_action

Represents a reader action to process a setup intent

```yaml
{"title": "TerminalReaderReaderResourceProcessSetupIntentAction", "required": ["setup_intent"], "type": "object", "properties": {"generated_card": {"maxLength": 5000, "type": "string", "description": "ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod."}, "process_config": {"$ref": "#/components/schemas/terminal_reader_reader_resource_process_setup_config"}, "setup_intent": {"description": "Most recent SetupIntent processed by the reader.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/setup_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/setup_intent"}]}}}, "description": "Represents a reader action to process a setup intent", "x-expandableFields": ["process_config", "setup_intent"]}
```
