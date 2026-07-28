---
title: Patch Settings
page_id: operation-patch-accounts-account-id-workers-scripts-script-name-settings-854e7ffd
path: operations/worker-script
description: Patch metadata or config, such as bindings or usage model.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/settings
operation_ids:
    - worker-script-patch-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Settings

`PATCH /accounts/{account_id}/workers/scripts/{script_name}/settings`

Operation ID: `worker-script-patch-settings`

Patch metadata or config, such as bindings or usage model.

## Definition

```yaml
{"operationId": "worker-script-patch-settings", "summary": "Patch Settings", "description": "Patch metadata or config, such as bindings or usage model.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"encoding": {"settings": {"contentType": "application/json"}}, "schema": {"type": "object", "properties": {"settings": {"$ref": "#/components/schemas/workers_script-and-version-settings-item"}}}}}}, "responses": {"200": {"description": "Patch settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-and-version-settings-response"}}}}, "4XX": {"description": "Patch settings failure. When the patch uses the declarative\n`exports` field and one or more entries fail reconciliation,\nthe response is the exports reconciliation error envelope\n(error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.script-and-version-settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
