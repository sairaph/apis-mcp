---
title: issuing_dispute_treasury
page_id: schema-issuing-dispute-treasury-d4b5e750
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_dispute_treasury

```yaml
{"title": "IssuingDisputeTreasury", "required": ["received_debit"], "type": "object", "properties": {"debit_reversal": {"maxLength": 5000, "type": "string", "description": "The Treasury [DebitReversal](https://docs.stripe.com/api/treasury/debit_reversals) representing this Issuing dispute", "nullable": true}, "received_debit": {"maxLength": 5000, "type": "string", "description": "The Treasury [ReceivedDebit](https://docs.stripe.com/api/treasury/received_debits) that is being disputed."}}, "description": "", "x-expandableFields": []}
```
