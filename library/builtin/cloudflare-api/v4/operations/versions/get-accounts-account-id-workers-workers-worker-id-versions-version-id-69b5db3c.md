---
title: Get Version
page_id: operation-get-accounts-account-id-workers-workers-worker-id-versions-version-id-75254357
path: operations/versions
description: Get details about a specific version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}
operation_ids:
    - getWorkerVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Version

`GET /accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}`

Operation ID: `getWorkerVersion`

Get details about a specific version.

## Definition

```yaml
{"operationId": "getWorkerVersion", "summary": "Get Version", "description": "Get details about a specific version.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}, {"name": "version_id", "in": "path", "required": true, "schema": {"description": "Identifier for the version, which can be a UUID, a UUID prefix (minimum length 8), or the literal \"latest\" to operate on the most recently created version.", "type": "string"}}, {"name": "include", "in": "query", "schema": {"description": "Whether to include the `modules` property of the version in the response, which contains code and sourcemap content and may add several megabytes to the response size.", "type": "string", "enum": ["modules"]}}], "responses": {"200": {"description": "Get version success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Version"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get version failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Versions"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers.versions", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
