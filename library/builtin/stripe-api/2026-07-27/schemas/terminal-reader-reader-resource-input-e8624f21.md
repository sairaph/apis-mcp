---
title: terminal_reader_reader_resource_input
page_id: schema-terminal-reader-reader-resource-input-e8624f21
path: schemas
description: Represents an input to be collected using the reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_input

Represents an input to be collected using the reader

```yaml
{"title": "TerminalReaderReaderResourceInput", "required": ["type"], "type": "object", "properties": {"custom_text": {"description": "Default text of input being collected.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/terminal_reader_reader_resource_custom_text"}]}, "email": {"$ref": "#/components/schemas/terminal_reader_reader_resource_email"}, "numeric": {"$ref": "#/components/schemas/terminal_reader_reader_resource_numeric"}, "phone": {"$ref": "#/components/schemas/terminal_reader_reader_resource_phone"}, "required": {"type": "boolean", "description": "Indicate that this input is required, disabling the skip button.", "nullable": true}, "selection": {"$ref": "#/components/schemas/terminal_reader_reader_resource_selection"}, "signature": {"$ref": "#/components/schemas/terminal_reader_reader_resource_signature"}, "skipped": {"type": "boolean", "description": "Indicate that this input was skipped by the user."}, "text": {"$ref": "#/components/schemas/terminal_reader_reader_resource_text"}, "toggles": {"type": "array", "description": "List of toggles being collected. Values are present if collection is complete.", "nullable": true, "items": {"$ref": "#/components/schemas/terminal_reader_reader_resource_toggle"}}, "type": {"type": "string", "description": "Type of input being collected.", "enum": ["email", "numeric", "phone", "selection", "signature", "text"]}}, "description": "Represents an input to be collected using the reader", "x-expandableFields": ["custom_text", "email", "numeric", "phone", "selection", "signature", "text", "toggles"]}
```
