---
title: Find a Secret
page_id: operation-get-v1-apps-secrets-find-f0c0966c
path: operations/untagged
description: <p>Finds a secret in the secret store by name and scope.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/apps/secrets/find
operation_ids:
    - GetAppsSecretsFind
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Find a Secret

`GET /v1/apps/secrets/find`

Operation ID: `GetAppsSecretsFind`

<p>Finds a secret in the secret store by name and scope.</p>

## Definition

```yaml
{"summary": "Find a Secret", "description": "<p>Finds a secret in the secret store by name and scope.</p>", "operationId": "GetAppsSecretsFind", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "name", "in": "query", "description": "A name for the secret that's unique within the scope.", "required": true, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "scope", "in": "query", "description": "Specifies the scoping of the secret. Requests originating from UI extensions can only access account-scoped secrets or secrets scoped to their own user.", "required": true, "style": "deepObject", "explode": true, "schema": {"title": "scope_param", "required": ["type"], "type": "object", "properties": {"type": {"type": "string", "enum": ["account", "user"]}, "user": {"maxLength": 5000, "type": "string"}}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/apps.secret"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
