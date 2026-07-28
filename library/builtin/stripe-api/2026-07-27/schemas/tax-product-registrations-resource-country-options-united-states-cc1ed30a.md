---
title: tax_product_registrations_resource_country_options_united_states
page_id: schema-tax-product-registrations-resource-country-options-united-states-cc1ed30a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_registrations_resource_country_options_united_states

```yaml
{"title": "TaxProductRegistrationsResourceCountryOptionsUnitedStates", "required": ["state", "type"], "type": "object", "properties": {"local_amusement_tax": {"$ref": "#/components/schemas/tax_product_registrations_resource_country_options_us_local_amusement_tax"}, "local_lease_tax": {"$ref": "#/components/schemas/tax_product_registrations_resource_country_options_us_local_lease_tax"}, "state": {"maxLength": 5000, "type": "string", "description": "Two-letter US state code ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2))."}, "state_sales_tax": {"$ref": "#/components/schemas/tax_product_registrations_resource_country_options_us_state_sales_tax"}, "type": {"type": "string", "description": "Type of registration in the US.", "enum": ["local_amusement_tax", "local_lease_tax", "state_communications_tax", "state_retail_delivery_fee", "state_sales_tax"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["local_amusement_tax", "local_lease_tax", "state_sales_tax"]}
```
