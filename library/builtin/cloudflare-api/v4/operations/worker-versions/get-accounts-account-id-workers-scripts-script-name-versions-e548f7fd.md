---
title: List Versions
page_id: operation-get-accounts-account-id-workers-scripts-script-name-versions-c074eec2
path: operations/worker-versions
description: List of Worker Versions. The first version in the list is the latest version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/versions
operation_ids:
    - worker-versions-list-versions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Versions

`GET /accounts/{account_id}/workers/scripts/{script_name}/versions`

Operation ID: `worker-versions-list-versions`

List of Worker Versions. The first version in the list is the latest version.

## Definition

```yaml
{"operationId": "worker-versions-list-versions", "summary": "List Versions", "description": "List of Worker Versions. The first version in the list is the latest version.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name-2"}}, {"name": "deployable", "in": "query", "description": "Only return versions that can be used in a deployment. Ignores pagination.", "schema": {"type": "boolean", "default": false}}, {"name": "page", "in": "query", "description": "Current page.", "schema": {"type": "integer", "default": 1}}, {"name": "per_page", "in": "query", "description": "Items per-page.", "schema": {"type": "integer"}}], "responses": {"200": {"description": "List Versions response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_versions-list-response"}}}}, "4XX": {"description": "List Versions response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_versions-list-response"}, {"$ref": "#/components/schemas/workers_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Versions"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.versions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
