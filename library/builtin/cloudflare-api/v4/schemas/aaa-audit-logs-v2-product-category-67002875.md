---
title: aaa_audit-logs-v2-product-category
page_id: schema-aaa-audit-logs-v2-product-category-67002875
path: schemas
description: A predefined product category and the resource products it expands to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-logs-v2-product-category

A predefined product category and the resource products it expands to.

```yaml
{"description": "A predefined product category and the resource products it expands to.", "type": "object", "properties": {"label": {"description": "A human-readable label for the product category.", "type": "string", "example": "Zero Trust"}, "products": {"description": "The resource products that the product category expands to.", "type": "array", "items": {"$ref": "#/components/schemas/aaa_audit-logs-v2-product-category-item"}}, "value": {"description": "The product category identifier used with the product_category filter.", "type": "string", "example": "zerotrust"}}}
```
