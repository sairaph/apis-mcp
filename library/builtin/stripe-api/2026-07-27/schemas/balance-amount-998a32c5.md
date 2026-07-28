---
title: balance_amount
page_id: schema-balance-amount-998a32c5
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_amount

```yaml
{"title": "BalanceAmount", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Balance amount."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "source_types": {"$ref": "#/components/schemas/balance_amount_by_source_type"}}, "description": "", "x-expandableFields": ["source_types"]}
```
