---
title: Cancel build
page_id: operation-put-accounts-account-id-builds-builds-build-uuid-cancel-c0420288
path: operations/builds
description: Cancel a running or queued build
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/builds/builds/{build_uuid}/cancel
operation_ids:
    - cancelBuildByUuid
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Cancel build

`PUT /accounts/{account_id}/builds/builds/{build_uuid}/cancel`

Operation ID: `cancelBuildByUuid`

Cancel a running or queued build

## Definition

```yaml
{"operationId": "cancelBuildByUuid", "summary": "Cancel build", "description": "Cancel a running or queued build", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_BuildUuid"}], "responses": {"200": {"description": "Build canceled successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_CanceledBuildResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Builds"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.builds", "x-fern-sdk-method-name": "cancel", "x-forge-hidden": true}
```
