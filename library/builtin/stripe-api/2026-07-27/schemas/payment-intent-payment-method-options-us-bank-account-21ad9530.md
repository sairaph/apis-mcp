---
title: payment_intent_payment_method_options_us_bank_account
page_id: schema-payment-intent-payment-method-options-us-bank-account-21ad9530
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_payment_method_options_us_bank_account

```yaml
{"title": "payment_intent_payment_method_options_us_bank_account", "type": "object", "properties": {"financial_connections": {"$ref": "#/components/schemas/linked_account_options_common"}, "mandate_options": {"$ref": "#/components/schemas/payment_method_options_us_bank_account_mandate_options"}, "setup_future_usage": {"type": "string", "description": "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).", "enum": ["none", "off_session", "on_session"]}, "target_date": {"maxLength": 5000, "type": "string", "description": "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now."}, "transaction_purpose": {"type": "string", "description": "The purpose of the transaction.", "enum": ["goods", "other", "services", "unspecified"]}, "verification_method": {"type": "string", "description": "Bank account verification method. The default value is `automatic`.", "enum": ["automatic", "instant", "microdeposits"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["financial_connections", "mandate_options"]}
```
