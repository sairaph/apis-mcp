---
title: payment_method_options_afterpay_clearpay
page_id: schema-payment-method-options-afterpay-clearpay-a95510b3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_options_afterpay_clearpay

```yaml
{"title": "payment_method_options_afterpay_clearpay", "type": "object", "properties": {"capture_method": {"type": "string", "description": "Controls when the funds will be captured from the customer's account.", "enum": ["manual"]}, "reference": {"maxLength": 5000, "type": "string", "description": "An internal identifier or reference that this payment corresponds to. You must limit the identifier to 128 characters, and it can only contain letters, numbers, underscores, backslashes, and dashes.\nThis field differs from the statement descriptor and item name.", "nullable": true}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).", "enum": ["none"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
