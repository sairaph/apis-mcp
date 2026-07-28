---
title: issuing_transaction_fuel_data
page_id: schema-issuing-transaction-fuel-data-ac1c58ac
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_fuel_data

```yaml
{"title": "IssuingTransactionFuelData", "required": ["type", "unit", "unit_cost_decimal"], "type": "object", "properties": {"industry_product_code": {"maxLength": 5000, "type": "string", "description": "[Conexxus Payment System Product Code](https://www.conexxus.org/conexxus-payment-system-product-codes) identifying the primary fuel product purchased.", "nullable": true}, "quantity_decimal": {"type": "string", "description": "The quantity of `unit`s of fuel that was dispensed, represented as a decimal string with at most 12 decimal places.", "format": "decimal", "nullable": true}, "type": {"maxLength": 5000, "type": "string", "description": "The type of fuel that was purchased. One of `diesel`, `unleaded_plus`, `unleaded_regular`, `unleaded_super`, or `other`."}, "unit": {"maxLength": 5000, "type": "string", "description": "The units for `quantity_decimal`. One of `charging_minute`, `imperial_gallon`, `kilogram`, `kilowatt_hour`, `liter`, `pound`, `us_gallon`, or `other`."}, "unit_cost_decimal": {"type": "string", "description": "The cost in cents per each unit of fuel, represented as a decimal string with at most 12 decimal places.", "format": "decimal"}}, "description": "", "x-expandableFields": []}
```
