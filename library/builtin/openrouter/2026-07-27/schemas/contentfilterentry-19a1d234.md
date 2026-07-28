---
title: ContentFilterEntry
page_id: schema-contentfilterentry-19a1d234
path: schemas
description: A custom regex content filter that scans request messages for matching patterns.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ContentFilterEntry

A custom regex content filter that scans request messages for matching patterns.

```yaml
{"description": "A custom regex content filter that scans request messages for matching patterns.", "example": {"action": "redact", "label": "[API_KEY]", "pattern": "\\b(sk-[a-zA-Z0-9]{48})\\b"}, "properties": {"action": {"$ref": "#/components/schemas/ContentFilterAction"}, "label": {"description": "Optional label used in redaction placeholders or error messages", "example": "[API_KEY]", "maxLength": 100, "type": ["string", "null"]}, "pattern": {"description": "A regex pattern to match against request content", "example": "\\b(sk-[a-zA-Z0-9]{48})\\b", "minLength": 1, "type": "string"}}, "required": ["pattern", "action"], "type": "object"}
```
