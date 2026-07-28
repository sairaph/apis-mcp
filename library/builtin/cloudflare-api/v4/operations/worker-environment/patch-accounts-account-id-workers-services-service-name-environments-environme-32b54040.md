---
title: Patch Script Settings
page_id: operation-patch-accounts-account-id-workers-services-service-name-environments-env-cf249058
path: operations/worker-environment
description: Patch script metadata, such as bindings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/services/{service_name}/environments/{environment_name}/settings
operation_ids:
    - worker-script-environment-patch-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Script Settings

`PATCH /accounts/{account_id}/workers/services/{service_name}/environments/{environment_name}/settings`

Operation ID: `worker-script-environment-patch-settings`

Patch script metadata, such as bindings.

## Definition

```yaml
{"operationId": "worker-script-environment-patch-settings", "summary": "Patch Script Settings", "description": "Patch script metadata, such as bindings.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "service_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_service"}}, {"name": "environment_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_environment"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-settings-response"}}}}, "responses": {"200": {"description": "Patch script settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-settings-response"}}}}, "4XX": {"description": "Patch script settings failure. When the patch uses the\ndeclarative `exports` field and one or more entries fail\nreconciliation, the response is the exports reconciliation\nerror envelope (error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Environment"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.services.environments.settings", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
