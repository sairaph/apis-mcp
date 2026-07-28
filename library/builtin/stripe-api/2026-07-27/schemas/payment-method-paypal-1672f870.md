---
title: payment_method_paypal
page_id: schema-payment-method-paypal-1672f870
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_paypal

```yaml
{"title": "payment_method_paypal", "type": "object", "properties": {"country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the buyer's country. Values are provided by PayPal directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}, "payer_email": {"maxLength": 5000, "type": "string", "description": "Owner's email. Values are provided by PayPal directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}, "payer_id": {"maxLength": 5000, "type": "string", "description": "PayPal account PayerID. This identifier uniquely identifies the PayPal customer.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
