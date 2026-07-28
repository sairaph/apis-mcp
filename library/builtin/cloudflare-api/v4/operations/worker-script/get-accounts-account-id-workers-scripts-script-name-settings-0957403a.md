---
title: Get Settings
page_id: operation-get-accounts-account-id-workers-scripts-script-name-settings-1936291a
path: operations/worker-script
description: Get metadata and config, such as bindings or usage model.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/settings
operation_ids:
    - worker-script-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Settings

`GET /accounts/{account_id}/workers/scripts/{script_name}/settings`

Operation ID: `worker-script-get-settings`

Get metadata and config, such as bindings or usage model.

## Definition

```yaml
{"operationId": "worker-script-get-settings", "summary": "Get Settings", "description": "Get metadata and config, such as bindings or usage model.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Fetch settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-and-version-settings-response"}}}}, "4XX": {"description": "Fetch settings failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.script-and-version-settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
