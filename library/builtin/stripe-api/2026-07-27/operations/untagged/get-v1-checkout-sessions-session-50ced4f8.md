---
title: Retrieve a Checkout Session
page_id: operation-get-v1-checkout-sessions-session-478458e1
path: operations/untagged
description: <p>Retrieves a Checkout Session object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/checkout/sessions/{session}
operation_ids:
    - GetCheckoutSessionsSession
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Checkout Session

`GET /v1/checkout/sessions/{session}`

Operation ID: `GetCheckoutSessionsSession`

<p>Retrieves a Checkout Session object.</p>

## Definition

```yaml
{"summary": "Retrieve a Checkout Session", "description": "<p>Retrieves a Checkout Session object.</p>", "operationId": "GetCheckoutSessionsSession", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 66, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/checkout.session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
