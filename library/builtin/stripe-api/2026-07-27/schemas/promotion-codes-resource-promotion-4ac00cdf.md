---
title: promotion_codes_resource_promotion
page_id: schema-promotion-codes-resource-promotion-4ac00cdf
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# promotion_codes_resource_promotion

```yaml
{"title": "PromotionCodesResourcePromotion", "required": ["type"], "type": "object", "properties": {"coupon": {"description": "If promotion `type` is `coupon`, the coupon for this promotion.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/coupon"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/coupon"}]}}, "type": {"type": "string", "description": "The type of promotion.", "enum": ["coupon"]}}, "description": "", "x-expandableFields": ["coupon"]}
```
