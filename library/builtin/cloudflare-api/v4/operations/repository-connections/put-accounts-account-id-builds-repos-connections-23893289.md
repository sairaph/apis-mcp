---
title: Create or update repository connection
page_id: operation-put-accounts-account-id-builds-repos-connections-f3038fb7
path: operations/repository-connections
description: Upsert a repository connection for CI/CD integration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/builds/repos/connections
operation_ids:
    - upsertRepoConnection
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create or update repository connection

`PUT /accounts/{account_id}/builds/repos/connections`

Operation ID: `upsertRepoConnection`

Upsert a repository connection for CI/CD integration

## Definition

```yaml
{"operationId": "upsertRepoConnection", "summary": "Create or update repository connection", "description": "Upsert a repository connection for CI/CD integration", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_UpsertRepoConnectionRequest"}}}}, "responses": {"200": {"description": "Repository connection upserted successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_UpsertRepoConnectionResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Repository Connections"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.repos.connections", "x-fern-sdk-method-name": "upsert", "x-forge-hidden": true}
```
