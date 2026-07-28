---
title: insights_resources_payment_evaluation_customer_details
page_id: schema-insights-resources-payment-evaluation-customer-details-ad80ded8
path: schemas
description: Customer details attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_customer_details

Customer details attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationCustomerDetails", "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string", "description": "The ID of the customer associated with the payment evaluation.", "nullable": true}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The ID of the Account representing the customer associated with the payment evaluation.", "nullable": true}, "email": {"maxLength": 5000, "type": "string", "description": "The customer's email address.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "The customer's full name or business name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "The customer's phone number.", "nullable": true}}, "description": "Customer details attached to this payment evaluation.", "x-expandableFields": []}
```
