---
title: Replace a preset
page_id: operation-put-accounts-account-id-realtime-kit-app-id-presets-preset-id-4e55efe9
path: operations/presets
description: Replace all details for the preset using the provided preset ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/presets/{preset_id}
operation_ids:
    - put-presets-preset_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace a preset

`PUT /accounts/{account_id}/realtime/kit/{app_id}/presets/{preset_id}`

Operation ID: `put-presets-preset_id`

Replace all details for the preset using the provided preset ID.

## Definition

```yaml
{"operationId": "put-presets-preset_id", "summary": "Replace a preset", "description": "Replace all details for the preset using the provided preset ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "preset_id", "in": "path", "description": "ID of the preset to replace", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_CreatePresetBody"}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetPresetBody"}}, "security": [{"api_token": []}], "tags": ["Presets"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
