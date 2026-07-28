---
title: Activate a billing alert
page_id: operation-post-v1-billing-alerts-id-activate-e969f75a
path: operations/untagged
description: <p>Reactivates this alert, allowing it to trigger again.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/alerts/{id}/activate
operation_ids:
    - PostBillingAlertsIdActivate
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Activate a billing alert

`POST /v1/billing/alerts/{id}/activate`

Operation ID: `PostBillingAlertsIdActivate`

<p>Reactivates this alert, allowing it to trigger again.</p>

## Definition

```yaml
{"summary": "Activate a billing alert", "description": "<p>Reactivates this alert, allowing it to trigger again.</p>", "operationId": "PostBillingAlertsIdActivate", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.alert"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
