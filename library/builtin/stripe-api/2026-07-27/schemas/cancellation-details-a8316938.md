---
title: cancellation_details
page_id: schema-cancellation-details-a8316938
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# cancellation_details

```yaml
{"title": "CancellationDetails", "type": "object", "properties": {"comment": {"maxLength": 5000, "type": "string", "description": "Additional comments about why the user canceled the subscription, if the subscription was canceled explicitly by the user.", "nullable": true}, "feedback": {"type": "string", "description": "The customer submitted reason for why they canceled, if the subscription was canceled explicitly by the user.", "nullable": true, "enum": ["customer_service", "low_quality", "missing_features", "other", "switched_service", "too_complex", "too_expensive", "unused"]}, "reason": {"type": "string", "description": "Why this subscription was canceled.", "nullable": true, "enum": ["canceled_by_retention_policy", "cancellation_requested", "payment_disputed", "payment_failed"]}}, "description": "", "x-expandableFields": []}
```
