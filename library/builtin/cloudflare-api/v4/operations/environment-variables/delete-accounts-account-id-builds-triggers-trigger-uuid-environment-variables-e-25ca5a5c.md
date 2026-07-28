---
title: Delete environment variable
page_id: operation-delete-accounts-account-id-builds-triggers-trigger-uuid-environment-vari-d8c36aa9
path: operations/environment-variables
description: Remove a specific environment variable from a trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables/{environment_variable_key}
operation_ids:
    - deleteEnvironmentVariable
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete environment variable

`DELETE /accounts/{account_id}/builds/triggers/{trigger_uuid}/environment_variables/{environment_variable_key}`

Operation ID: `deleteEnvironmentVariable`

Remove a specific environment variable from a trigger

## Definition

```yaml
{"operationId": "deleteEnvironmentVariable", "summary": "Delete environment variable", "description": "Remove a specific environment variable from a trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_TriggerUuid"}, {"name": "environment_variable_key", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/builds_environment_variable_key"}}], "responses": {"200": {"$ref": "#/components/responses/builds_SuccessEmpty"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Environment Variables"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers.environment-variables", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
