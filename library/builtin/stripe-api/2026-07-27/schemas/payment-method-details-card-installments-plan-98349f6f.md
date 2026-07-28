---
title: payment_method_details_card_installments_plan
page_id: schema-payment-method-details-card-installments-plan-98349f6f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_card_installments_plan

```yaml
{"title": "payment_method_details_card_installments_plan", "required": ["type"], "type": "object", "properties": {"count": {"type": "integer", "description": "For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.", "nullable": true}, "interval": {"type": "string", "description": "For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card.\nOne of `month`.", "nullable": true, "enum": ["month"]}, "type": {"type": "string", "description": "Type of installment plan, one of `fixed_count`, `bonus`, or `revolving`.", "enum": ["bonus", "fixed_count", "revolving"]}}, "description": "", "x-expandableFields": []}
```
