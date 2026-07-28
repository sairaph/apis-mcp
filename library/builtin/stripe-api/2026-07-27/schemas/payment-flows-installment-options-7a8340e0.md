---
title: payment_flows_installment_options
page_id: schema-payment-flows-installment-options-7a8340e0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_installment_options

```yaml
{"title": "PaymentFlowsInstallmentOptions", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}, "plan": {"$ref": "#/components/schemas/payment_method_details_card_installments_plan"}}, "description": "", "x-expandableFields": ["plan"]}
```
