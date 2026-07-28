---
title: payments_primitives_payment_records_resource_shipping_details
page_id: schema-payments-primitives-payment-records-resource-shipping-details-fcc4260e
path: schemas
description: The customer's shipping information associated with this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_shipping_details

The customer's shipping information associated with this payment.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceShippingDetails", "required": ["address"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/payments_primitives_payment_records_resource_address"}, "name": {"maxLength": 5000, "type": "string", "description": "The shipping recipient's name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "The shipping recipient's phone number.", "nullable": true}}, "description": "The customer's shipping information associated with this payment.", "x-expandableFields": ["address"]}
```
