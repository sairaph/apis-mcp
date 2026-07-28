---
title: payment_method_details_ach_credit_transfer
page_id: schema-payment-method-details-ach-credit-transfer-a4b85541
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_ach_credit_transfer

```yaml
{"title": "payment_method_details_ach_credit_transfer", "type": "object", "properties": {"account_number": {"maxLength": 5000, "type": "string", "description": "Account number to transfer funds to.", "nullable": true}, "bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the routing number.", "nullable": true}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing transit number for the bank account to transfer funds to.", "nullable": true}, "swift_code": {"maxLength": 5000, "type": "string", "description": "SWIFT code of the bank associated with the routing number.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
