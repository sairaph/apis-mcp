---
title: List Turnstile Widgets
page_id: operation-get-accounts-account-id-challenges-widgets-f5499311
path: operations/turnstile
description: Lists all turnstile widgets of an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/challenges/widgets
operation_ids:
    - accounts-turnstile-widgets-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Turnstile Widgets

`GET /accounts/{account_id}/challenges/widgets`

Operation ID: `accounts-turnstile-widgets-list`

Lists all turnstile widgets of an account.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/turnstile_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of items per page.", "type": "number", "default": 25, "maximum": 1000, "minimum": 5}}, {"name": "order", "in": "query", "schema": {"description": "Field to order widgets by.", "type": "string", "example": "id", "enum": ["id", "sitekey", "name", "created_on", "modified_on"]}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order widgets.", "type": "string", "example": "asc", "enum": ["asc", "desc"]}}, {"name": "filter", "in": "query", "description": "Filter widgets by field using case-insensitive substring matching.\nFormat: `field:value`\n\nSupported fields:\n- `name` - Filter by widget name (e.g., `filter=name:login-form`)\n- `sitekey` - Filter by sitekey (e.g., `filter=sitekey:0x4AAA`)\n\nReturns 400 Bad Request if the field is unsupported or format is invalid.\nAn empty filter value returns all results.\n", "schema": {"type": "string", "example": "name:my-widget"}}]
```

## Definition

```yaml
{"operationId": "accounts-turnstile-widgets-list", "summary": "List Turnstile Widgets", "description": "Lists all turnstile widgets of an account.", "responses": {"200": {"description": "List Turnstile Widgets", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/turnstile_api-response-common"}, {"properties": {"result_info": {"$ref": "#/components/schemas/turnstile_result_info"}}, "type": "object"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/turnstile_widget_list"}}}, "type": "object"}]}}}}, "4XX": {"description": "List Turnstile Widgets Error", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/turnstile_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Turnstile"], "x-api-token-group": ["Turnstile Sites Write", "Turnstile Sites Read", "Account Settings Write", "Account Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "turnstile.widgets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
