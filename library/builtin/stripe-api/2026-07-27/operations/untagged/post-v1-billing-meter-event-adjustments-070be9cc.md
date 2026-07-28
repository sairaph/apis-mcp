---
title: Create a billing meter event adjustment
page_id: operation-post-v1-billing-meter-event-adjustments-fc23456c
path: operations/untagged
description: <p>Creates a billing meter event adjustment.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/meter_event_adjustments
operation_ids:
    - PostBillingMeterEventAdjustments
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a billing meter event adjustment

`POST /v1/billing/meter_event_adjustments`

Operation ID: `PostBillingMeterEventAdjustments`

<p>Creates a billing meter event adjustment.</p>

## Definition

```yaml
{"summary": "Create a billing meter event adjustment", "description": "<p>Creates a billing meter event adjustment.</p>", "operationId": "PostBillingMeterEventAdjustments", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["event_name", "type"], "type": "object", "properties": {"cancel": {"title": "event_adjustment_cancel_settings_param", "type": "object", "properties": {"identifier": {"maxLength": 100, "type": "string"}}, "description": "Specifies which event to cancel."}, "event_name": {"maxLength": 100, "type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "type": {"type": "string", "description": "Specifies whether to cancel a single event or a range of events for a time period. Time period cancellation is not supported yet.", "enum": ["cancel"]}}, "additionalProperties": false}, "encoding": {"cancel": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.meter_event_adjustment"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
