---
title: balance_amount_net
page_id: schema-balance-amount-net-52e5791d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_amount_net

```yaml
{"title": "BalanceAmountNet", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Balance amount."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "net_available": {"type": "array", "description": "Breakdown of balance by destination.", "items": {"$ref": "#/components/schemas/balance_net_available"}}, "source_types": {"$ref": "#/components/schemas/balance_amount_by_source_type"}}, "description": "", "x-expandableFields": ["net_available", "source_types"]}
```
