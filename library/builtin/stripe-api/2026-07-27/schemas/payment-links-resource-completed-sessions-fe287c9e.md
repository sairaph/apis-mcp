---
title: payment_links_resource_completed_sessions
page_id: schema-payment-links-resource-completed-sessions-fe287c9e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_completed_sessions

```yaml
{"title": "PaymentLinksResourceCompletedSessions", "required": ["count", "limit"], "type": "object", "properties": {"count": {"type": "integer", "description": "The current number of checkout sessions that have been completed on the payment link which count towards the `completed_sessions` restriction to be met."}, "limit": {"type": "integer", "description": "The maximum number of checkout sessions that can be completed for the `completed_sessions` restriction to be met."}}, "description": "", "x-expandableFields": []}
```
