---
title: Update a preset
page_id: operation-patch-accounts-account-id-realtime-kit-app-id-presets-preset-id-310ee827
path: operations/presets
description: Update a preset by the provided preset ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/presets/{preset_id}
operation_ids:
    - patch-presets-preset_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a preset

`PATCH /accounts/{account_id}/realtime/kit/{app_id}/presets/{preset_id}`

Operation ID: `patch-presets-preset_id`

Update a preset by the provided preset ID

## Definition

```yaml
{"operationId": "patch-presets-preset_id", "summary": "Update a preset", "description": "Update a preset by the provided preset ID", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "preset_id", "in": "path", "description": "ID of the preset to fetch", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_UpdatePresetBody"}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetPresetBody"}}, "security": [{"api_token": []}], "tags": ["Presets"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
