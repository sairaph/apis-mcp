---
title: Delete repository connection
page_id: operation-delete-accounts-account-id-builds-repos-connections-repo-connection-uuid-4fc59b92
path: operations/repository-connections
description: Remove a repository connection
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/builds/repos/connections/{repo_connection_uuid}
operation_ids:
    - deleteRepoConnection
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete repository connection

`DELETE /accounts/{account_id}/builds/repos/connections/{repo_connection_uuid}`

Operation ID: `deleteRepoConnection`

Remove a repository connection

## Definition

```yaml
{"operationId": "deleteRepoConnection", "summary": "Delete repository connection", "description": "Remove a repository connection", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_RepoConnectionUuid"}], "responses": {"200": {"$ref": "#/components/responses/builds_SuccessEmpty"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Repository Connections"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.repos.connections", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
