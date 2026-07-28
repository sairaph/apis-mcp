---
title: insights_resources_payment_evaluation_money_movement_card
page_id: schema-insights-resources-payment-evaluation-money-movement-card-f6cc64c3
path: schemas
description: Money Movement card details attached to this payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_money_movement_card

Money Movement card details attached to this payment.

```yaml
{"title": "InsightsResourcesPaymentEvaluationMoneyMovementCard", "type": "object", "properties": {"customer_presence": {"type": "string", "description": "Describes the presence of the customer during the payment.", "nullable": true, "enum": ["off_session", "on_session"]}, "payment_type": {"type": "string", "description": "Describes the type of payment.", "nullable": true, "enum": ["one_off", "recurring", "setup_one_off", "setup_recurring"]}}, "description": "Money Movement card details attached to this payment.", "x-expandableFields": []}
```
