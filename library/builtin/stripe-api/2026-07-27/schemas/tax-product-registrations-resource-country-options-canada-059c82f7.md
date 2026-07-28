---
title: tax_product_registrations_resource_country_options_canada
page_id: schema-tax-product-registrations-resource-country-options-canada-059c82f7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_registrations_resource_country_options_canada

```yaml
{"title": "TaxProductRegistrationsResourceCountryOptionsCanada", "required": ["type"], "type": "object", "properties": {"province_standard": {"$ref": "#/components/schemas/tax_product_registrations_resource_country_options_ca_province_standard"}, "type": {"type": "string", "description": "Type of registration in Canada.", "enum": ["province_standard", "simplified", "standard"]}}, "description": "", "x-expandableFields": ["province_standard"]}
```
