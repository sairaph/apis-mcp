---
title: payment_method_naver_pay
page_id: schema-payment-method-naver-pay-dccfc4d2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_naver_pay

```yaml
{"title": "payment_method_naver_pay", "required": ["funding"], "type": "object", "properties": {"buyer_id": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular Naver Pay account. You can use this attribute to check whether two Naver Pay accounts are the same.", "nullable": true}, "funding": {"type": "string", "description": "Whether to fund this transaction with Naver Pay points or a card.", "enum": ["card", "points"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
