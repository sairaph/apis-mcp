---
title: checkout_wechat_pay_payment_method_options
page_id: schema-checkout-wechat-pay-payment-method-options-04a839b2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# checkout_wechat_pay_payment_method_options

```yaml
{"title": "CheckoutWechatPayPaymentMethodOptions", "type": "object", "properties": {"app_id": {"maxLength": 5000, "type": "string", "description": "The app ID registered with WeChat Pay. Only required when client is iOS or Android.", "nullable": true}, "client": {"type": "string", "description": "The client type that the end customer will pay from", "nullable": true, "enum": ["android", "ios", "web"], "x-stripeBypassValidation": true}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).", "enum": ["none"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
