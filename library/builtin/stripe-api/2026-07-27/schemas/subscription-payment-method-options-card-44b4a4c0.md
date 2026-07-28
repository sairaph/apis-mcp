---
title: subscription_payment_method_options_card
page_id: schema-subscription-payment-method-options-card-44b4a4c0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_payment_method_options_card

```yaml
{"title": "subscription_payment_method_options_card", "type": "object", "properties": {"mandate_options": {"$ref": "#/components/schemas/invoice_mandate_options_card"}, "network": {"type": "string", "description": "Selected network to process this Subscription on. Depends on the available networks of the card attached to the Subscription. Can be only set confirm-time.", "nullable": true, "enum": ["amex", "cartes_bancaires", "diners", "discover", "eftpos_au", "girocard", "interac", "jcb", "link", "mastercard", "unionpay", "unknown", "visa"]}, "request_three_d_secure": {"type": "string", "description": "We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.", "nullable": true, "enum": ["any", "automatic", "challenge"]}}, "description": "", "x-expandableFields": ["mandate_options"]}
```
