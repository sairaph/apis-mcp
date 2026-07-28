---
title: Deactivate a billing alert
page_id: operation-post-v1-billing-alerts-id-deactivate-1ef5c09c
path: operations/untagged
description: <p>Deactivates this alert, preventing it from triggering.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/alerts/{id}/deactivate
operation_ids:
    - PostBillingAlertsIdDeactivate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Deactivate a billing alert

`POST /v1/billing/alerts/{id}/deactivate`

Operation ID: `PostBillingAlertsIdDeactivate`

<p>Deactivates this alert, preventing it from triggering.</p>

## Definition

```yaml
{"summary": "Deactivate a billing alert", "description": "<p>Deactivates this alert, preventing it from triggering.</p>", "operationId": "PostBillingAlertsIdDeactivate", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.alert"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
