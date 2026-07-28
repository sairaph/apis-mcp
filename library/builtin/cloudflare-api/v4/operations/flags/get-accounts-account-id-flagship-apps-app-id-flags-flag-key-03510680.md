---
title: Get flag
page_id: operation-get-accounts-account-id-flagship-apps-app-id-flags-flag-key-684249aa
path: operations/flags
description: Returns the full flag definition including rules, variations, and audit fields.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/flagship/apps/{app_id}/flags/{flag_key}
operation_ids:
    - flagship_get_flag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get flag

`GET /accounts/{account_id}/flagship/apps/{app_id}/flags/{flag_key}`

Operation ID: `flagship_get_flag`

Returns the full flag definition including rules, variations, and audit fields.

## Definition

```yaml
{"operationId": "flagship_get_flag", "summary": "Get flag", "description": "Returns the full flag definition including rules, variations, and audit fields.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}, {"name": "app_id", "in": "path", "description": "App identifier.", "required": true, "schema": {"description": "App identifier.", "type": "string"}}, {"name": "flag_key", "in": "path", "description": "Flag key (slug).", "required": true, "schema": {"description": "Flag key (slug).", "type": "string"}}], "responses": {"200": {"description": "Flag details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/flagship_Flag"}, "success": {"type": "boolean"}}, "required": ["success", "result", "errors", "messages"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "429": {"description": "Too many requests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Flags"], "x-api-token-group": ["Flagship Read", "Flagship Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.app.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "flagship.apps.flags", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
