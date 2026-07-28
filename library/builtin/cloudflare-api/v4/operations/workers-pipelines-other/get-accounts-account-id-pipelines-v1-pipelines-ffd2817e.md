---
title: List Pipelines
page_id: operation-get-accounts-account-id-pipelines-v1-pipelines-08f5d5ed
path: operations/workers-pipelines-other
description: List/Filter Pipelines in Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/pipelines
operation_ids:
    - getV4AccountsByAccount_idPipelinesV1Pipelines
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Pipelines

`GET /accounts/{account_id}/pipelines/v1/pipelines`

Operation ID: `getV4AccountsByAccount_idPipelinesV1Pipelines`

List/Filter Pipelines in Account.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelinesV1Pipelines", "summary": "List Pipelines", "description": "List/Filter Pipelines in Account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "page", "in": "query", "schema": {"type": "number", "default": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "number"}}, {"name": "name", "in": "query", "description": "Filters pipelines by name (case-insensitive substring).", "schema": {"type": "string", "maxLength": 128, "minLength": 1}}], "responses": {"200": {"description": "Indicates a successfully listed Pipelines.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "array", "items": {"properties": {"created_at": {"type": "string"}, "id": {"description": "Indicates a unique identifier for this pipeline.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string"}, "name": {"description": "Indicates the name of the Pipeline.", "type": "string", "example": "my_pipeline", "maxLength": 128, "minLength": 1}, "sql": {"description": "Specifies SQL for the Pipeline processing flow.", "type": "string", "example": "insert into sink select * from source;"}, "status": {"description": "Indicates the current status of the Pipeline.", "type": "string"}}, "required": ["id", "name", "created_at", "modified_at", "sql", "status"], "type": "object"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Indicates the number of items on current page.", "type": "number", "example": 1}, "page": {"description": "Indicates the current page number.", "type": "number", "example": 0}, "per_page": {"description": "Indicates the number of items per page.", "type": "number", "example": 10}, "total_count": {"description": "Indicates the total number of items.", "type": "number", "example": 1}}, "required": ["count", "page", "per_page", "total_count"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result", "result_info"]}}}}, "4XX": {"description": "Indicates an error in listing Pipelines."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "list-v1", "x-forge-hidden": true}
```
