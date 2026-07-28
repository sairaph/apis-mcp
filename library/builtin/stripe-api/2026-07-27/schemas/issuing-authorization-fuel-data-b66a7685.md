---
title: issuing_authorization_fuel_data
page_id: schema-issuing-authorization-fuel-data-b66a7685
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_fuel_data

```yaml
{"title": "IssuingAuthorizationFuelData", "type": "object", "properties": {"industry_product_code": {"maxLength": 5000, "type": "string", "description": "[Conexxus Payment System Product Code](https://www.conexxus.org/conexxus-payment-system-product-codes) identifying the primary fuel product purchased.", "nullable": true}, "quantity_decimal": {"type": "string", "description": "The quantity of `unit`s of fuel that was dispensed, represented as a decimal string with at most 12 decimal places.", "format": "decimal", "nullable": true}, "type": {"type": "string", "description": "The type of fuel that was purchased.", "nullable": true, "enum": ["diesel", "other", "unleaded_plus", "unleaded_regular", "unleaded_super"]}, "unit": {"type": "string", "description": "The units for `quantity_decimal`.", "nullable": true, "enum": ["charging_minute", "imperial_gallon", "kilogram", "kilowatt_hour", "liter", "other", "pound", "us_gallon"]}, "unit_cost_decimal": {"type": "string", "description": "The cost in cents per each unit of fuel, represented as a decimal string with at most 12 decimal places.", "format": "decimal", "nullable": true}}, "description": "", "x-expandableFields": []}
```
