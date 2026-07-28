---
title: payment_method_options_card_installments
page_id: schema-payment-method-options-card-installments-57f2ea42
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_options_card_installments

```yaml
{"title": "payment_method_options_card_installments", "required": ["enabled"], "type": "object", "properties": {"available_plans": {"type": "array", "description": "Installment plans that may be selected for this PaymentIntent.", "nullable": true, "items": {"$ref": "#/components/schemas/payment_method_details_card_installments_plan"}}, "enabled": {"type": "boolean", "description": "Whether Installments are enabled for this PaymentIntent."}, "plan": {"description": "Installment plan selected for this PaymentIntent.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_method_details_card_installments_plan"}]}}, "description": "", "x-expandableFields": ["available_plans", "plan"]}
```
