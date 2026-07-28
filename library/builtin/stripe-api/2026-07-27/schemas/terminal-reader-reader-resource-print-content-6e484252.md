---
title: terminal_reader_reader_resource_print_content
page_id: schema-terminal-reader-reader-resource-print-content-6e484252
path: schemas
description: Represents a reader action to print content
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_print_content

Represents a reader action to print content

```yaml
{"title": "TerminalReaderReaderResourcePrintContent", "required": ["type"], "type": "object", "properties": {"image": {"$ref": "#/components/schemas/terminal_reader_reader_resource_file_metadata"}, "type": {"type": "string", "description": "The type of content to print. Currently supports `image`.", "enum": ["image"]}}, "description": "Represents a reader action to print content", "x-expandableFields": ["image"]}
```
