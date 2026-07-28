---
title: billing_credit_grants_resource_monetary_amount
page_id: schema-billing-credit-grants-resource-monetary-amount-aed3bb05
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_credit_grants_resource_monetary_amount

```yaml
{"title": "BillingCreditGrantsResourceMonetaryAmount", "required": ["currency", "value"], "type": "object", "properties": {"currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies)."}, "value": {"type": "integer", "description": "A positive integer representing the amount."}}, "description": "", "x-expandableFields": []}
```
