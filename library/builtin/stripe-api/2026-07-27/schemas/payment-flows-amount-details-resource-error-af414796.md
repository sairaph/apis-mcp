---
title: payment_flows_amount_details_resource_error
page_id: schema-payment-flows-amount-details-resource-error-af414796
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_amount_details_resource_error

```yaml
{"title": "PaymentFlowsAmountDetailsResourceError", "type": "object", "properties": {"code": {"type": "string", "description": "The code of the error that occurred when validating the current amount details.", "nullable": true, "enum": ["amount_details_amount_mismatch", "amount_details_tax_shipping_discount_greater_than_amount"]}, "message": {"maxLength": 5000, "type": "string", "description": "A message providing more details about the error.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
