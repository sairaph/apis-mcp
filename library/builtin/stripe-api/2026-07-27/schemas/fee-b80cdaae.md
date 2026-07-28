---
title: fee
page_id: schema-fee-b80cdaae
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# fee

```yaml
{"title": "Fee", "required": ["amount", "currency", "type"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount of the fee, in cents."}, "application": {"maxLength": 5000, "type": "string", "description": "ID of the Connect application that earned the fee.", "nullable": true}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users.", "nullable": true}, "type": {"maxLength": 5000, "type": "string", "description": "Type of the fee, one of: `application_fee`, `payment_method_passthrough_fee`, `stripe_fee`, `tax`, or `withheld_tax`."}}, "description": "", "x-expandableFields": []}
```
