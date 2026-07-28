---
title: billing.meter_event
page_id: schema-billing-meter-event-6512761b
path: schemas
description: Meter events represent actions that customers take in your system. You can use meter events to bill a customer based on their usage. Meter events are associated with billing meters, which define both the contents of the event’s payload and how to aggregate those events.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.meter_event

Meter events represent actions that customers take in your system. You can use meter events to bill a customer based on their usage. Meter events are associated with billing meters, which define both the contents of the event’s payload and how to aggregate those events.

```yaml
{"title": "BillingMeterEvent", "required": ["created", "event_name", "identifier", "livemode", "object", "payload", "timestamp"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "event_name": {"maxLength": 100, "type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "identifier": {"maxLength": 5000, "type": "string", "description": "A unique identifier for the event."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.meter_event"]}, "payload": {"type": "object", "additionalProperties": {"maxLength": 100, "type": "string"}, "description": "The payload of the event. This contains the fields corresponding to a meter's `customer_mapping.event_payload_key` (default is `stripe_customer_id`) and `value_settings.event_payload_key` (default is `value`). Read more about the [payload](https://docs.stripe.com/billing/subscriptions/usage-based/meters/configure#meter-configuration-attributes)."}, "timestamp": {"type": "integer", "description": "The timestamp passed in when creating the event. Measured in seconds since the Unix epoch.", "format": "unix-time"}}, "description": "Meter events represent actions that customers take in your system. You can use meter events to bill a customer based on their usage. Meter events are associated with billing meters, which define both the contents of the event’s payload and how to aggregate those events.", "x-expandableFields": [], "x-resourceId": "billing.meter_event"}
```
