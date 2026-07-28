---
title: Update Usage Model
page_id: operation-put-accounts-account-id-workers-scripts-script-name-usage-model-c40ddac1
path: operations/worker-script
description: Updates the Usage Model for a given Worker. Requires a Workers Paid subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/usage-model
operation_ids:
    - worker-script-update-usage-model
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Usage Model

`PUT /accounts/{account_id}/workers/scripts/{script_name}/usage-model`

Operation ID: `worker-script-update-usage-model`

Updates the Usage Model for a given Worker. Requires a Workers Paid subscription.

## Definition

```yaml
{"operationId": "worker-script-update-usage-model", "summary": "Update Usage Model", "description": "Updates the Usage Model for a given Worker. Requires a Workers Paid subscription.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"usage_model": {"$ref": "#/components/schemas/workers_usage_model"}, "user_limits": {"$ref": "#/components/schemas/workers_user_limits"}}}}}}, "responses": {"200": {"description": "Update Usage Model response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_usage-model-response"}}}}, "4XX": {"description": "Update Usage Model response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.usage-model", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
