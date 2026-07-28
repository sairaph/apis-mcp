---
title: discount_source
page_id: schema-discount-source-d3c214d9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# discount_source

```yaml
{"title": "DiscountSource", "required": ["type"], "type": "object", "properties": {"coupon": {"description": "The coupon that was redeemed to create this discount.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/coupon"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/coupon"}]}}, "type": {"type": "string", "description": "The source type of the discount.", "enum": ["coupon"]}}, "description": "", "x-expandableFields": ["coupon"]}
```
