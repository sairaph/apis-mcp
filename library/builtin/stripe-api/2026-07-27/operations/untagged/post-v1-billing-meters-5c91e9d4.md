---
title: Create a billing meter
page_id: operation-post-v1-billing-meters-f552865e
path: operations/untagged
description: <p>Creates a billing meter.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/meters
operation_ids:
    - PostBillingMeters
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a billing meter

`POST /v1/billing/meters`

Operation ID: `PostBillingMeters`

<p>Creates a billing meter.</p>

## Definition

```yaml
{"summary": "Create a billing meter", "description": "<p>Creates a billing meter.</p>", "operationId": "PostBillingMeters", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["default_aggregation", "display_name", "event_name"], "type": "object", "properties": {"customer_mapping": {"title": "customer_mapping_param", "required": ["event_payload_key", "type"], "type": "object", "properties": {"event_payload_key": {"maxLength": 100, "type": "string"}, "type": {"type": "string", "enum": ["by_id"]}}, "description": "Fields that specify how to map a meter event to a customer."}, "default_aggregation": {"title": "aggregation_settings_param", "required": ["formula"], "type": "object", "properties": {"formula": {"type": "string", "enum": ["count", "last", "sum"], "x-stripeBypassValidation": true}}, "description": "The default settings to aggregate a meter's events with."}, "display_name": {"maxLength": 250, "type": "string", "description": "The meter’s name. Not visible to the customer."}, "event_name": {"maxLength": 100, "type": "string", "description": "The name of the meter event to record usage for. Corresponds with the `event_name` field on meter events."}, "event_time_window": {"type": "string", "description": "The time window which meter events have been pre-aggregated for, if any.", "enum": ["day", "hour"]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "value_settings": {"title": "meter_value_settings_param", "required": ["event_payload_key"], "type": "object", "properties": {"event_payload_key": {"maxLength": 100, "type": "string"}}, "description": "Fields that specify how to calculate a meter event's value."}}, "additionalProperties": false}, "encoding": {"customer_mapping": {"style": "deepObject", "explode": true}, "default_aggregation": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "value_settings": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.meter"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
