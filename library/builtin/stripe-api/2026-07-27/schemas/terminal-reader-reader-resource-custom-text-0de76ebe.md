---
title: terminal_reader_reader_resource_custom_text
page_id: schema-terminal-reader-reader-resource-custom-text-0de76ebe
path: schemas
description: Represents custom text to be displayed when collecting the input using a reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_custom_text

Represents custom text to be displayed when collecting the input using a reader

```yaml
{"title": "TerminalReaderReaderResourceCustomText", "type": "object", "properties": {"description": {"maxLength": 5000, "type": "string", "description": "Customize the default description for this input", "nullable": true}, "skip_button": {"maxLength": 5000, "type": "string", "description": "Customize the default label for this input's skip button", "nullable": true}, "submit_button": {"maxLength": 5000, "type": "string", "description": "Customize the default label for this input's submit button", "nullable": true}, "title": {"maxLength": 5000, "type": "string", "description": "Customize the default title for this input", "nullable": true}}, "description": "Represents custom text to be displayed when collecting the input using a reader", "x-expandableFields": []}
```
