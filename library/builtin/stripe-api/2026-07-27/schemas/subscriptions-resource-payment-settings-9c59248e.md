---
title: subscriptions_resource_payment_settings
page_id: schema-subscriptions-resource-payment-settings-9c59248e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_payment_settings

```yaml
{"title": "SubscriptionsResourcePaymentSettings", "type": "object", "properties": {"payment_method_options": {"description": "Payment-method-specific configuration to provide to invoices created by the subscription.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/subscriptions_resource_payment_method_options"}]}, "payment_method_types": {"type": "array", "description": "The list of payment method types to provide to every invoice created by the subscription. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).", "nullable": true, "items": {"type": "string", "enum": ["ach_credit_transfer", "ach_debit", "acss_debit", "affirm", "amazon_pay", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "cashapp", "crypto", "custom", "customer_balance", "eps", "fpx", "giropay", "grabpay", "ideal", "jp_credit_transfer", "kakao_pay", "klarna", "konbini", "kr_card", "link", "multibanco", "naver_pay", "nz_bank_account", "p24", "pay_by_bank", "payco", "paynow", "paypal", "payto", "pix", "promptpay", "revolut_pay", "satispay", "sepa_credit_transfer", "sepa_debit", "sofort", "swish", "twint", "upi", "us_bank_account", "wechat_pay"], "x-stripeBypassValidation": true}}, "save_default_payment_method": {"type": "string", "description": "Configure whether Stripe updates `subscription.default_payment_method` when payment succeeds. Defaults to `off`.", "nullable": true, "enum": ["off", "on_subscription"]}}, "description": "", "x-expandableFields": ["payment_method_options"]}
```
