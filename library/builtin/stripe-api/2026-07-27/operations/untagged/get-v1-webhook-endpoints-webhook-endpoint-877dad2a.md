---
title: Retrieve a webhook endpoint
page_id: operation-get-v1-webhook-endpoints-webhook-endpoint-068a4b22
path: operations/untagged
description: <p>Retrieves the webhook endpoint with the given ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/webhook_endpoints/{webhook_endpoint}
operation_ids:
    - GetWebhookEndpointsWebhookEndpoint
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a webhook endpoint

`GET /v1/webhook_endpoints/{webhook_endpoint}`

Operation ID: `GetWebhookEndpointsWebhookEndpoint`

<p>Retrieves the webhook endpoint with the given ID.</p>

## Definition

```yaml
{"summary": "Retrieve a webhook endpoint", "description": "<p>Retrieves the webhook endpoint with the given ID.</p>", "operationId": "GetWebhookEndpointsWebhookEndpoint", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "webhook_endpoint", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/webhook_endpoint"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
