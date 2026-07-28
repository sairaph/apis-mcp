---
title: Update trigger
page_id: operation-patch-accounts-account-id-builds-triggers-trigger-uuid-b331d55b
path: operations/triggers
description: Update an existing CI/CD trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/builds/triggers/{trigger_uuid}
operation_ids:
    - updateTrigger
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update trigger

`PATCH /accounts/{account_id}/builds/triggers/{trigger_uuid}`

Operation ID: `updateTrigger`

Update an existing CI/CD trigger

## Definition

```yaml
{"operationId": "updateTrigger", "summary": "Update trigger", "description": "Update an existing CI/CD trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_TriggerUuid"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_UpdateTriggerRequest"}}}}, "responses": {"200": {"description": "Trigger updated successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_TriggerResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Triggers"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
