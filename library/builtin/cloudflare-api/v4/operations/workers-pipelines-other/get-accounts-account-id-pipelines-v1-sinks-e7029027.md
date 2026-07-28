---
title: List Sinks
page_id: operation-get-accounts-account-id-pipelines-v1-sinks-a6e3e926
path: operations/workers-pipelines-other
description: List/Filter Sinks in Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/sinks
operation_ids:
    - getV4AccountsByAccount_idPipelinesV1Sinks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Sinks

`GET /accounts/{account_id}/pipelines/v1/sinks`

Operation ID: `getV4AccountsByAccount_idPipelinesV1Sinks`

List/Filter Sinks in Account.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelinesV1Sinks", "summary": "List Sinks", "description": "List/Filter Sinks in Account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "pipeline_id", "in": "query", "schema": {"type": "string"}}, {"name": "name", "in": "query", "description": "Filters sinks by name (case-insensitive substring).", "schema": {"type": "string", "maxLength": 128, "minLength": 1}}, {"name": "page", "in": "query", "schema": {"type": "number", "default": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "number"}}], "responses": {"200": {"description": "Indicates successfully listed Sinks.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"config": {"description": "Defines the configuration of the R2 Sink.", "oneOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_r2TablePublic"}, {"$ref": "#/components/schemas/cloudflare-pipelines_r2_data_catalogTablePublic"}]}, "created_at": {"type": "string", "format": "date-time"}, "format": {"$ref": "#/components/schemas/cloudflare-pipelines_Format"}, "id": {"description": "Indicates a unique identifier for this sink.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"description": "Defines the name of the Sink.", "type": "string", "example": "my_sink", "maxLength": 128, "minLength": 1}, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_ConnectionSchema"}, "type": {"description": "Specifies the type of sink.", "type": "string", "example": "r2", "enum": ["r2", "r2_data_catalog"]}}, "required": ["id", "name", "created_at", "modified_at", "type"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Indicates the number of items on current page.", "type": "number", "example": 1}, "page": {"description": "Indicates the current page number.", "type": "number", "example": 0}, "per_page": {"description": "Indicates the number of items per page.", "type": "number", "example": 10}, "total_count": {"description": "Indicates the total number of items.", "type": "number", "example": 1}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result", "result_info"]}}}}, "4XX": {"description": "Indicates an error in listing Sinks."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.sinks", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
