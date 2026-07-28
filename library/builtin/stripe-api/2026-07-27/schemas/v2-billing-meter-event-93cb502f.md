---
title: v2.billing.meter_event
page_id: schema-v2-billing-meter-event-93cb502f
path: schemas
description: A Meter Event is a usage record that captures billable activity for usage-based billing. Meter Events contain an event name, timestamp, and payload with customer mapping and usage value, enabling accurate usage tracking and billing.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.billing.meter_event

A Meter Event is a usage record that captures billable activity for usage-based billing. Meter Events contain an event name, timestamp, and payload with customer mapping and usage value, enabling accurate usage tracking and billing.

```yaml
{"title": "Meter Event", "required": ["created", "event_name", "identifier", "livemode", "object", "payload", "timestamp"], "type": "object", "properties": {"created": {"type": "string", "description": "The creation time of this meter event.", "format": "date-time"}, "event_name": {"type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "identifier": {"type": "string", "description": "A unique identifier for the event. If not provided, one will be generated. We recommend using a globally unique identifier for this. We'll enforce uniqueness within a rolling 24 hour period."}, "livemode": {"type": "boolean", "description": "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value of the object field.", "enum": ["v2.billing.meter_event"]}, "payload": {"type": "object", "additionalProperties": {"type": "string"}, "description": "The payload of the event. This must contain the fields corresponding to a meter's\n`customer_mapping.event_payload_key` (default is `stripe_customer_id`) and\n`value_settings.event_payload_key` (default is `value`). Read more about\nthe [payload](https://docs.stripe.com/billing/subscriptions/usage-based/recording-usage#payload-key-overrides).."}, "timestamp": {"type": "string", "description": "The time of the event. Must be within the past 35 calendar days or up to\n5 minutes in the future. Defaults to current timestamp if not specified.", "format": "date-time"}}, "description": "A Meter Event is a usage record that captures billable activity for usage-based billing. Meter Events contain an event name, timestamp, and payload with customer mapping and usage value, enabling accurate usage tracking and billing."}
```
