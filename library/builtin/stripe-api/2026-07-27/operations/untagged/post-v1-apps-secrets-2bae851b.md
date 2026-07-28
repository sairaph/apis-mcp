---
title: Set a Secret
page_id: operation-post-v1-apps-secrets-70dd99e4
path: operations/untagged
description: <p>Create or replace a secret in the secret store.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/apps/secrets
operation_ids:
    - PostAppsSecrets
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Set a Secret

`POST /v1/apps/secrets`

Operation ID: `PostAppsSecrets`

<p>Create or replace a secret in the secret store.</p>

## Definition

```yaml
{"summary": "Set a Secret", "description": "<p>Create or replace a secret in the secret store.</p>", "operationId": "PostAppsSecrets", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["name", "payload", "scope"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"type": "integer", "description": "The Unix timestamp for the expiry time of the secret, after which the secret deletes.", "format": "unix-time"}, "name": {"maxLength": 5000, "type": "string", "description": "A name for the secret that's unique within the scope."}, "payload": {"maxLength": 5000, "type": "string", "description": "The plaintext secret value to be stored."}, "scope": {"title": "scope_param", "required": ["type"], "type": "object", "properties": {"type": {"type": "string", "enum": ["account", "user"]}, "user": {"maxLength": 5000, "type": "string"}}, "description": "Specifies the scoping of the secret. Requests originating from UI extensions can only access account-scoped secrets or secrets scoped to their own user."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "scope": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/apps.secret"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
