---
title: payments_primitives_payment_records_resource_customer_details
page_id: schema-payments-primitives-payment-records-resource-customer-details-93b39403
path: schemas
description: Information about the customer for this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_customer_details

Information about the customer for this payment.

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourceCustomerDetails", "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string", "description": "ID of the Stripe Customer associated with this payment.", "nullable": true}, "email": {"maxLength": 5000, "type": "string", "description": "The customer's email address.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "The customer's name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "The customer's phone number.", "nullable": true}}, "description": "Information about the customer for this payment.", "x-expandableFields": []}
```
