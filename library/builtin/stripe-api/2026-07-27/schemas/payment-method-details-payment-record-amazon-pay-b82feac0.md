---
title: payment_method_details_payment_record_amazon_pay
page_id: schema-payment-method-details-payment-record-amazon-pay-b82feac0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_payment_record_amazon_pay

```yaml
{"title": "payment_method_details_payment_record_amazon_pay", "type": "object", "properties": {"funding": {"$ref": "#/components/schemas/payments_primitives_payment_records_resource_payment_method_amazon_pay_details_resource_funding"}, "transaction_id": {"maxLength": 5000, "type": "string", "description": "The Amazon Pay transaction ID associated with this payment.", "nullable": true}}, "description": "", "x-expandableFields": ["funding"]}
```
