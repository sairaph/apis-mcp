---
title: insights_resources_payment_evaluation_merchant_blocked
page_id: schema-insights-resources-payment-evaluation-merchant-blocked-088eb5eb
path: schemas
description: Details of a merchant_blocked outcome attached to this payment evaluation.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# insights_resources_payment_evaluation_merchant_blocked

Details of a merchant_blocked outcome attached to this payment evaluation.

```yaml
{"title": "InsightsResourcesPaymentEvaluationMerchantBlocked", "required": ["reason"], "type": "object", "properties": {"reason": {"type": "string", "description": "The reason the payment was blocked by the merchant.", "enum": ["authentication_required", "blocked_for_fraud", "invalid_payment", "other"]}}, "description": "Details of a merchant_blocked outcome attached to this payment evaluation.", "x-expandableFields": []}
```
