---
title: balance_amount_by_source_type
page_id: schema-balance-amount-by-source-type-b465c115
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_amount_by_source_type

```yaml
{"title": "BalanceAmountBySourceType", "type": "object", "properties": {"bank_account": {"type": "integer", "description": "Amount coming from [legacy US ACH payments](https://docs.stripe.com/ach-deprecated)."}, "card": {"type": "integer", "description": "Amount coming from most payment methods, including cards as well as [non-legacy bank debits](https://docs.stripe.com/payments/bank-debits)."}, "fpx": {"type": "integer", "description": "Amount coming from [FPX](https://docs.stripe.com/payments/fpx), a Malaysian payment method."}}, "description": "", "x-expandableFields": []}
```
