---
title: Create a Meter Event with asynchronous validation
page_id: operation-post-v2-billing-meter-event-stream-b7d3e8f4
path: operations/untagged
description: Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/billing/meter_event_stream
operation_ids:
    - PostV2BillingMeterEventStream
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Meter Event with asynchronous validation

`POST /v2/billing/meter_event_stream`

Operation ID: `PostV2BillingMeterEventStream`

Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.

## Definition

```yaml
{"summary": "Create a Meter Event with asynchronous validation", "description": "Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.", "operationId": "PostV2BillingMeterEventStream", "requestBody": {"content": {"application/json": {"schema": {"required": ["events"], "type": "object", "properties": {"events": {"type": "array", "description": "List of meter events to include in the request. Supports up to 100 events per request.", "items": {"required": ["event_name", "payload"], "type": "object", "properties": {"event_name": {"type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "identifier": {"type": "string", "description": "A unique identifier for the event. If not provided, one will be generated.\nWe recommend using a globally unique identifier for this. We'll enforce\nuniqueness within a rolling 24 hour period."}, "payload": {"type": "object", "additionalProperties": {"type": "string"}, "description": "The payload of the event. This must contain the fields corresponding to a meter's\n`customer_mapping.event_payload_key` (default is `stripe_customer_id`) and\n`value_settings.event_payload_key` (default is `value`). Read more about\nthe\n[payload](https://docs.stripe.com/billing/subscriptions/usage-based/recording-usage#payload-key-overrides)."}, "timestamp": {"type": "string", "description": "The time of the event. Must be within the past 35 calendar days or up to\n5 minutes in the future. Defaults to current timestamp if not specified.", "format": "date-time"}}}}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.billing_meter_event_session_expired"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}, "servers": [{"url": "https://meter-events.stripe.com/"}]}
```
