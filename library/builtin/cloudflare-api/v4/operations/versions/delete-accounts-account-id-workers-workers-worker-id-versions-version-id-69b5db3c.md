---
title: Delete Version
page_id: operation-delete-accounts-account-id-workers-workers-worker-id-versions-version-id-cb73156a
path: operations/versions
description: Delete a version.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}
operation_ids:
    - deleteWorkerVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Version

`DELETE /accounts/{account_id}/workers/workers/{worker_id}/versions/{version_id}`

Operation ID: `deleteWorkerVersion`

Delete a version.

## Definition

```yaml
{"operationId": "deleteWorkerVersion", "summary": "Delete Version", "description": "Delete a version.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}, {"name": "version_id", "in": "path", "required": true, "schema": {"description": "Identifier for the version, which can be a UUID, a UUID prefix (minimum length 8), or the literal \"latest\" to operate on the most recently created version.", "type": "string"}}], "responses": {"200": {"description": "Delete version success.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common"}}}}, "4XX": {"description": "Delete version failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Versions"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers.versions", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
