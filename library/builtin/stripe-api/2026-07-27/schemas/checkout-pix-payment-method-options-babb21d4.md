---
title: checkout_pix_payment_method_options
page_id: schema-checkout-pix-payment-method-options-babb21d4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# checkout_pix_payment_method_options

```yaml
{"title": "CheckoutPixPaymentMethodOptions", "type": "object", "properties": {"amount_includes_iof": {"type": "string", "description": "Determines if the amount includes the IOF tax.", "enum": ["always", "never"]}, "expires_after_seconds": {"type": "integer", "description": "The number of seconds after which Pix payment will expire.", "nullable": true}, "mandate_options": {"$ref": "#/components/schemas/payment_method_options_mandate_options_pix"}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).", "enum": ["none", "off_session"]}}, "description": "", "x-expandableFields": ["mandate_options"]}
```
