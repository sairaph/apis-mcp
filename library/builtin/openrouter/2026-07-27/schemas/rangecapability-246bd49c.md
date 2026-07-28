---
title: RangeCapability
page_id: schema-rangecapability-246bd49c
path: schemas
description: A parameter that accepts any value within an inclusive numeric range.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# RangeCapability

A parameter that accepts any value within an inclusive numeric range.

```yaml
{"description": "A parameter that accepts any value within an inclusive numeric range.", "example": {"max": 100, "min": 0, "type": "range"}, "properties": {"max": {"type": "number"}, "min": {"type": "number"}, "type": {"enum": ["range"], "type": "string"}}, "required": ["type", "min", "max"], "type": "object"}
```
