---
title: payment_method_kr_card
page_id: schema-payment-method-kr-card-8945c866
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_kr_card

```yaml
{"title": "payment_method_kr_card", "type": "object", "properties": {"brand": {"type": "string", "description": "The local credit or debit card brand.", "nullable": true, "enum": ["bc", "citi", "hana", "hyundai", "jeju", "jeonbuk", "kakaobank", "kbank", "kdbbank", "kookmin", "kwangju", "lotte", "mg", "nh", "post", "samsung", "savingsbank", "shinhan", "shinhyup", "suhyup", "tossbank", "woori"]}, "last4": {"maxLength": 4, "type": "string", "description": "The last four digits of the card. This may not be present for American Express cards.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
