---
title: Get Script Settings
page_id: operation-get-accounts-account-id-workers-scripts-script-name-script-settings-e0fdd4a5
path: operations/worker-script
description: Get script-level settings when using [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions). Includes Logpush and Tail Consumers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/script-settings
operation_ids:
    - worker-script-settings-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Script Settings

`GET /accounts/{account_id}/workers/scripts/{script_name}/script-settings`

Operation ID: `worker-script-settings-get-settings`

Get script-level settings when using [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions). Includes Logpush and Tail Consumers.

## Definition

```yaml
{"operationId": "worker-script-settings-get-settings", "summary": "Get Script Settings", "description": "Get script-level settings when using [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions). Includes Logpush and Tail Consumers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Fetch script settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-settings-response"}}}}, "4XX": {"description": "Fetch script settings failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
