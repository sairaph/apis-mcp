---
title: payment_method_acss_debit
page_id: schema-payment-method-acss-debit-4c6d138d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_acss_debit

```yaml
{"title": "payment_method_acss_debit", "type": "object", "properties": {"bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the bank account.", "nullable": true}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "institution_number": {"maxLength": 5000, "type": "string", "description": "Institution number of the bank account.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "transit_number": {"maxLength": 5000, "type": "string", "description": "Transit number of the bank account.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
