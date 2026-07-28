---
title: notification_event_request
page_id: schema-notification-event-request-d6abb85b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# notification_event_request

```yaml
{"title": "NotificationEventRequest", "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "ID of the API request that caused the event. If null, the event was automatic (e.g., Stripe's automatic subscription handling). Request logs are available in the [dashboard](https://dashboard.stripe.com/logs), but currently not in the API.", "nullable": true}, "idempotency_key": {"maxLength": 5000, "type": "string", "description": "The idempotency key transmitted during the request, if any. *Note: This property is populated only for events on or after May 23, 2017*.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
