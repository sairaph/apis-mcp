---
title: Delete app
page_id: operation-delete-accounts-account-id-flagship-apps-app-id-7bc18078
path: operations/apps
description: Deletes an app and all its flags and changelog history. Returns 409 if any Worker still references this app via a Flagship binding.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/flagship/apps/{app_id}
operation_ids:
    - flagship_delete_app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete app

`DELETE /accounts/{account_id}/flagship/apps/{app_id}`

Operation ID: `flagship_delete_app`

Deletes an app and all its flags and changelog history. Returns 409 if any Worker still references this app via a Flagship binding.

## Definition

```yaml
{"operationId": "flagship_delete_app", "summary": "Delete app", "description": "Deletes an app and all its flags and changelog history. Returns 409 if any Worker still references this app via a Flagship binding.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}, {"name": "app_id", "in": "path", "description": "App identifier.", "required": true, "schema": {"description": "App identifier.", "type": "string"}}], "responses": {"200": {"description": "App deleted.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/flagship_DeleteAppResult"}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "409": {"description": "Conflict — app is bound to Worker scripts.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "429": {"description": "Too many requests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Apps"], "x-api-token-group": ["Flagship Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.app.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "flagship.apps", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
