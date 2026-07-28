---
title: payment_method_options_konbini
page_id: schema-payment-method-options-konbini-e8571cba
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_options_konbini

```yaml
{"title": "payment_method_options_konbini", "type": "object", "properties": {"confirmation_number": {"maxLength": 5000, "type": "string", "description": "An optional 10 to 11 digit numeric-only string determining the confirmation code at applicable convenience stores.", "nullable": true}, "expires_after_days": {"type": "integer", "description": "The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST.", "nullable": true}, "expires_at": {"type": "integer", "description": "The timestamp at which the Konbini payment instructions will expire. Only one of `expires_after_days` or `expires_at` may be set.", "format": "unix-time", "nullable": true}, "product_description": {"maxLength": 5000, "type": "string", "description": "A product descriptor of up to 22 characters, which will appear to customers at the convenience store.", "nullable": true}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).", "enum": ["none"]}}, "description": "", "x-expandableFields": []}
```
