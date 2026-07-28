---
title: Create a billing alert
page_id: operation-post-v1-billing-alerts-68485972
path: operations/untagged
description: <p>Creates a billing alert</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/alerts
operation_ids:
    - PostBillingAlerts
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a billing alert

`POST /v1/billing/alerts`

Operation ID: `PostBillingAlerts`

<p>Creates a billing alert</p>

## Definition

```yaml
{"summary": "Create a billing alert", "description": "<p>Creates a billing alert</p>", "operationId": "PostBillingAlerts", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["alert_type", "title"], "type": "object", "properties": {"alert_type": {"type": "string", "description": "The type of alert to create.", "enum": ["usage_threshold"], "x-stripeBypassValidation": true}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "title": {"maxLength": 256, "type": "string", "description": "The title of the alert."}, "usage_threshold": {"title": "usage_threshold_config", "required": ["gte", "meter", "recurrence"], "type": "object", "properties": {"filters": {"type": "array", "items": {"title": "usage_alert_filter", "required": ["type"], "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["customer"]}}}}, "gte": {"type": "integer"}, "meter": {"maxLength": 5000, "type": "string"}, "recurrence": {"type": "string", "enum": ["one_time"], "x-stripeBypassValidation": true}}, "description": "The configuration of the usage threshold."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "usage_threshold": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.alert"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
