---
title: tax_product_registrations_resource_country_options_us_state_sales_tax_election
page_id: schema-tax-product-registrations-resource-country-options-us-state-sales-tax-el-244c7a54
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_registrations_resource_country_options_us_state_sales_tax_election

```yaml
{"title": "TaxProductRegistrationsResourceCountryOptionsUsStateSalesTaxElection", "required": ["type"], "type": "object", "properties": {"jurisdiction": {"maxLength": 5000, "type": "string", "description": "A [FIPS code](https://www.census.gov/library/reference/code-lists/ansi.html) representing the local jurisdiction."}, "type": {"type": "string", "description": "The type of the election for the state sales tax registration.", "enum": ["local_use_tax", "simplified_sellers_use_tax", "single_local_use_tax"]}}, "description": "", "x-expandableFields": []}
```
