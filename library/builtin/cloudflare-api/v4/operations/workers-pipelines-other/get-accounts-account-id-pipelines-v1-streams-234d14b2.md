---
title: List Streams
page_id: operation-get-accounts-account-id-pipelines-v1-streams-406a2308
path: operations/workers-pipelines-other
description: List/Filter Streams in Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/streams
operation_ids:
    - getV4AccountsByAccount_idPipelinesV1Streams
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Streams

`GET /accounts/{account_id}/pipelines/v1/streams`

Operation ID: `getV4AccountsByAccount_idPipelinesV1Streams`

List/Filter Streams in Account.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelinesV1Streams", "summary": "List Streams", "description": "List/Filter Streams in Account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "pipeline_id", "in": "query", "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-pipeline-id"}}, {"name": "name", "in": "query", "description": "Filters streams by name (case-insensitive substring).", "schema": {"type": "string", "maxLength": 128, "minLength": 1}}, {"name": "page", "in": "query", "schema": {"type": "number", "default": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "number"}}], "responses": {"200": {"description": "Indicates a successfully created Stream.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"type": "string", "format": "date-time"}, "endpoint": {"description": "Indicates the endpoint URL of this stream.", "type": "string", "format": "uri", "example": "https://01234567890123457689012345678901.ingest.cloudflare.com/v1"}, "format": {"$ref": "#/components/schemas/cloudflare-pipelines_Format"}, "http": {"type": "object", "properties": {"authentication": {"description": "Indicates that authentication is required for the HTTP endpoint.", "type": "boolean", "example": false}, "cors": {"description": "Specifies the CORS options for the HTTP endpoint.", "type": "object", "example": {}, "properties": {"origins": {"type": "array", "items": {"type": "string"}, "maxItems": 5}}}, "enabled": {"description": "Indicates that the HTTP endpoint is enabled.", "type": "boolean", "example": true}}, "required": ["enabled", "authentication"]}, "id": {"description": "Indicates a unique identifier for this stream.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"description": "Indicates the name of the Stream.", "type": "string", "example": "my_stream", "maxLength": 128, "minLength": 1}, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_ConnectionSchema"}, "version": {"description": "Indicates the current version of this stream.", "type": "integer", "example": 3}, "worker_binding": {"type": "object", "properties": {"enabled": {"description": "Indicates that the worker binding is enabled.", "type": "boolean", "example": true}}, "required": ["enabled"]}}, "required": ["id", "name", "version", "created_at", "modified_at", "http", "worker_binding"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Indicates the number of items on current page.", "type": "number", "example": 1}, "page": {"description": "Indicates the current page number.", "type": "number", "example": 0}, "per_page": {"description": "Indicates the number of items per page.", "type": "number", "example": 10}, "total_count": {"description": "Indicates the total number of items.", "type": "number", "example": 1}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result", "result_info"]}}}}, "4XX": {"description": "Indicates an error in listing Streams."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.streams", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
