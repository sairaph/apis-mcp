---
title: issuing_transaction_receipt_data
page_id: schema-issuing-transaction-receipt-data-56f67217
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_receipt_data

```yaml
{"title": "IssuingTransactionReceiptData", "type": "object", "properties": {"description": {"maxLength": 5000, "type": "string", "description": "The description of the item. The maximum length of this field is 26 characters.", "nullable": true}, "quantity": {"type": "number", "description": "The quantity of the item.", "nullable": true}, "total": {"type": "integer", "description": "The total for this line item in cents.", "nullable": true}, "unit_cost": {"type": "integer", "description": "The unit cost of the item in cents.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
