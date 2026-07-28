---
title: Get repository configuration autofill
page_id: operation-get-accounts-account-id-builds-repos-provider-type-provider-account-id-r-cc50960a
path: operations/github-integration
description: Analyze repository for automatic configuration detection
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/repos/{provider_type}/{provider_account_id}/{repo_id}/config_autofill
operation_ids:
    - getWorkerConfigAutofill
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get repository configuration autofill

`GET /accounts/{account_id}/builds/repos/{provider_type}/{provider_account_id}/{repo_id}/config_autofill`

Operation ID: `getWorkerConfigAutofill`

Analyze repository for automatic configuration detection

## Definition

```yaml
{"operationId": "getWorkerConfigAutofill", "summary": "Get repository configuration autofill", "description": "Analyze repository for automatic configuration detection", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"name": "provider_type", "in": "path", "description": "SCM provider type", "required": true, "schema": {"$ref": "#/components/schemas/builds_SCMProviderType"}}, {"name": "provider_account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/builds_provider_account_id"}}, {"name": "repo_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/builds_repo_id"}}, {"name": "branch", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/builds_branch"}}, {"name": "root_directory", "in": "query", "schema": {"$ref": "#/components/schemas/builds_root_directory"}}], "responses": {"200": {"description": "Configuration autofill data retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_ConfigAutofillResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["GitHub Integration"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.repos.config-autofill", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
