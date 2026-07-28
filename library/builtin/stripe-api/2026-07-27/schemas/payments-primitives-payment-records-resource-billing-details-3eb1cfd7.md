---
title: payments_primitives_payment_records_resource_billing_details
page_id: schema-payments-primitives-payment-records-resource-billing-details-3eb1cfd7
path: schemas
description: Billing details used by the customer for this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_billing_details

Billing details used by the customer for this payment.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceBillingDetails", "required": ["address"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/payments_primitives_payment_records_resource_address"}, "email": {"maxLength": 5000, "type": "string", "description": "The billing email associated with the method of payment.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "The billing name associated with the method of payment.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "The billing phone number associated with the method of payment.", "nullable": true}}, "description": "Billing details used by the customer for this payment.", "x-expandableFields": ["address"]}
```
