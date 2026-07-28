---
title: terminal_reader_reader_resource_line_item
page_id: schema-terminal-reader-reader-resource-line-item-3ad6765a
path: schemas
description: Represents a line item to be displayed on the reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_line_item

Represents a line item to be displayed on the reader

```yaml
{"title": "TerminalReaderReaderResourceLineItem", "required": ["amount", "description", "quantity"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount of the line item. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}, "description": {"maxLength": 5000, "type": "string", "description": "Description of the line item."}, "quantity": {"type": "integer", "description": "The quantity of the line item."}}, "description": "Represents a line item to be displayed on the reader", "x-expandableFields": []}
```
