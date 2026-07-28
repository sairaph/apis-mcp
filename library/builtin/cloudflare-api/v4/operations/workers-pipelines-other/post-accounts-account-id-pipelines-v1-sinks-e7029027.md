---
title: Create Sink
page_id: operation-post-accounts-account-id-pipelines-v1-sinks-19b528af
path: operations/workers-pipelines-other
description: Create a new Sink.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/sinks
operation_ids:
    - postV4AccountsByAccount_idPipelinesV1Sinks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Sink

`POST /accounts/{account_id}/pipelines/v1/sinks`

Operation ID: `postV4AccountsByAccount_idPipelinesV1Sinks`

Create a new Sink.

## Definition

```yaml
{"operationId": "postV4AccountsByAccount_idPipelinesV1Sinks", "summary": "Create Sink", "description": "Create a new Sink.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"description": "Defines the configuration of the R2 Sink.", "oneOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_r2Table"}, {"$ref": "#/components/schemas/cloudflare-pipelines_r2_data_catalogTable"}]}, "format": {"$ref": "#/components/schemas/cloudflare-pipelines_Format"}, "name": {"description": "Defines the name of the Sink.", "type": "string", "example": "my_sink", "maxLength": 128, "minLength": 1}, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_ConnectionSchema"}, "type": {"description": "Specifies the type of sink.", "type": "string", "example": "r2", "enum": ["r2", "r2_data_catalog"]}}, "required": ["name", "type"]}}}}, "responses": {"200": {"description": "Indicates a successfully created Sink.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"config": {"oneOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_r2Table"}, {"$ref": "#/components/schemas/cloudflare-pipelines_r2_data_catalogTable"}]}, "created_at": {"type": "string", "format": "date-time"}, "format": {"$ref": "#/components/schemas/cloudflare-pipelines_Format"}, "id": {"description": "Indicates a unique identifier for this sink.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"description": "Defines the name of the Sink.", "type": "string", "example": "my_sink", "maxLength": 128, "minLength": 1}, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_ConnectionSchema"}, "type": {"description": "Specifies the type of sink.", "type": "string", "example": "r2", "enum": ["r2", "r2_data_catalog"]}}, "required": ["id", "name", "created_at", "modified_at", "type"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in creating a Sink."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.sinks", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
