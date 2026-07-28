---
title: payment_method_sepa_debit
page_id: schema-payment-method-sepa-debit-8aff4cf2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_sepa_debit

```yaml
{"title": "payment_method_sepa_debit", "type": "object", "properties": {"bank_code": {"maxLength": 5000, "type": "string", "description": "Bank code of bank associated with the bank account.", "nullable": true}, "branch_code": {"maxLength": 5000, "type": "string", "description": "Branch code of bank associated with the bank account.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country the bank account is located in.", "nullable": true}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "generated_from": {"description": "Information about the object that generated this PaymentMethod.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/sepa_debit_generated_from"}]}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four characters of the IBAN.", "nullable": true}}, "description": "", "x-expandableFields": ["generated_from"]}
```
