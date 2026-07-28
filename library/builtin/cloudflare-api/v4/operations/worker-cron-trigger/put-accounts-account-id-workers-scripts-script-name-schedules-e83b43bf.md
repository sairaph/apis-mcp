---
title: Update Cron Triggers
page_id: operation-put-accounts-account-id-workers-scripts-script-name-schedules-de6ec172
path: operations/worker-cron-trigger
description: Updates Cron Triggers for a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/schedules
operation_ids:
    - worker-cron-trigger-update-cron-triggers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Cron Triggers

`PUT /accounts/{account_id}/workers/scripts/{script_name}/schedules`

Operation ID: `worker-cron-trigger-update-cron-triggers`

Updates Cron Triggers for a Worker.

## Definition

```yaml
{"operationId": "worker-cron-trigger-update-cron-triggers", "summary": "Update Cron Triggers", "description": "Updates Cron Triggers for a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/workers_schedule"}}}}}, "responses": {"200": {"description": "Update Cron Triggers response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"schedules": {"type": "array", "items": {"$ref": "#/components/schemas/workers_schedule"}}}, "required": ["schedules"]}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Update Cron Triggers response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Cron Trigger"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.schedules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
