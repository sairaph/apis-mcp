---
title: Update Stream
page_id: operation-patch-accounts-account-id-pipelines-v1-streams-stream-id-531c715d
path: operations/workers-pipelines-other
description: Update a Stream.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/streams/{stream_id}
operation_ids:
    - patchV4AccountsByAccount_idPipelinesV1StreamsByStream_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Stream

`PATCH /accounts/{account_id}/pipelines/v1/streams/{stream_id}`

Operation ID: `patchV4AccountsByAccount_idPipelinesV1StreamsByStream_id`

Update a Stream.

## Definition

```yaml
{"operationId": "patchV4AccountsByAccount_idPipelinesV1StreamsByStream_id", "summary": "Update Stream", "description": "Update a Stream.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "stream_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-stream-id"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"http": {"type": "object", "properties": {"authentication": {"description": "Indicates that authentication is required for the HTTP endpoint.", "type": "boolean", "example": false}, "cors": {"description": "Specifies the CORS options for the HTTP endpoint.", "type": "object", "example": {}, "properties": {"origins": {"type": "array", "items": {"type": "string"}, "maxItems": 5}}}, "enabled": {"description": "Indicates that the HTTP endpoint is enabled.", "type": "boolean", "example": true}}, "required": ["enabled", "authentication"]}, "worker_binding": {"type": "object", "properties": {"enabled": {"description": "Indicates that the worker binding is enabled.", "type": "boolean", "example": true}}, "required": ["enabled"]}}}}}}, "responses": {"200": {"description": "Indicates a successfully updated Stream.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string", "format": "date-time"}, "endpoint": {"description": "Indicates the endpoint URL of this stream.", "type": "string", "format": "uri", "example": "https://01234567890123457689012345678901.ingest.cloudflare.com/v1"}, "format": {"$ref": "#/components/schemas/cloudflare-pipelines_Format"}, "http": {"type": "object", "properties": {"authentication": {"description": "Indicates that authentication is required for the HTTP endpoint.", "type": "boolean", "example": false}, "cors": {"description": "Specifies the CORS options for the HTTP endpoint.", "type": "object", "example": {}, "properties": {"origins": {"type": "array", "items": {"type": "string"}, "maxItems": 5}}}, "enabled": {"description": "Indicates that the HTTP endpoint is enabled.", "type": "boolean", "example": true}}, "required": ["enabled", "authentication"]}, "id": {"description": "Indicates a unique identifier for this stream.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"description": "Indicates the name of the Stream.", "type": "string", "example": "my_stream", "maxLength": 128, "minLength": 1}, "version": {"description": "Indicates the current version of this stream.", "type": "integer", "example": 3}, "worker_binding": {"type": "object", "properties": {"enabled": {"description": "Indicates that the worker binding is enabled.", "type": "boolean", "example": true}}, "required": ["enabled"]}}, "required": ["id", "name", "version", "created_at", "modified_at", "http", "worker_binding"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in creating a Stream."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.streams", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
