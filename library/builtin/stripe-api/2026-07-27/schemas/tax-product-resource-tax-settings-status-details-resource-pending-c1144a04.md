---
title: tax_product_resource_tax_settings_status_details_resource_pending
page_id: schema-tax-product-resource-tax-settings-status-details-resource-pending-c1144a04
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_settings_status_details_resource_pending

```yaml
{"title": "TaxProductResourceTaxSettingsStatusDetailsResourcePending", "type": "object", "properties": {"missing_fields": {"type": "array", "description": "The list of missing fields that are required to perform calculations. It includes the entry `head_office` when the status is `pending`. It is recommended to set the optional values even if they aren't listed as required for calculating taxes. Calculations can fail if missing fields aren't explicitly provided on every call.", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": []}
```
