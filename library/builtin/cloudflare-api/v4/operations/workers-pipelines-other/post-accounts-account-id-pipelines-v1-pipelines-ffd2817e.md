---
title: Create Pipeline
page_id: operation-post-accounts-account-id-pipelines-v1-pipelines-9ff11f0d
path: operations/workers-pipelines-other
description: Create a new Pipeline.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/pipelines
operation_ids:
    - postV4AccountsByAccount_idPipelinesV1Pipelines
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Pipeline

`POST /accounts/{account_id}/pipelines/v1/pipelines`

Operation ID: `postV4AccountsByAccount_idPipelinesV1Pipelines`

Create a new Pipeline.

## Definition

```yaml
{"operationId": "postV4AccountsByAccount_idPipelinesV1Pipelines", "summary": "Create Pipeline", "description": "Create a new Pipeline.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"description": "Specifies the name of the Pipeline.", "type": "string", "example": "my_pipeline", "maxLength": 128, "minLength": 1}, "sql": {"description": "Specifies SQL for the Pipeline processing flow.", "type": "string", "example": "insert into sink select * from source;"}}, "required": ["name", "sql"]}}}}, "responses": {"200": {"description": "Indicates a successfully created Pipeline.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string"}, "id": {"description": "Indicates a unique identifier for this pipeline.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string"}, "name": {"description": "Indicates the name of the Pipeline.", "type": "string", "example": "my_pipeline", "maxLength": 128, "minLength": 1}, "sql": {"description": "Specifies SQL for the Pipeline processing flow.", "type": "string", "example": "insert into sink select * from source;"}, "status": {"description": "Indicates the current status of the Pipeline.", "type": "string"}}, "required": ["id", "name", "created_at", "modified_at", "sql", "status"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in creating a Pipeline."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "create-v1", "x-forge-hidden": true}
```
