---
title: payment_method_details_payment_record_payto
page_id: schema-payment-method-details-payment-record-payto-2ea74697
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_payment_record_payto

```yaml
{"title": "payment_method_details_payment_record_payto", "type": "object", "properties": {"bsb_number": {"maxLength": 5000, "type": "string", "description": "Bank-State-Branch number of the bank account.", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "Last four digits of the bank account number.", "nullable": true}, "mandate": {"maxLength": 5000, "type": "string", "description": "ID of the mandate used to make this payment."}, "pay_id": {"maxLength": 5000, "type": "string", "description": "The PayID alias for the bank account.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
