---
title: payment_method_details_affirm
page_id: schema-payment-method-details-affirm-cf1077c6
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_affirm

```yaml
{"title": "payment_method_details_affirm", "type": "object", "properties": {"location": {"maxLength": 5000, "type": "string", "description": "ID of the location that this reader is assigned to."}, "reader": {"maxLength": 5000, "type": "string", "description": "ID of the reader this transaction was made on."}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "The Affirm transaction ID associated with this payment.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
