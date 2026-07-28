---
title: terminal_reader_reader_resource_set_reader_display_action
page_id: schema-terminal-reader-reader-resource-set-reader-display-action-818e0efc
path: schemas
description: Represents a reader action to set the reader display
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_set_reader_display_action

Represents a reader action to set the reader display

```yaml
{"title": "TerminalReaderReaderResourceSetReaderDisplayAction", "required": ["type"], "type": "object", "properties": {"cart": {"description": "Cart object to be displayed by the reader, including line items, amounts, and currency.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/terminal_reader_reader_resource_cart"}]}, "type": {"type": "string", "description": "Type of information to be displayed by the reader. Only `cart` is currently supported.", "enum": ["cart"]}}, "description": "Represents a reader action to set the reader display", "x-expandableFields": ["cart"]}
```
