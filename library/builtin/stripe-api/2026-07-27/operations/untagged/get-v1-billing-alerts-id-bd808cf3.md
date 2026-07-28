---
title: Retrieve a billing alert
page_id: operation-get-v1-billing-alerts-id-f7a4d103
path: operations/untagged
description: <p>Retrieves a billing alert given an ID</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/billing/alerts/{id}
operation_ids:
    - GetBillingAlertsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a billing alert

`GET /v1/billing/alerts/{id}`

Operation ID: `GetBillingAlertsId`

<p>Retrieves a billing alert given an ID</p>

## Definition

```yaml
{"summary": "Retrieve a billing alert", "description": "<p>Retrieves a billing alert given an ID</p>", "operationId": "GetBillingAlertsId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billing.alert"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
