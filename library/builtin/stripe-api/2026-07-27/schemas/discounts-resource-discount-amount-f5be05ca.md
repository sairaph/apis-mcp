---
title: discounts_resource_discount_amount
page_id: schema-discounts-resource-discount-amount-f5be05ca
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# discounts_resource_discount_amount

```yaml
{"title": "DiscountsResourceDiscountAmount", "required": ["amount", "discount"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount, in cents (or local equivalent), of the discount."}, "discount": {"description": "The discount that was applied to get this discount amount.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/discount"}, {"$ref": "#/components/schemas/deleted_discount"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/discount"}, {"$ref": "#/components/schemas/deleted_discount"}]}}}, "description": "", "x-expandableFields": ["discount"]}
```
