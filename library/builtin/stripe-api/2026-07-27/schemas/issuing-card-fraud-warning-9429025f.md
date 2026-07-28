---
title: issuing_card_fraud_warning
page_id: schema-issuing-card-fraud-warning-9429025f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_card_fraud_warning

```yaml
{"title": "IssuingCardFraudWarning", "type": "object", "properties": {"started_at": {"type": "integer", "description": "Timestamp of the most recent fraud warning.", "format": "unix-time", "nullable": true}, "type": {"type": "string", "description": "The type of fraud warning that most recently took place on this card. This field updates with every new fraud warning, so the value changes over time. If populated, cancel and reissue the card.", "nullable": true, "enum": ["card_testing_exposure", "fraud_dispute_filed", "third_party_reported", "user_indicated_fraud"]}}, "description": "", "x-expandableFields": []}
```
