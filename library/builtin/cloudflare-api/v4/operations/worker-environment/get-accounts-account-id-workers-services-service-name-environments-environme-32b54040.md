---
title: Get Script Settings
page_id: operation-get-accounts-account-id-workers-services-service-name-environments-envir-0419015b
path: operations/worker-environment
description: Get script settings from a worker with an environment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/services/{service_name}/environments/{environment_name}/settings
operation_ids:
    - worker-script-environment-get-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Script Settings

`GET /accounts/{account_id}/workers/services/{service_name}/environments/{environment_name}/settings`

Operation ID: `worker-script-environment-get-settings`

Get script settings from a worker with an environment.

## Definition

```yaml
{"operationId": "worker-script-environment-get-settings", "summary": "Get Script Settings", "description": "Get script settings from a worker with an environment.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "service_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_service"}}, {"name": "environment_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_environment"}}], "responses": {"200": {"description": "Fetch script settings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-settings-response"}}}}, "4XX": {"description": "Fetch script settings failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Environment"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.services.environments.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
