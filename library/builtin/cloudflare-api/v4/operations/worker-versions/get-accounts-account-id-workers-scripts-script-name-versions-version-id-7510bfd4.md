---
title: Get Version Detail
page_id: operation-get-accounts-account-id-workers-scripts-script-name-versions-version-id-e4e13c26
path: operations/worker-versions
description: Retrieves detailed information about a specific version of a Workers script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/versions/{version_id}
operation_ids:
    - worker-versions-get-version-detail
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Version Detail

`GET /accounts/{account_id}/workers/scripts/{script_name}/versions/{version_id}`

Operation ID: `worker-versions-get-version-detail`

Retrieves detailed information about a specific version of a Workers script.

## Definition

```yaml
{"operationId": "worker-versions-get-version-detail", "summary": "Get Version Detail", "description": "Retrieves detailed information about a specific version of a Workers script.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name-2"}}, {"name": "version_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_version_identifier"}}], "responses": {"200": {"description": "Get Version Detail response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_versions-single-response"}}}}, "4XX": {"description": "Get Version Detail response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_versions-single-response"}, {"$ref": "#/components/schemas/workers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Versions"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.versions", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
