---
title: Create a Meter Event with synchronous validation
page_id: operation-post-v2-billing-meter-events-087e0b28
path: operations/untagged
description: Creates a meter event. Events are validated synchronously, but are processed asynchronously. Supports up to 1,000 events per second in livemode. For higher rate-limits, please use meter event streams instead.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/billing/meter_events
operation_ids:
    - PostV2BillingMeterEvents
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Meter Event with synchronous validation

`POST /v2/billing/meter_events`

Operation ID: `PostV2BillingMeterEvents`

Creates a meter event. Events are validated synchronously, but are processed asynchronously. Supports up to 1,000 events per second in livemode. For higher rate-limits, please use meter event streams instead.

## Definition

```yaml
{"summary": "Create a Meter Event with synchronous validation", "description": "Creates a meter event. Events are validated synchronously, but are processed asynchronously. Supports up to 1,000 events per second in livemode. For higher rate-limits, please use meter event streams instead.", "operationId": "PostV2BillingMeterEvents", "requestBody": {"content": {"application/json": {"schema": {"required": ["event_name", "payload"], "type": "object", "properties": {"event_name": {"type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "identifier": {"type": "string", "description": "A unique identifier for the event. If not provided, one will be generated.\nWe recommend using a globally unique identifier for this. We'll enforce\nuniqueness within a rolling 24 hour period."}, "payload": {"type": "object", "additionalProperties": {"type": "string"}, "description": "The payload of the event. This must contain the fields corresponding to a meter's\n`customer_mapping.event_payload_key` (default is `stripe_customer_id`) and\n`value_settings.event_payload_key` (default is `value`). Read more about\nthe\n[payload](https://docs.stripe.com/billing/subscriptions/usage-based/recording-usage#payload-key-overrides)."}, "timestamp": {"type": "string", "description": "The time of the event. Must be within the past 35 calendar days or up to\n5 minutes in the future. Defaults to current timestamp if not specified.", "format": "date-time"}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.billing.meter_event"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.archived_meter"}, {"$ref": "#/components/schemas/v2.error.duplicate_meter_event"}, {"$ref": "#/components/schemas/v2.error.no_meter"}, {"$ref": "#/components/schemas/v2.error.payload_invalid_value"}, {"$ref": "#/components/schemas/v2.error.payload_no_customer_defined"}, {"$ref": "#/components/schemas/v2.error.payload_no_value_defined"}, {"$ref": "#/components/schemas/v2.error.too_many_concurrent_requests"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
