---
title: payment_method_au_becs_debit
page_id: schema-payment-method-au-becs-debit-b85b10ad
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_au_becs_debit

```yaml
{"title": "payment_method_au_becs_debit", "type": "object", "properties": {"bsb_number": {"maxLength": 5000, "type": "string", "description": "Six-digit number identifying bank and branch associated with this bank account.", "nullable": true}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
