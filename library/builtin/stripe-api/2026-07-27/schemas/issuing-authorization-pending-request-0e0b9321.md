---
title: issuing_authorization_pending_request
page_id: schema-issuing-authorization-pending-request-0e0b9321
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_pending_request

```yaml
{"title": "IssuingAuthorizationPendingRequest", "required": ["amount", "currency", "is_amount_controllable", "merchant_amount", "merchant_currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The additional amount Stripe will hold if the authorization is approved, in the card's [currency](https://docs.stripe.com/api#issuing_authorization_object-pending-request-currency) and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}, "amount_details": {"description": "Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/issuing_authorization_amount_details"}]}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "is_amount_controllable": {"type": "boolean", "description": "If set `true`, you may provide [amount](https://docs.stripe.com/api/issuing/authorizations/approve#approve_issuing_authorization-amount) to control how much to hold for the authorization."}, "merchant_amount": {"type": "integer", "description": "The amount the merchant is requesting to be authorized in the `merchant_currency`. The amount is in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}, "merchant_currency": {"type": "string", "description": "The local currency the merchant is requesting to authorize.", "format": "currency"}, "network_risk_score": {"type": "integer", "description": "The card network's estimate of the likelihood that an authorization is fraudulent. Takes on values between 1 and 99.", "nullable": true}}, "description": "", "x-expandableFields": ["amount_details"]}
```
