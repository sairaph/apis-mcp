---
title: Create app
page_id: operation-post-accounts-account-id-flagship-apps-098d19f9
path: operations/apps
description: Creates an app. The returned `id` is used in all subsequent flag, changelog, and evaluation requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/flagship/apps
operation_ids:
    - flagship_create_app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create app

`POST /accounts/{account_id}/flagship/apps`

Operation ID: `flagship_create_app`

Creates an app. The returned `id` is used in all subsequent flag, changelog, and evaluation requests.

## Definition

```yaml
{"operationId": "flagship_create_app", "summary": "Create app", "description": "Creates an app. The returned `id` is used in all subsequent flag, changelog, and evaluation requests.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_-]+$"}}, "additionalProperties": false, "required": ["name"]}}}}, "responses": {"201": {"description": "App created.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/flagship_App"}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "429": {"description": "Too many requests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Apps"], "x-api-token-group": ["Flagship Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "flagship.apps", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
