---
title: payment_method_payto
page_id: schema-payment-method-payto-16224abf
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_payto

```yaml
{"title": "payment_method_payto", "type": "object", "properties": {"bsb_number": {"maxLength": 5000, "type": "string", "description": "Bank-State-Branch number of the bank account.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "pay_id": {"maxLength": 5000, "type": "string", "description": "The PayID alias for the bank account.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
