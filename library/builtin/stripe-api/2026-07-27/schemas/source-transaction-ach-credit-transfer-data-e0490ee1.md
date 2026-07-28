---
title: source_transaction_ach_credit_transfer_data
page_id: schema-source-transaction-ach-credit-transfer-data-e0490ee1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_transaction_ach_credit_transfer_data

```yaml
{"title": "SourceTransactionAchCreditTransferData", "type": "object", "properties": {"customer_data": {"maxLength": 5000, "type": "string", "description": "Customer data associated with the transfer."}, "fingerprint": {"maxLength": 5000, "type": "string", "description": "Bank account fingerprint associated with the transfer."}, "last4": {"maxLength": 5000, "type": "string", "description": "Last 4 digits of the account number associated with the transfer."}, "routing_number": {"maxLength": 5000, "type": "string", "description": "Routing number associated with the transfer."}}, "description": "", "x-expandableFields": []}
```
