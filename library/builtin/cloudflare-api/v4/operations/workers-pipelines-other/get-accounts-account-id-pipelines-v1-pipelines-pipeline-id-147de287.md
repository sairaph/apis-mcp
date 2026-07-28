---
title: Get Pipeline Details
page_id: operation-get-accounts-account-id-pipelines-v1-pipelines-pipeline-id-0fc7dba6
path: operations/workers-pipelines-other
description: Get Pipelines Details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/pipelines/{pipeline_id}
operation_ids:
    - getV4AccountsByAccount_idPipelinesV1PipelinesByPipeline_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Pipeline Details

`GET /accounts/{account_id}/pipelines/v1/pipelines/{pipeline_id}`

Operation ID: `getV4AccountsByAccount_idPipelinesV1PipelinesByPipeline_id`

Get Pipelines Details.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelinesV1PipelinesByPipeline_id", "summary": "Get Pipeline Details", "description": "Get Pipelines Details.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "pipeline_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-pipeline-id"}}], "responses": {"200": {"description": "Indicates a successfully retrieved Pipeline.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string"}, "failure_reason": {"description": "Indicates the reason for the failure of the Pipeline.", "type": "string"}, "id": {"description": "Indicates a unique identifier for this pipeline.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string"}, "name": {"description": "Indicates the name of the Pipeline.", "type": "string", "example": "my_pipeline", "maxLength": 128, "minLength": 1}, "sql": {"description": "Specifies SQL for the Pipeline processing flow.", "type": "string", "example": "insert into sink select * from source;"}, "status": {"description": "Indicates the current status of the Pipeline.", "type": "string"}, "tables": {"description": "List of streams and sinks used by this pipeline.", "type": "array", "items": {"properties": {"id": {"description": "Unique identifier for the connection (stream or sink).", "type": "string", "example": "1c9200d5872c018bb34e93e2cd8a438e"}, "latest": {"description": "Latest available version of the connection.", "type": "integer", "example": 5}, "name": {"description": "Name of the connection.", "type": "string", "example": "my_table", "maxLength": 128, "minLength": 1}, "type": {"description": "Type of the connection.", "type": "string", "example": "stream", "enum": ["stream", "sink"]}, "version": {"description": "Current version of the connection used by this pipeline.", "type": "integer", "example": 4}}, "required": ["id", "version", "latest", "type", "name"], "type": "object"}}}, "required": ["id", "name", "created_at", "modified_at", "sql", "status", "tables"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in retrieving Pipelines."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "get-v1", "x-forge-hidden": true}
```
