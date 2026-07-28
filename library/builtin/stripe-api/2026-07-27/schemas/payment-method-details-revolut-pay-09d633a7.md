---
title: payment_method_details_revolut_pay
page_id: schema-payment-method-details-revolut-pay-09d633a7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_revolut_pay

```yaml
{"title": "payment_method_details_revolut_pay", "type": "object", "properties": {"funding": {"$ref": "#/components/schemas/revolut_pay_underlying_payment_method_funding_details"}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "The Revolut Pay transaction ID associated with this payment.", "nullable": true}}, "description": "", "x-expandableFields": ["funding"]}
```
