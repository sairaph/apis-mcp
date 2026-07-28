---
title: notification_event_data
page_id: schema-notification-event-data-57d06963
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# notification_event_data

```yaml
{"title": "NotificationEventData", "required": ["object"], "type": "object", "properties": {"object": {"type": "object", "description": "Object containing the API resource relevant to the event. For example, an `invoice.created` event will have a full [invoice object](https://api.stripe.com#invoice_object) as the value of the object key."}, "previous_attributes": {"type": "object", "description": "Object containing the names of the updated attributes and their values prior to the event (only included in events of type `*.updated`). If an array attribute has any updated elements, this object contains the entire array. In Stripe API versions 2017-04-06 or earlier, an updated array attribute in this object includes only the updated array elements."}}, "description": "", "x-expandableFields": []}
```
