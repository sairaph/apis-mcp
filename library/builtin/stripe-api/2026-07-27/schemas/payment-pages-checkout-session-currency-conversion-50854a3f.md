---
title: payment_pages_checkout_session_currency_conversion
page_id: schema-payment-pages-checkout-session-currency-conversion-50854a3f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_pages_checkout_session_currency_conversion

```yaml
{"title": "PaymentPagesCheckoutSessionCurrencyConversion", "required": ["amount_subtotal", "amount_total", "fx_rate", "source_currency"], "type": "object", "properties": {"amount_subtotal": {"type": "integer", "description": "Total of all items in source currency before discounts or taxes are applied."}, "amount_total": {"type": "integer", "description": "Total of all items in source currency after discounts and taxes are applied."}, "fx_rate": {"type": "string", "description": "Exchange rate used to convert source currency amounts to customer currency amounts", "format": "decimal"}, "source_currency": {"maxLength": 5000, "type": "string", "description": "Creation currency of the CheckoutSession before localization"}}, "description": "", "x-expandableFields": []}
```
