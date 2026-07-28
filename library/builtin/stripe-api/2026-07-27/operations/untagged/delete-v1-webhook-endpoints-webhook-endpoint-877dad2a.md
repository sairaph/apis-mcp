---
title: Delete a webhook endpoint
page_id: operation-delete-v1-webhook-endpoints-webhook-endpoint-3d65519d
path: operations/untagged
description: <p>You can also delete webhook endpoints via the <a href="https://dashboard.stripe.com/account/webhooks">webhook endpoint management</a> page of the Stripe dashboard.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/webhook_endpoints/{webhook_endpoint}
operation_ids:
    - DeleteWebhookEndpointsWebhookEndpoint
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a webhook endpoint

`DELETE /v1/webhook_endpoints/{webhook_endpoint}`

Operation ID: `DeleteWebhookEndpointsWebhookEndpoint`

<p>You can also delete webhook endpoints via the <a href="https://dashboard.stripe.com/account/webhooks">webhook endpoint management</a> page of the Stripe dashboard.</p>

## Definition

```yaml
{"summary": "Delete a webhook endpoint", "description": "<p>You can also delete webhook endpoints via the <a href=\"https://dashboard.stripe.com/account/webhooks\">webhook endpoint management</a> page of the Stripe dashboard.</p>", "operationId": "DeleteWebhookEndpointsWebhookEndpoint", "parameters": [{"name": "webhook_endpoint", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_webhook_endpoint"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
