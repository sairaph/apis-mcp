---
title: Delete a Secret
page_id: operation-post-v1-apps-secrets-delete-68c3666c
path: operations/untagged
description: <p>Deletes a secret from the secret store by name and scope.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/apps/secrets/delete
operation_ids:
    - PostAppsSecretsDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a Secret

`POST /v1/apps/secrets/delete`

Operation ID: `PostAppsSecretsDelete`

<p>Deletes a secret from the secret store by name and scope.</p>

## Definition

```yaml
{"summary": "Delete a Secret", "description": "<p>Deletes a secret from the secret store by name and scope.</p>", "operationId": "PostAppsSecretsDelete", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["name", "scope"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "name": {"maxLength": 5000, "type": "string", "description": "A name for the secret that's unique within the scope."}, "scope": {"title": "scope_param", "required": ["type"], "type": "object", "properties": {"type": {"type": "string", "enum": ["account", "user"]}, "user": {"maxLength": 5000, "type": "string"}}, "description": "Specifies the scoping of the secret. Requests originating from UI extensions can only access account-scoped secrets or secrets scoped to their own user."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "scope": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/apps.secret"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
