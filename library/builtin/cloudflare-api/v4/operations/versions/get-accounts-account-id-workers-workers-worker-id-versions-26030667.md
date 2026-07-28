---
title: List Versions
page_id: operation-get-accounts-account-id-workers-workers-worker-id-versions-3b5c6495
path: operations/versions
description: List all versions for a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}/versions
operation_ids:
    - listWorkerVersions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Versions

`GET /accounts/{account_id}/workers/workers/{worker_id}/versions`

Operation ID: `listWorkerVersions`

List all versions for a Worker.

## Definition

```yaml
{"operationId": "listWorkerVersions", "summary": "List Versions", "description": "List all versions for a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}, {"name": "page", "in": "query", "description": "Current page.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Items per-page.", "schema": {"type": "integer", "default": 10, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List versions success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_Version"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List versions failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Versions"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers.versions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
