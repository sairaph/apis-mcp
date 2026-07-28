---
title: Delete Pipelines
page_id: operation-delete-accounts-account-id-pipelines-v1-pipelines-pipeline-id-71c7e5dd
path: operations/workers-pipelines-other
description: Delete Pipeline in Account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pipelines/v1/pipelines/{pipeline_id}
operation_ids:
    - deleteV4AccountsByAccount_idPipelinesV1PipelinesByPipeline_id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Pipelines

`DELETE /accounts/{account_id}/pipelines/v1/pipelines/{pipeline_id}`

Operation ID: `deleteV4AccountsByAccount_idPipelinesV1PipelinesByPipeline_id`

Delete Pipeline in Account.

## Definition

```yaml
{"operationId": "deleteV4AccountsByAccount_idPipelinesV1PipelinesByPipeline_id", "summary": "Delete Pipelines", "description": "Delete Pipeline in Account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "pipeline_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-pipeline-id"}}], "responses": {"200": {"description": "Indicates a successfully deleted Pipeline.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object"}, "success": {"$ref": "#/components/schemas/cloudflare-pipelines_worker-pipelines-common-success"}}, "required": ["success", "result"]}}}}, "4XX": {"description": "Indicates an error in listing Pipelines."}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "delete-v1", "x-forge-hidden": true}
```
