---
title: Delete build token
page_id: operation-delete-accounts-account-id-builds-tokens-build-token-uuid-4c3dbf15
path: operations/build-tokens
description: Remove a build authentication token
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/builds/tokens/{build_token_uuid}
operation_ids:
    - deleteBuildToken
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete build token

`DELETE /accounts/{account_id}/builds/tokens/{build_token_uuid}`

Operation ID: `deleteBuildToken`

Remove a build authentication token

## Definition

```yaml
{"operationId": "deleteBuildToken", "summary": "Delete build token", "description": "Remove a build authentication token", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"name": "build_token_uuid", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/builds_build_token_uuid"}}], "responses": {"200": {"$ref": "#/components/responses/builds_SuccessEmpty"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Build Tokens"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.tokens", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
