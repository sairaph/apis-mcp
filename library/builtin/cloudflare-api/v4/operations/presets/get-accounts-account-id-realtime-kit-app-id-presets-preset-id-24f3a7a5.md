---
title: Fetch details of a preset
page_id: operation-get-accounts-account-id-realtime-kit-app-id-presets-preset-id-ec3ef597
path: operations/presets
description: Fetches details of a preset using the provided preset ID
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/presets/{preset_id}
operation_ids:
    - get-presets-preset_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch details of a preset

`GET /accounts/{account_id}/realtime/kit/{app_id}/presets/{preset_id}`

Operation ID: `get-presets-preset_id`

Fetches details of a preset using the provided preset ID

## Definition

```yaml
{"operationId": "get-presets-preset_id", "summary": "Fetch details of a preset", "description": "Fetches details of a preset using the provided preset ID", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "preset_id", "in": "path", "description": "ID of the preset to fetch", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetPresetBody"}}, "security": [{"api_token": []}], "tags": ["Presets"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
