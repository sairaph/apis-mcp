---
title: payment_method_details_kr_card
page_id: schema-payment-method-details-kr-card-88e5f242
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_kr_card

```yaml
{"title": "payment_method_details_kr_card", "type": "object", "properties": {"brand": {"type": "string", "description": "The local credit or debit card brand.", "nullable": true, "enum": ["bc", "citi", "hana", "hyundai", "jeju", "jeonbuk", "kakaobank", "kbank", "kdbbank", "kookmin", "kwangju", "lotte", "mg", "nh", "post", "samsung", "savingsbank", "shinhan", "shinhyup", "suhyup", "tossbank", "woori"]}, "buyer_id": {"maxLength": 5000, "type": "string", "description": "A unique identifier for the buyer as determined by the local payment processor.", "nullable": true}, "last4": {"maxLength": 4, "type": "string", "description": "The last four digits of the card. This may not be present for American Express cards.", "nullable": true}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "The Korean Card transaction ID associated with this payment.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
