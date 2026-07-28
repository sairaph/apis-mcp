---
title: CompoundFilter
page_id: schema-compoundfilter-8d3e4db6
path: schemas
description: A compound filter that combines multiple comparison or compound filters
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CompoundFilter

A compound filter that combines multiple comparison or compound filters

```yaml
{"description": "A compound filter that combines multiple comparison or compound filters", "example": {"filters": [{"key": "author", "type": "eq", "value": "Alice"}], "type": "and"}, "properties": {"filters": {"items": {"additionalProperties": {}, "type": "object"}, "type": "array"}, "type": {"enum": ["and", "or"], "type": "string", "x-speakeasy-unknown-values": "allow"}}, "required": ["type", "filters"], "type": "object"}
```
