---
title: Edit Worker
page_id: operation-patch-accounts-account-id-workers-workers-worker-id-131b9105
path: operations/workers
description: Perform a partial update on a Worker, where omitted properties are left unchanged from their current values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/workers/workers/{worker_id}
operation_ids:
    - editWorker
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Worker

`PATCH /accounts/{account_id}/workers/workers/{worker_id}`

Operation ID: `editWorker`

Perform a partial update on a Worker, where omitted properties are left unchanged from their current values.

## Definition

```yaml
{"operationId": "editWorker", "summary": "Edit Worker", "description": "Perform a partial update on a Worker, where omitted properties are left unchanged from their current values.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "worker_id", "in": "path", "required": true, "schema": {"description": "Identifier for the Worker, which can be ID or name.", "type": "string"}}], "requestBody": {"description": "Worker partial update parameters.", "required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_Worker"}, {"type": "object"}]}}, "application/merge-patch+json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_Worker"}, {"type": "object"}]}}}}, "responses": {"200": {"description": "Partially Update Worker success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Worker"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Bad Request - Invalid input data.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10003": "#/components/schemas/workers_ErrorMissingParam", "10016": "#/components/schemas/workers_ErrorWorkerNameInvalid", "10021": "#/components/schemas/workers_ErrorWorkerInvalid", "100102": "#/components/schemas/workers_ErrorWorkerTagLengthLimit", "100132": "#/components/schemas/workers_ErrorWorkerNameSubdomainLengthLimit", "100134": "#/components/schemas/workers_ErrorWorkerTagInvalid", "100308": "#/components/schemas/workers_ErrorWorkerObservabilitySamplingRateInvalid", "100315": "#/components/schemas/workers_ErrorWorkerNamePreviewLengthLimit"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorMissingParam"}, {"$ref": "#/components/schemas/workers_ErrorWorkerNameInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerTagLengthLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerTagInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerNameSubdomainLengthLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerNamePreviewLengthLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerObservabilitySamplingRateInvalid"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "401": {"$ref": "#/components/responses/workers_ErrorAuth"}, "403": {"description": "Forbidden - Insufficient permissions or quota exceeded.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10037": "#/components/schemas/workers_ErrorWorkerLimit", "100103": "#/components/schemas/workers_ErrorWorkerTagLimit"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorWorkerLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerTagLimit"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/workers_ErrorWorkerNotFound"}, "409": {"description": "Conflict - Worker name already exists.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10040": "#/components/schemas/workers_ErrorWorkerNameConflict"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorWorkerNameConflict"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "500": {"$ref": "#/components/responses/workers_ErrorInternalServer"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
