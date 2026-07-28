---
title: Delete Worker
page_id: operation-delete-accounts-account-id-workers-workers-worker-id-e113e315
path: operations/workers
description: Delete a Worker and all its associated resources (versions, deployments, etc.).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}
operation_ids:
    - deleteWorker
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Worker

`DELETE /accounts/{account_id}/workers/workers/{worker_id}`

Operation ID: `deleteWorker`

Delete a Worker and all its associated resources (versions, deployments, etc.).

## Definition

```yaml
{"operationId": "deleteWorker", "summary": "Delete Worker", "description": "Delete a Worker and all its associated resources (versions, deployments, etc.).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}], "responses": {"200": {"description": "Delete Worker success.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common"}}}}, "400": {"description": "Bad Request - Missing or invalid parameters.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10003": "#/components/schemas/workers_ErrorMissingParam"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorMissingParam"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "401": {"$ref": "#/components/responses/workers_ErrorAuth"}, "404": {"$ref": "#/components/responses/workers_ErrorWorkerNotFound"}, "500": {"$ref": "#/components/responses/workers_ErrorInternalServer"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
