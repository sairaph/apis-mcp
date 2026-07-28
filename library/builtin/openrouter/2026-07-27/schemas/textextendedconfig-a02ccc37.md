---
title: TextExtendedConfig
page_id: schema-textextendedconfig-a02ccc37
path: schemas
description: Text output configuration including format and verbosity
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TextExtendedConfig

Text output configuration including format and verbosity

```yaml
{"allOf": [{"$ref": "#/components/schemas/TextConfig"}, {"properties": {"verbosity": {"enum": ["low", "medium", "high", "xhigh", "max", null], "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}}, "type": "object"}], "description": "Text output configuration including format and verbosity", "example": {"format": {"type": "text"}}}
```
