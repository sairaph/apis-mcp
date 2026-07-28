---
title: Create a Meter Event Adjustment
page_id: operation-post-v2-billing-meter-event-adjustments-1bc13068
path: operations/untagged
description: Creates a meter event adjustment to cancel a previously sent meter event.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v2/billing/meter_event_adjustments
operation_ids:
    - PostV2BillingMeterEventAdjustments
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a Meter Event Adjustment

`POST /v2/billing/meter_event_adjustments`

Operation ID: `PostV2BillingMeterEventAdjustments`

Creates a meter event adjustment to cancel a previously sent meter event.

## Definition

```yaml
{"summary": "Create a Meter Event Adjustment", "description": "Creates a meter event adjustment to cancel a previously sent meter event.", "operationId": "PostV2BillingMeterEventAdjustments", "requestBody": {"content": {"application/json": {"schema": {"required": ["cancel", "event_name", "type"], "type": "object", "properties": {"cancel": {"required": ["identifier"], "type": "object", "properties": {"identifier": {"type": "string", "description": "The identifier that was originally assigned to the meter event. You can only cancel events within 24 hours of Stripe receiving them."}}, "description": "Specifies which event to cancel."}, "event_name": {"type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "type": {"type": "string", "description": "Specifies the type of cancellation. Currently supports canceling a single event.", "enum": ["cancel"]}}}}}}, "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.billing.meter_event_adjustment"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.invalid_cancel_configuration"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
