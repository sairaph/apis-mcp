---
title: Fetch Usage Model
page_id: operation-get-accounts-account-id-workers-scripts-script-name-usage-model-e7b50897
path: operations/worker-script
description: Fetches the Usage Model for a given Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/usage-model
operation_ids:
    - worker-script-fetch-usage-model
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch Usage Model

`GET /accounts/{account_id}/workers/scripts/{script_name}/usage-model`

Operation ID: `worker-script-fetch-usage-model`

Fetches the Usage Model for a given Worker.

## Definition

```yaml
{"operationId": "worker-script-fetch-usage-model", "summary": "Fetch Usage Model", "description": "Fetches the Usage Model for a given Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "Fetch Usage Model response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_usage-model-response"}}}}, "4XX": {"description": "Fetch Usage Model response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.usage-model", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
