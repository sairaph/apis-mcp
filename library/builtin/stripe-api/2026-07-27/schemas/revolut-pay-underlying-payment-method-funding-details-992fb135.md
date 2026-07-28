---
title: revolut_pay_underlying_payment_method_funding_details
page_id: schema-revolut-pay-underlying-payment-method-funding-details-992fb135
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# revolut_pay_underlying_payment_method_funding_details

```yaml
{"title": "revolut_pay_underlying_payment_method_funding_details", "type": "object", "properties": {"card": {"$ref": "#/components/schemas/payment_method_details_passthrough_card"}, "type": {"type": "string", "description": "funding type of the underlying payment method.", "nullable": true, "enum": ["card"]}}, "description": "", "x-expandableFields": ["card"]}
```
