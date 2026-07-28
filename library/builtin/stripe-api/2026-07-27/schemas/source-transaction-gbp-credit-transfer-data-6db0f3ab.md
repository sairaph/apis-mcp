---
title: source_transaction_gbp_credit_transfer_data
page_id: schema-source-transaction-gbp-credit-transfer-data-6db0f3ab
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_transaction_gbp_credit_transfer_data

```yaml
{"title": "SourceTransactionGbpCreditTransferData", "type": "object", "properties": {"fingerprint": {"maxLength": 5000, "type": "string", "description": "Bank account fingerprint associated with the Stripe owned bank account receiving the transfer."}, "funding_method": {"maxLength": 5000, "type": "string", "description": "The credit transfer rails the sender used to push this transfer. The possible rails are: Faster Payments, BACS, CHAPS, and wire transfers. Currently only Faster Payments is supported."}, "last4": {"maxLength": 5000, "type": "string", "description": "Last 4 digits of sender account number associated with the transfer."}, "reference": {"maxLength": 5000, "type": "string", "description": "Sender entered arbitrary information about the transfer."}, "sender_account_number": {"maxLength": 5000, "type": "string", "description": "Sender account number associated with the transfer."}, "sender_name": {"maxLength": 5000, "type": "string", "description": "Sender name associated with the transfer."}, "sender_sort_code": {"maxLength": 5000, "type": "string", "description": "Sender sort code associated with the transfer."}}, "description": "", "x-expandableFields": []}
```
