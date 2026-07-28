---
title: terminal_reader_reader_resource_choice
page_id: schema-terminal-reader-reader-resource-choice-49a83828
path: schemas
description: Choice to be selected on a Reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_choice

Choice to be selected on a Reader

```yaml
{"title": "TerminalReaderReaderResourceChoice", "required": ["text"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "The identifier for the selected choice. Maximum 50 characters.", "nullable": true}, "style": {"type": "string", "description": "The button style for the choice. Can be `primary` or `secondary`.", "nullable": true, "enum": ["primary", "secondary"]}, "text": {"maxLength": 5000, "type": "string", "description": "The text to be selected. Maximum 30 characters."}}, "description": "Choice to be selected on a Reader", "x-expandableFields": []}
```
