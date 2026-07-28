---
title: payment_flows_amount_details_resource_tax
page_id: schema-payment-flows-amount-details-resource-tax-6efac30f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_amount_details_resource_tax

```yaml
{"title": "PaymentFlowsAmountDetailsResourceTax", "type": "object", "properties": {"total_tax_amount": {"type": "integer", "description": "The total amount of tax on the transaction represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Required for L2 rates. An integer greater than or equal to 0.\n\nThis field is mutually exclusive with the `amount_details[line_items][#][tax][total_tax_amount]` field.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
