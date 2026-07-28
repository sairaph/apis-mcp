---
title: account_annual_revenue
page_id: schema-account-annual-revenue-c43048f4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_annual_revenue

```yaml
{"title": "AccountAnnualRevenue", "type": "object", "properties": {"amount": {"type": "integer", "description": "A non-negative integer representing the amount in the [smallest currency unit](/currencies#zero-decimal).", "nullable": true}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency", "nullable": true}, "fiscal_year_end": {"maxLength": 5000, "type": "string", "description": "The close-out date of the preceding fiscal year in ISO 8601 format. E.g. 2023-12-31 for the 31st of December, 2023.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
