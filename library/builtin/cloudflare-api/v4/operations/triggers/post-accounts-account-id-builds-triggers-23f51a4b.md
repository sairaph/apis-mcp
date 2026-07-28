---
title: Create trigger
page_id: operation-post-accounts-account-id-builds-triggers-5c89d657
path: operations/triggers
description: Create a new CI/CD trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/builds/triggers
operation_ids:
    - createTrigger
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create trigger

`POST /accounts/{account_id}/builds/triggers`

Operation ID: `createTrigger`

Create a new CI/CD trigger

## Definition

```yaml
{"operationId": "createTrigger", "summary": "Create trigger", "description": "Create a new CI/CD trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_CreateTriggerRequest"}}}}, "responses": {"200": {"description": "Trigger created successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_TriggerResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Triggers"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
