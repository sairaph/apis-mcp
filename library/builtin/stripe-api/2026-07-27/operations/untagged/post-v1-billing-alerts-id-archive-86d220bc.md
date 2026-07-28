---
title: Archive a billing alert
page_id: operation-post-v1-billing-alerts-id-archive-4fd5e186
path: operations/untagged
description: <p>Archives this alert, removing it from the list view and APIs. This is non-reversible.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/billing/alerts/{id}/archive
operation_ids:
    - PostBillingAlertsIdArchive
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Archive a billing alert

`POST /v1/billing/alerts/{id}/archive`

Operation ID: `PostBillingAlertsIdArchive`

<p>Archives this alert, removing it from the list view and APIs. This is non-reversible.</p>

## Definition

```yaml
{"summary": "Archive a billing alert", "description": "<p>Archives this alert, removing it from the list view and APIs. This is non-reversible.</p>", "operationId": "PostBillingAlertsIdArchive", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.alert"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
