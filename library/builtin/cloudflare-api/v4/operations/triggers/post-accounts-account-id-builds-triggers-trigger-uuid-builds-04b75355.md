---
title: Create manual build
page_id: operation-post-accounts-account-id-builds-triggers-trigger-uuid-builds-af0b83c6
path: operations/triggers
description: Trigger a manual build for a specific trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/builds/triggers/{trigger_uuid}/builds
operation_ids:
    - createManualBuild
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create manual build

`POST /accounts/{account_id}/builds/triggers/{trigger_uuid}/builds`

Operation ID: `createManualBuild`

Trigger a manual build for a specific trigger

## Definition

```yaml
{"operationId": "createManualBuild", "summary": "Create manual build", "description": "Trigger a manual build for a specific trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_TriggerUuid"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_CreateBuildRequest"}}}}, "responses": {"200": {"description": "Build created successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_InsertBuildResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Triggers"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers", "x-fern-sdk-method-name": "create-build", "x-forge-hidden": true}
```
