---
title: promotion_codes_resource_restrictions
page_id: schema-promotion-codes-resource-restrictions-abf11b84
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# promotion_codes_resource_restrictions

```yaml
{"title": "PromotionCodesResourceRestrictions", "required": ["first_time_transaction"], "type": "object", "properties": {"currency_options": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/promotion_code_currency_option"}, "description": "Promotion code restrictions defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies)."}, "first_time_transaction": {"type": "boolean", "description": "A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices"}, "minimum_amount": {"type": "integer", "description": "Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).", "nullable": true}, "minimum_amount_currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount", "nullable": true}}, "description": "", "x-expandableFields": ["currency_options"]}
```
