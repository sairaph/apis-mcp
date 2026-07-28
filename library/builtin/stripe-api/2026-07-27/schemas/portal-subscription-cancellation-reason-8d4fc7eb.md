---
title: portal_subscription_cancellation_reason
page_id: schema-portal-subscription-cancellation-reason-8d4fc7eb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_subscription_cancellation_reason

```yaml
{"title": "PortalSubscriptionCancellationReason", "required": ["enabled", "options"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Whether the feature is enabled."}, "options": {"type": "array", "description": "Which cancellation reasons will be given as options to the customer.", "items": {"type": "string", "enum": ["customer_service", "low_quality", "missing_features", "other", "switched_service", "too_complex", "too_expensive", "unused"]}}}, "description": "", "x-expandableFields": []}
```
