---
title: payments_primitives_payment_records_resource_processor_details_resource_custom_details
page_id: schema-payments-primitives-payment-records-resource-processor-details-resource-379f5106
path: schemas
description: |-
    Custom processors represent payment processors not modeled directly in
    the Stripe API. This resource consists of details about the custom processor
    used for this payment attempt.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_processor_details_resource_custom_details

Custom processors represent payment processors not modeled directly in
the Stripe API. This resource consists of details about the custom processor
used for this payment attempt.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceProcessorDetailsResourceCustomDetails", "type": "object", "properties": {"payment_reference": {"maxLength": 5000, "type": "string", "description": "An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.", "nullable": true}}, "description": "Custom processors represent payment processors not modeled directly in\nthe Stripe API. This resource consists of details about the custom processor\nused for this payment attempt.", "x-expandableFields": []}
```
