---
title: List environment variables
page_id: operation-get-accounts-account-id-builds-triggers-trigger-uuid-environment-variabl-fdcf9afc
path: operations/environment-variables
description: Get all environment variables for a trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables
operation_ids:
    - listEnvironmentVariables
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List environment variables

`GET /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables`

Operation ID: `listEnvironmentVariables`

Get all environment variables for a trigger

## Definition

```yaml
{"operationId": "listEnvironmentVariables", "summary": "List environment variables", "description": "Get all environment variables for a trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_TriggerUuid"}], "responses": {"200": {"description": "Environment variables retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_EnvironmentVariablesResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Environment Variables"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers.environment-variables", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
