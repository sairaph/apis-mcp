---
title: Get Worker
page_id: operation-get-accounts-account-id-workers-workers-worker-id-2c37bd22
path: operations/workers
description: Get details about a specific Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}
operation_ids:
    - getWorker
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Worker

`GET /accounts/{account_id}/workers/workers/{worker_id}`

Operation ID: `getWorker`

Get details about a specific Worker.

## Definition

```yaml
{"operationId": "getWorker", "summary": "Get Worker", "description": "Get details about a specific Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}], "responses": {"200": {"description": "Get Worker success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Worker"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Bad Request - Missing or invalid parameters.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10003": "#/components/schemas/workers_ErrorMissingParam"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorMissingParam"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/workers_ErrorWorkerNotFound"}, "500": {"$ref": "#/components/responses/workers_ErrorInternalServer"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
