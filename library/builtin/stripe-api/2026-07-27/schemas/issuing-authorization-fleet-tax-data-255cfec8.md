---
title: issuing_authorization_fleet_tax_data
page_id: schema-issuing-authorization-fleet-tax-data-255cfec8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_fleet_tax_data

```yaml
{"title": "IssuingAuthorizationFleetTaxData", "type": "object", "properties": {"local_amount_decimal": {"type": "string", "description": "Amount of state or provincial Sales Tax included in the transaction amount. `null` if not reported by merchant or not subject to tax.", "format": "decimal", "nullable": true}, "national_amount_decimal": {"type": "string", "description": "Amount of national Sales Tax or VAT included in the transaction amount. `null` if not reported by merchant or not subject to tax.", "format": "decimal", "nullable": true}}, "description": "", "x-expandableFields": []}
```
