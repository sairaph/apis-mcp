---
title: issuing_authorization_fraud_challenge
page_id: schema-issuing-authorization-fraud-challenge-6a821147
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_authorization_fraud_challenge

```yaml
{"title": "IssuingAuthorizationFraudChallenge", "required": ["channel", "status"], "type": "object", "properties": {"channel": {"type": "string", "description": "The method by which the fraud challenge was delivered to the cardholder.", "enum": ["sms"]}, "status": {"type": "string", "description": "The status of the fraud challenge.", "enum": ["expired", "pending", "rejected", "undeliverable", "verified"]}, "undeliverable_reason": {"type": "string", "description": "If the challenge is not deliverable, the reason why.", "nullable": true, "enum": ["no_phone_number", "unsupported_phone_number"]}}, "description": "", "x-expandableFields": []}
```
