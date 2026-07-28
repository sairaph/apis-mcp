---
title: Get Sink Details
page_id: operation-get-accounts-account-id-pipelines-v1-sinks-sink-id-3e4dc9fb
path: operations/workers-pipelines-other
description: Get Sink Details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/sinks/{sink_id}
operation_ids:
    - getV4AccountsByAccount_idPipelinesV1SinksBySink_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Sink Details

`GET /accounts/{account_id}/pipelines/v1/sinks/{sink_id}`

Operation ID: `getV4AccountsByAccount_idPipelinesV1SinksBySink_id`

Get Sink Details.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelinesV1SinksBySink_id", "summary": "Get Sink Details", "description": "Get Sink Details.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "sink_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-sink-id"}}], "responses": {"200": {"description": "Indicates that Sink was retrieved.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"config": {"description": "Defines the configuration of the R2 Sink.", "oneOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_r2TablePublic"}, {"$ref": "#/components/schemas/cloudflare-pipelines_r2_data_catalogTablePublic"}]}, "created_at": {"type": "string", "format": "date-time"}, "format": {"$ref": "#/components/schemas/cloudflare-pipelines_Format"}, "id": {"description": "Indicates a unique identifier for this sink.", "type": "string", "example": "01234567890123457689012345678901"}, "modified_at": {"type": "string", "format": "date-time"}, "name": {"description": "Defines the name of the Sink.", "type": "string", "example": "my_sink", "maxLength": 128, "minLength": 1}, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_ConnectionSchema"}, "type": {"description": "Specifies the type of sink.", "type": "string", "example": "r2", "enum": ["r2", "r2_data_catalog"]}}, "required": ["id", "name", "created_at", "modified_at", "type"]}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in listing Sinks."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines.sinks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
