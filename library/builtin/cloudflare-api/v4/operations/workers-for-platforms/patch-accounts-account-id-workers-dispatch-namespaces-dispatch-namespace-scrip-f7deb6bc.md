---
title: Patch Script Settings
page_id: operation-patch-accounts-account-id-workers-dispatch-namespaces-dispatch-namespace-3632b297
path: operations/workers-for-platforms
description: Patch script metadata, such as bindings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings
operation_ids:
    - namespace-worker-patch-script-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Script Settings

`PATCH /accounts/{account_id}/workers/dispatch/namespaces/{dispatch_namespace}/scripts/{script_name}/settings`

Operation ID: `namespace-worker-patch-script-settings`

Patch script metadata, such as bindings.

## Definition

```yaml
{"operationId": "namespace-worker-patch-script-settings", "summary": "Patch Script Settings", "description": "Patch script metadata, such as bindings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "dispatch_namespace", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_dispatch_namespace_name"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"encoding": {"settings": {"contentType": "application/json"}}, "schema": {"type": "object", "properties": {"settings": {"$ref": "#/components/schemas/workers_namespace-script-and-version-settings-item"}}}}}}, "responses": {"200": {"description": "Patch script settings.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_namespace-script-and-version-settings-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Patch script settings failure. When the patch uses the\ndeclarative `exports` field and one or more entries fail\nreconciliation, the response is the exports reconciliation\nerror envelope (error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers for Platforms"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dispatch-namespaces.scripts.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
