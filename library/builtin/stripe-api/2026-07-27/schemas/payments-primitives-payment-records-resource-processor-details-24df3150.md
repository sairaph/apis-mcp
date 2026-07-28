---
title: payments_primitives_payment_records_resource_processor_details
page_id: schema-payments-primitives-payment-records-resource-processor-details-24df3150
path: schemas
description: Processor information associated with this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_processor_details

Processor information associated with this payment.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceProcessorDetails", "required": ["type"], "type": "object", "properties": {"custom": {"$ref": "#/components/schemas/payments_primitives_payment_records_resource_processor_details_resource_custom_details"}, "type": {"type": "string", "description": "The processor used for this payment attempt.", "enum": ["custom"], "x-stripeBypassValidation": true}}, "description": "Processor information associated with this payment.", "x-expandableFields": ["custom"]}
```
