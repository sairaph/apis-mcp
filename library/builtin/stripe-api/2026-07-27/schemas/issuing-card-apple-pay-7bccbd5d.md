---
title: issuing_card_apple_pay
page_id: schema-issuing-card-apple-pay-7bccbd5d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_card_apple_pay

```yaml
{"title": "IssuingCardApplePay", "required": ["eligible"], "type": "object", "properties": {"eligible": {"type": "boolean", "description": "Apple Pay Eligibility"}, "ineligible_reason": {"type": "string", "description": "Reason the card is ineligible for Apple Pay", "nullable": true, "enum": ["missing_agreement", "missing_cardholder_contact", "unsupported_region"]}}, "description": "", "x-expandableFields": []}
```
