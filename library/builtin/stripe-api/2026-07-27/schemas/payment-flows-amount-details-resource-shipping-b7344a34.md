---
title: payment_flows_amount_details_resource_shipping
page_id: schema-payment-flows-amount-details-resource-shipping-b7344a34
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_amount_details_resource_shipping

```yaml
{"title": "PaymentFlowsAmountDetailsResourceShipping", "type": "object", "properties": {"amount": {"type": "integer", "description": "If a physical good is being shipped, the cost of shipping represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). An integer greater than or equal to 0.", "nullable": true}, "from_postal_code": {"maxLength": 5000, "type": "string", "description": "If a physical good is being shipped, the postal code of where it is being shipped from. At most 10 alphanumeric characters long, hyphens and spaces are allowed.", "nullable": true}, "to_postal_code": {"maxLength": 5000, "type": "string", "description": "If a physical good is being shipped, the postal code of where it is being shipped to. At most 10 alphanumeric characters long, hyphens and spaces are allowed.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
