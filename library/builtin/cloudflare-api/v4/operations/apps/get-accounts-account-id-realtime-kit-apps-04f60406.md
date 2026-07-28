---
title: Fetch all apps
page_id: operation-get-accounts-account-id-realtime-kit-apps-47fc2ae0
path: operations/apps
description: Fetch all apps for your account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/apps
operation_ids:
    - get_apps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all apps

`GET /accounts/{account_id}/realtime/kit/apps`

Operation ID: `get_apps`

Fetch all apps for your account

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}]
```

## Definition

```yaml
{"operationId": "get_apps", "summary": "Fetch all apps", "description": "Fetch all apps for your account", "parameters": [{"name": "page_no", "in": "query", "description": "The page number from which you want your page search results to be displayed.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"type": "integer", "default": 20, "minimum": 1}}, {"name": "search", "in": "query", "description": "Search string that matches apps by name.", "schema": {"type": "string", "maxLength": 125}}, {"name": "sort_order", "in": "query", "description": "Sort order for apps by creation time.", "schema": {"type": "string", "default": "DESC", "enum": ["ASC", "DESC"]}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"Example 1": {"value": {"data": [{"created_at": "2025-01-01T08:16:40.644Z", "id": "14a396e7-ca44-4937-bf1f-050a69118543", "name": "my-first-app"}], "paging": {"end_offset": 1, "start_offset": 1, "total_count": 1}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}}, "type": "object"}}, "paging": {"type": "object", "properties": {"end_offset": {"type": "number"}, "start_offset": {"type": "number"}, "total_count": {"type": "number"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Apps"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
