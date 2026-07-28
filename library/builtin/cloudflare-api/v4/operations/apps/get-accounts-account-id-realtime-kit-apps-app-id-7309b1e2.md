---
title: Fetch app details
page_id: operation-get-accounts-account-id-realtime-kit-apps-app-id-e10f6599
path: operations/apps
description: Fetch details for an app in your account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/apps/{app_id}
operation_ids:
    - get_app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch app details

`GET /accounts/{account_id}/realtime/kit/apps/{app_id}`

Operation ID: `get_app`

Fetch details for an app in your account.

## Definition

```yaml
{"operationId": "get_app", "summary": "Fetch app details", "description": "Fetch details for an app in your account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"Example 1": {"value": {"data": {"created_at": "2025-01-01T08:16:40.644Z", "id": "14a396e7-ca44-4937-bf1f-050a69118543", "name": "my-first-app"}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string", "nullable": true}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Apps"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
