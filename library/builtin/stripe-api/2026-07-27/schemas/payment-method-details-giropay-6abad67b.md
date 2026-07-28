---
title: payment_method_details_giropay
page_id: schema-payment-method-details-giropay-6abad67b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_giropay

```yaml
{"title": "payment_method_details_giropay", "type": "object", "properties": {"bank_code": {"maxLength": 5000, "type": "string", "description": "Bank code of bank associated with the bank account.", "nullable": true}, "bank_name": {"maxLength": 5000, "type": "string", "description": "Name of the bank associated with the bank account.", "nullable": true}, "bic": {"maxLength": 5000, "type": "string", "description": "Bank Identifier Code of the bank associated with the bank account.", "nullable": true}, "verified_name": {"maxLength": 5000, "type": "string", "description": "Owner's verified full name. Values are verified or provided by Giropay directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.\nGiropay rarely provides this information so the attribute is usually empty.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
