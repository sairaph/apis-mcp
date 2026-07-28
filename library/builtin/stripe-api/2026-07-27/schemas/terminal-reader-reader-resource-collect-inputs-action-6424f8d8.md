---
title: terminal_reader_reader_resource_collect_inputs_action
page_id: schema-terminal-reader-reader-resource-collect-inputs-action-6424f8d8
path: schemas
description: Represents a reader action to collect customer inputs
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_collect_inputs_action

Represents a reader action to collect customer inputs

```yaml
{"title": "TerminalReaderReaderResourceCollectInputsAction", "required": ["inputs"], "type": "object", "properties": {"inputs": {"type": "array", "description": "List of inputs to be collected.", "items": {"$ref": "#/components/schemas/terminal_reader_reader_resource_input"}}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}}, "description": "Represents a reader action to collect customer inputs", "x-expandableFields": ["inputs"]}
```
