---
title: terminal_reader_reader_resource_toggle
page_id: schema-terminal-reader-reader-resource-toggle-6aa36358
path: schemas
description: Information about an input's toggle
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_toggle

Information about an input's toggle

```yaml
{"title": "TerminalReaderReaderResourceToggle", "type": "object", "properties": {"default_value": {"type": "string", "description": "The toggle's default value. Can be `enabled` or `disabled`.", "nullable": true, "enum": ["disabled", "enabled"]}, "description": {"maxLength": 5000, "type": "string", "description": "The toggle's description text. Maximum 50 characters.", "nullable": true}, "title": {"maxLength": 5000, "type": "string", "description": "The toggle's title text. Maximum 50 characters.", "nullable": true}, "value": {"type": "string", "description": "The toggle's collected value. Can be `enabled` or `disabled`.", "nullable": true, "enum": ["disabled", "enabled"]}}, "description": "Information about an input's toggle", "x-expandableFields": []}
```
