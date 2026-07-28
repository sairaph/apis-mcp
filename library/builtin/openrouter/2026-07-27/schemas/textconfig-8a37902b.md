---
title: TextConfig
page_id: schema-textconfig-8a37902b
path: schemas
description: Text output configuration including format and verbosity
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TextConfig

Text output configuration including format and verbosity

```yaml
{"description": "Text output configuration including format and verbosity", "example": {"format": {"type": "text"}, "verbosity": "medium"}, "properties": {"format": {"$ref": "#/components/schemas/Formats"}, "verbosity": {"enum": ["high", "low", "medium", null], "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}}, "type": "object"}
```
