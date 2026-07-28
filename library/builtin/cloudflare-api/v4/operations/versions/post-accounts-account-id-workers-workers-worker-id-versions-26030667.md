---
title: Create Version
page_id: operation-post-accounts-account-id-workers-workers-worker-id-versions-81406923
path: operations/versions
description: Create a new version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}/versions
operation_ids:
    - createWorkerVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Version

`POST /accounts/{account_id}/workers/workers/{worker_id}/versions`

Operation ID: `createWorkerVersion`

Create a new version.

## Definition

```yaml
{"operationId": "createWorkerVersion", "summary": "Create Version", "description": "Create a new version.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}, {"name": "deploy", "in": "query", "schema": {"description": "If true, a deployment will be created that sends 100% of traffic to the new version.", "type": "boolean"}}], "requestBody": {"description": "Version creation parameters.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_Version"}}}}, "responses": {"200": {"description": "Create version success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Version"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create version failure. When the upload uses the declarative\n`exports` field and one or more entries fail reconciliation,\nthe response is the exports reconciliation error envelope\n(error code 100402) with per-class detail in\n`errors[].meta.details`.\n", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"$ref": "#/components/schemas/workers_exports_reconciliation_error_response"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Versions"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers.versions", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
