---
title: payments_primitives_payment_records_resource_payment_method_mobilepay_details_resource_card
page_id: schema-payments-primitives-payment-records-resource-payment-method-mobilepay-de-6e56f9ad
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_payment_method_mobilepay_details_resource_card

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourcePaymentMethodMobilepayDetailsResourceCard", "type": "object", "properties": {"brand": {"maxLength": 5000, "type": "string", "description": "Brand of the card used in the transaction", "nullable": true}, "country": {"maxLength": 5000, "type": "string", "description": "Two-letter ISO code representing the country of the card", "nullable": true}, "exp_month": {"type": "integer", "description": "Two digit number representing the card's expiration month", "nullable": true}, "exp_year": {"type": "integer", "description": "Two digit number representing the card's expiration year", "nullable": true}, "last4": {"maxLength": 5000, "type": "string", "description": "The last 4 digits of the card", "nullable": true}}, "description": "", "x-expandableFields": []}
```
