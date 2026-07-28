---
title: Create App
page_id: operation-post-accounts-account-id-realtime-kit-apps-6a9cf186
path: operations/apps
description: Create new app for your account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/apps
operation_ids:
    - create_app
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create App

`POST /accounts/{account_id}/realtime/kit/apps`

Operation ID: `create_app`

Create new app for your account

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}]
```

## Definition

```yaml
{"operationId": "create_app", "summary": "Create App", "description": "Create new app for your account", "requestBody": {"$ref": "#/components/requestBodies/realtimekit_CreateApp"}, "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"Example 1": {"value": {"data": {"app": {"created_at": "2025-01-01T08:16:40.644Z", "id": "14a396e7-ca44-4937-bf1f-050a69118543", "name": "my-new-app"}}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"app": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}}}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Apps"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
