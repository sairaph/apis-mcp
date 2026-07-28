---
title: Update app
page_id: operation-put-accounts-account-id-flagship-apps-app-id-48cfed82
path: operations/apps
description: Updates an app. Only `name` is mutable.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/flagship/apps/{app_id}
operation_ids:
    - flagship_update_app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update app

`PUT /accounts/{account_id}/flagship/apps/{app_id}`

Operation ID: `flagship_update_app`

Updates an app. Only `name` is mutable.

## Definition

```yaml
{"operationId": "flagship_update_app", "summary": "Update app", "description": "Updates an app. Only `name` is mutable.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}, {"name": "app_id", "in": "path", "description": "App identifier.", "required": true, "schema": {"description": "App identifier.", "type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_-]+$"}}, "additionalProperties": false}}}}, "responses": {"200": {"description": "Updated app.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/flagship_App"}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "429": {"description": "Too many requests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Apps"], "x-api-token-group": ["Flagship Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.app.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "flagship.apps", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
