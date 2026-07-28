---
title: stackable_discount_with_discount_settings_and_discount_end
page_id: schema-stackable-discount-with-discount-settings-and-discount-end-3aae910c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# stackable_discount_with_discount_settings_and_discount_end

```yaml
{"title": "StackableDiscountWithDiscountSettingsAndDiscountEnd", "type": "object", "properties": {"coupon": {"description": "ID of the coupon to create a new discount for.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/coupon"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/coupon"}]}}, "discount": {"description": "ID of an existing discount on the object (or one of its ancestors) to reuse.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/discount"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/discount"}]}}, "promotion_code": {"description": "ID of the promotion code to create a new discount for.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/promotion_code"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/promotion_code"}]}}}, "description": "", "x-expandableFields": ["coupon", "discount", "promotion_code"]}
```
