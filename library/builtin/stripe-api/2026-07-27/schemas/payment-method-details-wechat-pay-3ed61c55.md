---
title: payment_method_details_wechat_pay
page_id: schema-payment-method-details-wechat-pay-3ed61c55
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_wechat_pay

```yaml
{"title": "payment_method_details_wechat_pay", "type": "object", "properties": {"fingerprint": {"maxLength": 5000, "type": "string", "description": "Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.", "nullable": true}, "location": {"maxLength": 5000, "type": "string", "description": "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to."}, "reader": {"maxLength": 5000, "type": "string", "description": "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on."}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "Transaction ID of this particular WeChat Pay transaction.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
