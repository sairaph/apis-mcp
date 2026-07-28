---
title: v2.billing.meter_event_adjustment
page_id: schema-v2-billing-meter-event-adjustment-530f2e4e
path: schemas
description: A Meter Event Adjustment is used to cancel or modify previously recorded meter events. Meter Event Adjustments allow you to correct billing data by canceling individual events or event ranges, with tracking of adjustment status and creation time.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.billing.meter_event_adjustment

A Meter Event Adjustment is used to cancel or modify previously recorded meter events. Meter Event Adjustments allow you to correct billing data by canceling individual events or event ranges, with tracking of adjustment status and creation time.

```yaml
{"title": "Meter Event Adjustment", "required": ["cancel", "created", "event_name", "id", "livemode", "object", "status", "type"], "type": "object", "properties": {"cancel": {"required": ["identifier"], "type": "object", "properties": {"identifier": {"type": "string", "description": "The identifier that was originally assigned to the meter event. You can only cancel events within 24 hours of Stripe receiving them."}}, "description": "Specifies which event to cancel."}, "created": {"type": "string", "description": "The time the adjustment was created.", "format": "date-time"}, "event_name": {"type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "id": {"type": "string", "description": "The unique ID of this meter event adjustment."}, "livemode": {"type": "boolean", "description": "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value of the object field.", "enum": ["v2.billing.meter_event_adjustment"]}, "status": {"type": "string", "description": "Open Enum. The meter event adjustment's status.", "enum": ["complete", "pending"]}, "type": {"type": "string", "description": "Open Enum. Specifies the type of cancellation. Currently supports canceling a single event.", "enum": ["cancel"]}}, "description": "A Meter Event Adjustment is used to cancel or modify previously recorded meter events. Meter Event Adjustments allow you to correct billing data by canceling individual events or event ranges, with tracking of adjustment status and creation time."}
```
