---
title: payment_source
page_id: schema-payment-source-6644104f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_source

```yaml
{"title": "Polymorphic", "anyOf": [{"$ref": "#/components/schemas/account"}, {"$ref": "#/components/schemas/bank_account"}, {"$ref": "#/components/schemas/card"}, {"$ref": "#/components/schemas/source"}], "x-resourceId": "payment_source", "x-stripeBypassValidation": true}
```
