---
title: terminal_reader_reader_resource_selection
page_id: schema-terminal-reader-reader-resource-selection-23989e83
path: schemas
description: Information about a selection being collected using a reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_selection

Information about a selection being collected using a reader

```yaml
{"title": "TerminalReaderReaderResourceSelection", "required": ["choices"], "type": "object", "properties": {"choices": {"type": "array", "description": "List of possible choices to be selected", "items": {"$ref": "#/components/schemas/terminal_reader_reader_resource_choice"}}, "id": {"maxLength": 5000, "type": "string", "description": "The id of the selected choice", "nullable": true}, "text": {"maxLength": 5000, "type": "string", "description": "The text of the selected choice", "nullable": true}}, "description": "Information about a selection being collected using a reader", "x-expandableFields": ["choices"]}
```
