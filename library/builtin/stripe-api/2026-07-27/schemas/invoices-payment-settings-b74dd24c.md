---
title: invoices_payment_settings
page_id: schema-invoices-payment-settings-b74dd24c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_payment_settings

```yaml
{"title": "InvoicesPaymentSettings", "type": "object", "properties": {"default_mandate": {"maxLength": 5000, "type": "string", "description": "ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the invoice's default_payment_method or default_source, if set.", "nullable": true}, "payment_method_options": {"description": "Payment-method-specific configuration to provide to the invoice’s PaymentIntent.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/invoices_payment_method_options"}]}, "payment_method_types": {"type": "array", "description": "The list of payment method types (e.g. card) to provide to the invoice’s PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).", "nullable": true, "items": {"type": "string", "enum": ["ach_credit_transfer", "ach_debit", "acss_debit", "affirm", "amazon_pay", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "cashapp", "crypto", "custom", "customer_balance", "eps", "fpx", "giropay", "grabpay", "ideal", "jp_credit_transfer", "kakao_pay", "klarna", "konbini", "kr_card", "link", "multibanco", "naver_pay", "nz_bank_account", "p24", "pay_by_bank", "payco", "paynow", "paypal", "payto", "pix", "promptpay", "revolut_pay", "satispay", "sepa_credit_transfer", "sepa_debit", "sofort", "swish", "twint", "upi", "us_bank_account", "wechat_pay"], "x-stripeBypassValidation": true}}}, "description": "", "x-expandableFields": ["payment_method_options"]}
```
