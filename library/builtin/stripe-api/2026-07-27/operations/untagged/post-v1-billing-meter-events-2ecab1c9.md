---
title: Create a billing meter event
page_id: operation-post-v1-billing-meter-events-601664c7
path: operations/untagged
description: <p>Creates a billing meter event.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/meter_events
operation_ids:
    - PostBillingMeterEvents
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a billing meter event

`POST /v1/billing/meter_events`

Operation ID: `PostBillingMeterEvents`

<p>Creates a billing meter event.</p>

## Definition

```yaml
{"summary": "Create a billing meter event", "description": "<p>Creates a billing meter event.</p>", "operationId": "PostBillingMeterEvents", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["event_name", "payload"], "type": "object", "properties": {"event_name": {"maxLength": 100, "type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "identifier": {"maxLength": 100, "type": "string", "description": "A unique identifier for the event. If not provided, one is generated. We recommend using UUID-like identifiers. Stripe enforces uniqueness within a rolling period of at least 24 hours. The enforcement of uniqueness primarily addresses issues arising from accidental retries or other problems occurring within extremely brief time intervals. This approach helps prevent duplicate entries and ensures data integrity in high-frequency operations."}, "payload": {"type": "object", "additionalProperties": {"type": "string"}, "description": "The payload of the event. This must contain the fields corresponding to a meter's `customer_mapping.event_payload_key` (default is `stripe_customer_id`) and `value_settings.event_payload_key` (default is `value`). Read more about the [payload](https://docs.stripe.com/billing/subscriptions/usage-based/meters/configure#meter-configuration-attributes)."}, "timestamp": {"type": "integer", "description": "The time of the event. Measured in seconds since the Unix epoch. Must be within the past 35 calendar days or up to 5 minutes in the future. Defaults to current timestamp if not specified.", "format": "unix-time"}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "payload": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.meter_event"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
