---
title: payments_primitives_payment_records_resource_amount
page_id: schema-payments-primitives-payment-records-resource-amount-39c8901e
path: schemas
description: A representation of an amount of money, consisting of an amount and a currency.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_amount

A representation of an amount of money, consisting of an amount and a currency.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceAmount", "required": ["currency", "value"], "type": "object", "properties": {"currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "value": {"type": "integer", "description": "A positive integer representing the amount in the currency's [minor unit](https://docs.stripe.com/currencies#zero-decimal). For example, `100` can represent 1 USD or 100 JPY."}}, "description": "A representation of an amount of money, consisting of an amount and a currency.", "x-expandableFields": []}
```
