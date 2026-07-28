---
title: List flags
page_id: operation-get-accounts-account-id-flagship-apps-app-id-flags-f665eec2
path: operations/flags
description: Lists an app's flags ordered by key. Pass `cursor` from `result_info` to page forward; a null cursor indicates the last page.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/flagship/apps/{app_id}/flags
operation_ids:
    - flagship_list_flags
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List flags

`GET /accounts/{account_id}/flagship/apps/{app_id}/flags`

Operation ID: `flagship_list_flags`

Lists an app's flags ordered by key. Pass `cursor` from `result_info` to page forward; a null cursor indicates the last page.

## Definition

```yaml
{"operationId": "flagship_list_flags", "summary": "List flags", "description": "Lists an app's flags ordered by key. Pass `cursor` from `result_info` to page forward; a null cursor indicates the last page.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"description": "Cloudflare account ID.", "type": "string"}}, {"name": "app_id", "in": "path", "description": "App identifier.", "required": true, "schema": {"description": "App identifier.", "type": "string"}}, {"name": "limit", "in": "query", "description": "Max items to return (1–200).", "schema": {"description": "Max items to return (1–200).", "type": "string"}}, {"name": "cursor", "in": "query", "description": "Pagination cursor from a previous response.", "schema": {"description": "Pagination cursor from a previous response.", "type": "string"}}], "responses": {"200": {"description": "Paginated list of flags.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/flagship_FlagsPage"}, "result_info": {"$ref": "#/components/schemas/flagship_ResultInfo"}, "success": {"type": "boolean"}}, "required": ["success", "result", "result_info", "errors", "messages"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "429": {"description": "Too many requests.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/flagship_Error"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Flags"], "x-api-token-group": ["Flagship Read", "Flagship Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.flagship.app.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "flagship.apps.flags", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
