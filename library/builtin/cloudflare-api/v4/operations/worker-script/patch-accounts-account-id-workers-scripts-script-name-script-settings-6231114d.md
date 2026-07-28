---
title: Patch Script Settings
page_id: operation-patch-accounts-account-id-workers-scripts-script-name-script-settings-3adbc462
path: operations/worker-script
description: Patch script-level settings when using [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions). Including but not limited to Logpush and Tail Consumers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/script-settings
operation_ids:
    - worker-script-settings-patch-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Script Settings

`PATCH /accounts/{account_id}/workers/scripts/{script_name}/script-settings`

Operation ID: `worker-script-settings-patch-settings`

Patch script-level settings when using [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions). Including but not limited to Logpush and Tail Consumers.

## Definition

```yaml
{"operationId": "worker-script-settings-patch-settings", "summary": "Patch Script Settings", "description": "Patch script-level settings when using [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions). Including but not limited to Logpush and Tail Consumers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-settings-item"}}}}, "responses": {"200": {"description": "Patch script settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-settings-response"}}}}, "4XX": {"description": "Patch script settings failure. When the patch uses the\ndeclarative `exports` field and one or more entries fail\nreconciliation, the response is the exports reconciliation\nerror envelope (error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
