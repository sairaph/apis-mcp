---
title: payments_primitives_payment_records_resource_payment_method_custom_details
page_id: schema-payments-primitives-payment-records-resource-payment-method-custom-detai-ed8e48b9
path: schemas
description: |-
    Custom Payment Methods represent Payment Method types not modeled directly in
    the Stripe API. This resource consists of details about the custom payment method
    used for this payment attempt.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_payment_method_custom_details

Custom Payment Methods represent Payment Method types not modeled directly in
the Stripe API. This resource consists of details about the custom payment method
used for this payment attempt.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCustomDetails", "required": ["display_name"], "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string", "description": "Display name for the custom (user-defined) payment method type used to make this payment."}, "type": {"maxLength": 5000, "type": "string", "description": "The custom payment method type associated with this payment.", "nullable": true}}, "description": "Custom Payment Methods represent Payment Method types not modeled directly in\nthe Stripe API. This resource consists of details about the custom payment method\nused for this payment attempt.", "x-expandableFields": []}
```
