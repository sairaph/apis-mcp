---
title: Create Worker
page_id: operation-post-accounts-account-id-workers-workers-26217f8f
path: operations/workers
description: Create a new Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/workers
operation_ids:
    - createWorker
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Worker

`POST /accounts/{account_id}/workers/workers`

Operation ID: `createWorker`

Create a new Worker.

## Definition

```yaml
{"operationId": "createWorker", "summary": "Create Worker", "description": "Create a new Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"description": "Worker creation parameters.", "required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_Worker"}, {"required": ["name"], "type": "object"}]}}}}, "responses": {"200": {"description": "Create Worker success.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Worker"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Bad Request - Invalid input data.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10016": "#/components/schemas/workers_ErrorWorkerNameInvalid", "10021": "#/components/schemas/workers_ErrorWorkerInvalid", "100102": "#/components/schemas/workers_ErrorWorkerTagLengthLimit", "100132": "#/components/schemas/workers_ErrorWorkerNameSubdomainLengthLimit", "100134": "#/components/schemas/workers_ErrorWorkerTagInvalid", "100308": "#/components/schemas/workers_ErrorWorkerObservabilitySamplingRateInvalid", "100315": "#/components/schemas/workers_ErrorWorkerNamePreviewLengthLimit"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorWorkerNameInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerTagLengthLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerNameSubdomainLengthLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerTagInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerObservabilitySamplingRateInvalid"}, {"$ref": "#/components/schemas/workers_ErrorWorkerNamePreviewLengthLimit"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "401": {"$ref": "#/components/responses/workers_ErrorAuth"}, "403": {"description": "Forbidden - Access denied or limit exceeded.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10037": "#/components/schemas/workers_ErrorWorkerLimit", "100103": "#/components/schemas/workers_ErrorWorkerTagLimit"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorWorkerLimit"}, {"$ref": "#/components/schemas/workers_ErrorWorkerTagLimit"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "409": {"description": "Conflict - Resource already exists.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common-failure"}, {"properties": {"errors": {"type": "array", "items": {"discriminator": {"mapping": {"10040": "#/components/schemas/workers_ErrorWorkerNameConflict"}, "propertyName": "code"}, "oneOf": [{"$ref": "#/components/schemas/workers_ErrorWorkerNameConflict"}], "type": "object"}}}, "required": ["errors"], "type": "object"}]}}}}, "500": {"$ref": "#/components/responses/workers_ErrorInternalServer"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.beta.workers", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
