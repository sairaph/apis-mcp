---
title: invoice_payment_method_options_customer_balance
page_id: schema-invoice-payment-method-options-customer-balance-f0362630
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment_method_options_customer_balance

```yaml
{"title": "invoice_payment_method_options_customer_balance", "type": "object", "properties": {"bank_transfer": {"$ref": "#/components/schemas/invoice_payment_method_options_customer_balance_bank_transfer"}, "funding_type": {"type": "string", "description": "The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.", "nullable": true, "enum": ["bank_transfer"]}}, "description": "", "x-expandableFields": ["bank_transfer"]}
```
