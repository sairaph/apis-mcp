---
title: Retrieve a Source MandateNotification
page_id: operation-get-v1-sources-source-mandate-notifications-mandate-notification-17cd8354
path: operations/untagged
description: <p>Retrieves a new Source MandateNotification.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/sources/{source}/mandate_notifications/{mandate_notification}
operation_ids:
    - GetSourcesSourceMandateNotificationsMandateNotification
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Source MandateNotification

`GET /v1/sources/{source}/mandate_notifications/{mandate_notification}`

Operation ID: `GetSourcesSourceMandateNotificationsMandateNotification`

<p>Retrieves a new Source MandateNotification.</p>

## Definition

```yaml
{"summary": "Retrieve a Source MandateNotification", "description": "<p>Retrieves a new Source MandateNotification.</p>", "operationId": "GetSourcesSourceMandateNotificationsMandateNotification", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "mandate_notification", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "source", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/source_mandate_notification"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
