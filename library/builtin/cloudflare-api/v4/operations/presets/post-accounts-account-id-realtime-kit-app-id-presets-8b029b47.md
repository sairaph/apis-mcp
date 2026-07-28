---
title: Create a preset
page_id: operation-post-accounts-account-id-realtime-kit-app-id-presets-50bd8a50
path: operations/presets
description: Creates a preset belonging to the current App
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/presets
operation_ids:
    - post-presets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a preset

`POST /accounts/{account_id}/realtime/kit/{app_id}/presets`

Operation ID: `post-presets`

Creates a preset belonging to the current App

## Definition

```yaml
{"operationId": "post-presets", "summary": "Create a preset", "description": "Creates a preset belonging to the current App", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_CreatePresetBody"}, "responses": {"201": {"$ref": "#/components/responses/realtimekit_GetPresetBody"}}, "security": [{"api_token": []}], "tags": ["Presets"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
