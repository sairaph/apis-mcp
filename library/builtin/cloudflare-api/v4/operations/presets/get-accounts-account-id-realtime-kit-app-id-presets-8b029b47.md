---
title: Fetch all presets
page_id: operation-get-accounts-account-id-realtime-kit-app-id-presets-5c601e4d
path: operations/presets
description: Fetches all the presets belonging to an App.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/presets
operation_ids:
    - get-presets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all presets

`GET /accounts/{account_id}/realtime/kit/{app_id}/presets`

Operation ID: `get-presets`

Fetches all the presets belonging to an App.

## Definition

```yaml
{"operationId": "get-presets", "summary": "Fetch all presets", "description": "Fetches all the presets belonging to an App.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_perPage"}, {"$ref": "#/components/parameters/realtimekit_pageNo"}, {"name": "search", "in": "query", "description": "Search presets by name.", "schema": {"type": "string"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetAllPresets"}}, "security": [{"api_token": []}], "tags": ["Presets"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
