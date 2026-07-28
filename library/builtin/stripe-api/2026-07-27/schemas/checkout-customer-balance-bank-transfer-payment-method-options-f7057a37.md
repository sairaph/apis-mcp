---
title: checkout_customer_balance_bank_transfer_payment_method_options
page_id: schema-checkout-customer-balance-bank-transfer-payment-method-options-f7057a37
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# checkout_customer_balance_bank_transfer_payment_method_options

```yaml
{"title": "CheckoutCustomerBalanceBankTransferPaymentMethodOptions", "type": "object", "properties": {"eu_bank_transfer": {"$ref": "#/components/schemas/payment_method_options_customer_balance_eu_bank_account"}, "requested_address_types": {"type": "array", "description": "List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.\n\nPermitted values include: `sort_code`, `zengin`, `iban`, or `spei`.", "items": {"type": "string", "enum": ["aba", "iban", "sepa", "sort_code", "spei", "swift", "zengin"], "x-stripeBypassValidation": true}}, "type": {"type": "string", "description": "The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.", "nullable": true, "enum": ["eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["eu_bank_transfer"]}
```
