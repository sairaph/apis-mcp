---
title: tax_product_resource_postal_address
page_id: schema-tax-product-resource-postal-address-22931e24
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_postal_address

```yaml
{"title": "TaxProductResourcePostalAddress", "required": ["country"], "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string", "description": "City, district, suburb, town, or village.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2))."}, "line1": {"maxLength": 5000, "type": "string", "description": "Address line 1, such as the street, PO Box, or company name.", "nullable": true}, "line2": {"maxLength": 5000, "type": "string", "description": "Address line 2, such as the apartment, suite, unit, or building.", "nullable": true}, "postal_code": {"maxLength": 5000, "type": "string", "description": "ZIP or postal code.", "nullable": true}, "state": {"maxLength": 5000, "type": "string", "description": "State/province as an [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) subdivision code, without country prefix, such as \"NY\" or \"TX\".", "nullable": true}}, "description": "", "x-expandableFields": []}
```
