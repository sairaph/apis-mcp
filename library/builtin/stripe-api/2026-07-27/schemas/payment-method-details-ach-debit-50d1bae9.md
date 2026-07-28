---
title: payment_method_details_ach_debit
page_id: schema-payment-method-details-ach-debit-50d1bae9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_ach_debit

```yaml
{"title": "payment_method_details_ach_debit", "type": "object", "properties": {"account_holder_type": {"type": "string", "description": "Type of entity that holds the account. This can be either `individual` or `company`.", "nullable": true, "enum": ["company", "individual"]}, "bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the bank account.", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country the bank account is located in.", "nullable": true}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing transit number of the bank account.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
