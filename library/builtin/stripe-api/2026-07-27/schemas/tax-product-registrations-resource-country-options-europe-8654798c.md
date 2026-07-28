---
title: tax_product_registrations_resource_country_options_europe
page_id: schema-tax-product-registrations-resource-country-options-europe-8654798c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_registrations_resource_country_options_europe

```yaml
{"title": "TaxProductRegistrationsResourceCountryOptionsEurope", "required": ["type"], "type": "object", "properties": {"standard": {"$ref": "#/components/schemas/tax_product_registrations_resource_country_options_eu_standard"}, "type": {"type": "string", "description": "Type of registration in an EU country.", "enum": ["ioss", "oss_non_union", "oss_union", "standard"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["standard"]}
```
