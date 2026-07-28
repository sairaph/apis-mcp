---
title: exchange_rate
page_id: schema-exchange-rate-c15ee584
path: schemas
description: |-
    [Deprecated] The `ExchangeRate` APIs are deprecated. Please use the [FX Quotes API](https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api) instead.

    `ExchangeRate` objects allow you to determine the rates that Stripe is currently
    using to convert from one currency to another. Since this number is variable
    throughout the day, there are various reasons why you might want to know the current
    rate (for example, to dynamically price an item for a user with a default
    payment in a foreign currency).

    Please refer to our [Exchange Rates API](https://docs.stripe.com/fx-rates) guide for more details.

    *[Note: this integration path is supported but no longer recommended]* Additionally,
    you can guarantee that a charge is made with an exchange rate that you expect is
    current. To do so, you must pass in the exchange_rate to charges endpoints. If the
    value is no longer up to date, the charge won't go through. Please refer to our
    [Using with charges](https://docs.stripe.com/exchange-rates) guide for more details.

    -----

    &nbsp;

    *This Exchange Rates API is a Beta Service and is subject to Stripe's terms of service. You may use the API solely for the purpose of transacting on Stripe. For example, the API may be queried in order to:*

    - *localize prices for processing payments on Stripe*
    - *reconcile Stripe transactions*
    - *determine how much money to send to a connected account*
    - *determine app fees to charge a connected account*

    *Using this Exchange Rates API beta for any purpose other than to transact on Stripe is strictly prohibited and constitutes a violation of Stripe's terms of service.*
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# exchange_rate

[Deprecated] The `ExchangeRate` APIs are deprecated. Please use the [FX Quotes API](https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api) instead.

`ExchangeRate` objects allow you to determine the rates that Stripe is currently
using to convert from one currency to another. Since this number is variable
throughout the day, there are various reasons why you might want to know the current
rate (for example, to dynamically price an item for a user with a default
payment in a foreign currency).

Please refer to our [Exchange Rates API](https://docs.stripe.com/fx-rates) guide for more details.

*[Note: this integration path is supported but no longer recommended]* Additionally,
you can guarantee that a charge is made with an exchange rate that you expect is
current. To do so, you must pass in the exchange_rate to charges endpoints. If the
value is no longer up to date, the charge won't go through. Please refer to our
[Using with charges](https://docs.stripe.com/exchange-rates) guide for more details.

-----

&nbsp;

*This Exchange Rates API is a Beta Service and is subject to Stripe's terms of service. You may use the API solely for the purpose of transacting on Stripe. For example, the API may be queried in order to:*

- *localize prices for processing payments on Stripe*
- *reconcile Stripe transactions*
- *determine how much money to send to a connected account*
- *determine app fees to charge a connected account*

*Using this Exchange Rates API beta for any purpose other than to transact on Stripe is strictly prohibited and constitutes a violation of Stripe's terms of service.*

```yaml
{"title": "ExchangeRate", "required": ["id", "object", "rates"], "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object. Represented as the three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) in lowercase."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["exchange_rate"]}, "rates": {"type": "object", "additionalProperties": {"type": "number"}, "description": "Hash where the keys are supported currencies and the values are the exchange rate at which the base id currency converts to the key currency."}}, "description": "[Deprecated] The `ExchangeRate` APIs are deprecated. Please use the [FX Quotes API](https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api) instead.\n\n`ExchangeRate` objects allow you to determine the rates that Stripe is currently\nusing to convert from one currency to another. Since this number is variable\nthroughout the day, there are various reasons why you might want to know the current\nrate (for example, to dynamically price an item for a user with a default\npayment in a foreign currency).\n\nPlease refer to our [Exchange Rates API](https://docs.stripe.com/fx-rates) guide for more details.\n\n*[Note: this integration path is supported but no longer recommended]* Additionally,\nyou can guarantee that a charge is made with an exchange rate that you expect is\ncurrent. To do so, you must pass in the exchange_rate to charges endpoints. If the\nvalue is no longer up to date, the charge won't go through. Please refer to our\n[Using with charges](https://docs.stripe.com/exchange-rates) guide for more details.\n\n-----\n\n&nbsp;\n\n*This Exchange Rates API is a Beta Service and is subject to Stripe's terms of service. You may use the API solely for the purpose of transacting on Stripe. For example, the API may be queried in order to:*\n\n- *localize prices for processing payments on Stripe*\n- *reconcile Stripe transactions*\n- *determine how much money to send to a connected account*\n- *determine app fees to charge a connected account*\n\n*Using this Exchange Rates API beta for any purpose other than to transact on Stripe is strictly prohibited and constitutes a violation of Stripe's terms of service.*", "x-expandableFields": [], "x-resourceId": "exchange_rate"}
```
