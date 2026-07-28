---
title: payment_flows_payment_details
page_id: schema-payment-flows-payment-details-b2444151
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_payment_details

```yaml
{"title": "PaymentFlowsPaymentDetails", "type": "object", "properties": {"customer_reference": {"maxLength": 5000, "type": "string", "description": "A unique value to identify the customer. This field is available only for card payments.\n\nThis field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.", "nullable": true}, "order_reference": {"maxLength": 5000, "type": "string", "description": "A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.\n\nFor Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks. For Klarna, this field is truncated to 255 characters and is visible to customers when they view the order in the Klarna app.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
