---
title: '[DEPRECATED] Delete Pipeline'
page_id: operation-delete-accounts-account-id-pipelines-pipeline-name-5bc16f18
path: operations/workers-pipelines-other
description: '[DEPRECATED] Delete a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pipelines/{pipeline_name}
operation_ids:
    - deleteV4AccountsByAccount_idPipelinesByPipeline_name_deprecated
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# [DEPRECATED] Delete Pipeline

`DELETE /accounts/{account_id}/pipelines/{pipeline_name}`

Operation ID: `deleteV4AccountsByAccount_idPipelinesByPipeline_name_deprecated`

[DEPRECATED] Delete a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.

## Definition

```yaml
{"operationId": "deleteV4AccountsByAccount_idPipelinesByPipeline_name_deprecated", "summary": "[DEPRECATED] Delete Pipeline", "description": "[DEPRECATED] Delete a pipeline. Use the new /pipelines/v1/pipelines endpoint instead.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudflare-pipelines_workers-pipelines-account-id"}}, {"name": "pipeline_name", "in": "path", "required": true, "schema": {"description": "Defines the name of the pipeline.", "type": "string", "example": "sample_pipeline", "maxLength": 128, "minLength": 1}}], "responses": {"200": {"description": "[DEPRECATED] Indicates a successfully deleted pipeline."}, "4XX": {"description": "Indicates an error in deleting a pipeline."}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["workers_pipelines_other"], "x-api-token-group": ["Pipelines Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pipelines", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
