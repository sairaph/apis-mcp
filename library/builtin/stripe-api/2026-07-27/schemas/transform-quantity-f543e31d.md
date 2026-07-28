---
title: transform_quantity
page_id: schema-transform-quantity-f543e31d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# transform_quantity

```yaml
{"title": "TransformQuantity", "required": ["divide_by", "round"], "type": "object", "properties": {"divide_by": {"type": "integer", "description": "Divide usage by this number."}, "round": {"type": "string", "description": "After division, either round the result `up` or `down`.", "enum": ["down", "up"]}}, "description": "", "x-expandableFields": []}
```
