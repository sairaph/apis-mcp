---
title: tax_product_resource_jurisdiction
page_id: schema-tax-product-resource-jurisdiction-464a8d5c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_jurisdiction

```yaml
{"title": "TaxProductResourceJurisdiction", "required": ["country", "display_name", "level"], "type": "object", "properties": {"country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2))."}, "display_name": {"maxLength": 5000, "type": "string", "description": "A human-readable name for the jurisdiction imposing the tax."}, "level": {"type": "string", "description": "Indicates the level of the jurisdiction imposing the tax.", "enum": ["city", "country", "county", "district", "state"]}, "state": {"maxLength": 5000, "type": "string", "description": "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2), without country prefix. For example, \"NY\" for New York, United States.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
