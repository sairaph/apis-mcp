---
title: invoice_payment_method_options_customer_balance_bank_transfer
page_id: schema-invoice-payment-method-options-customer-balance-bank-transfer-7341314e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment_method_options_customer_balance_bank_transfer

```yaml
{"title": "invoice_payment_method_options_customer_balance_bank_transfer", "type": "object", "properties": {"eu_bank_transfer": {"$ref": "#/components/schemas/invoice_payment_method_options_customer_balance_bank_transfer_eu_bank_transfer"}, "type": {"type": "string", "description": "The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.", "nullable": true}}, "description": "", "x-expandableFields": ["eu_bank_transfer"]}
```
