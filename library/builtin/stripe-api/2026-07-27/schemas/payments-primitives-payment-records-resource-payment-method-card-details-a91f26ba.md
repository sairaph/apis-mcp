---
title: payments_primitives_payment_records_resource_payment_method_card_details_resource_installment_plan
page_id: schema-payments-primitives-payment-records-resource-payment-method-card-details-a91f26ba
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_payment_method_card_details_resource_installment_plan

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCardDetailsResourceInstallmentPlan", "required": ["type"], "type": "object", "properties": {"count": {"type": "integer", "description": "For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.", "nullable": true}, "interval": {"type": "string", "description": "For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.", "nullable": true, "enum": ["month"]}, "type": {"type": "string", "description": "Type of installment plan, one of `fixed_count`, `revolving`, or `bonus`.", "enum": ["bonus", "fixed_count", "revolving"]}}, "description": "", "x-expandableFields": []}
```
