---
title: issuing_transaction_treasury
page_id: schema-issuing-transaction-treasury-5315d9be
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_transaction_treasury

```yaml
{"title": "IssuingTransactionTreasury", "type": "object", "properties": {"received_credit": {"maxLength": 5000, "type": "string", "description": "The Treasury [ReceivedCredit](https://docs.stripe.com/api/treasury/received_credits) representing this Issuing transaction if it is a refund", "nullable": true}, "received_debit": {"maxLength": 5000, "type": "string", "description": "The Treasury [ReceivedDebit](https://docs.stripe.com/api/treasury/received_debits) representing this Issuing transaction if it is a capture", "nullable": true}}, "description": "", "x-expandableFields": []}
```
