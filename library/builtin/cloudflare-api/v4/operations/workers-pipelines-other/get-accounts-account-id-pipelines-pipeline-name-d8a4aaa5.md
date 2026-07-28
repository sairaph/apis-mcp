---
title: '[DEPRECATED] Get Pipeline'
page_id: operation-get-accounts-account-id-pipelines-pipeline-name-40b978bd
path: operations/workers-pipelines-other
description: '[DEPRECATED] Get configuration of a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pipelines/{pipeline_name}
operation_ids:
    - getV4AccountsByAccount_idPipelinesByPipeline_name_deprecated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# [DEPRECATED] Get Pipeline

`GET /accounts/{account_id}/pipelines/{pipeline_name}`

Operation ID: `getV4AccountsByAccount_idPipelinesByPipeline_name_deprecated`

[DEPRECATED] Get configuration of a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.

## Definition

```yaml
{"operationId": "getV4AccountsByAccount_idPipelinesByPipeline_name_deprecated", "summary": "[DEPRECATED] Get Pipeline", "description": "[DEPRECATED] Get configuration of a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "pipeline_name", "in": "path", "required": true, "schema": {"description": "Defines the name of the pipeline.", "type": "string", "example": "sample_pipeline", "maxLength": 128, "minLength": 1}}], "responses": {"200": {"description": "[DEPRECATED] Describes the configuration of a pipeline.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-pipeline"}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "404": {"description": "Indicates that the pipeline was not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number"}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object"}, "example": []}, "results": {"type": "object", "nullable": true, "x-stainless-empty-object": true}, "success": {"example": false, "allOf": [{"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}]}}, "required": ["success", "results", "errors"]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write", "Pipelines Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
