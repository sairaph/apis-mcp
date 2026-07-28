---
title: EnumCapability
page_id: schema-enumcapability-08794bc6
path: schemas
description: A parameter that accepts one of a discrete set of string values.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# EnumCapability

A parameter that accepts one of a discrete set of string values.

```yaml
{"description": "A parameter that accepts one of a discrete set of string values.", "example": {"type": "enum", "values": ["1K", "2K", "4K"]}, "properties": {"type": {"enum": ["enum"], "type": "string"}, "values": {"items": {"type": "string"}, "type": "array"}}, "required": ["type", "values"], "type": "object"}
```
