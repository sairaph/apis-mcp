---
title: issuing_transaction_network_data
page_id: schema-issuing-transaction-network-data-50d378d1
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_network_data

```yaml
{"title": "IssuingTransactionNetworkData", "type": "object", "properties": {"authorization_code": {"maxLength": 5000, "type": "string", "description": "A code created by Stripe which is shared with the merchant to validate the authorization. This field will be populated if the authorization message was approved. The code typically starts with the letter \"S\", followed by a six-digit number. For example, \"S498162\". Please note that the code is not guaranteed to be unique across authorizations.", "nullable": true}, "processing_date": {"maxLength": 5000, "type": "string", "description": "The date the transaction was processed by the card network. This can be different from the date the seller recorded the transaction depending on when the acquirer submits the transaction to the network.", "nullable": true}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the authorization assigned by the card network used to match subsequent messages, disputes, and transactions.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
