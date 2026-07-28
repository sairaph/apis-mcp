---
title: account_monthly_estimated_revenue
page_id: schema-account-monthly-estimated-revenue-cc62fedd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_monthly_estimated_revenue

```yaml
{"title": "AccountMonthlyEstimatedRevenue", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A non-negative integer representing how much to charge in the [smallest currency unit](/currencies#zero-decimal)."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}}, "description": "", "x-expandableFields": []}
```
