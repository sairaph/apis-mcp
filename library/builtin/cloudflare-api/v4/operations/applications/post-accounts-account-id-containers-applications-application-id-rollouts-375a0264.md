---
title: Create a new rollout for an application
page_id: operation-post-accounts-account-id-containers-applications-application-id-rollouts-4a8ce533
path: operations/applications
description: A rollout can be used to update the application's configuration across instances with minimal downtime.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/rollouts
operation_ids:
    - createApplicationRollout
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new rollout for an application

`POST /accounts/{account_id}/containers/applications/{application_id}/rollouts`

Operation ID: `createApplicationRollout`

A rollout can be used to update the application's configuration across instances with minimal downtime.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "createApplicationRollout", "summary": "Create a new rollout for an application", "description": "A rollout can be used to update the application's configuration across instances with minimal downtime.", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_CreateApplicationRolloutRequest"}}}}, "responses": {"201": {"description": "Application rollout created successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ApplicationRollout"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Can't update the application rollout because it has bad inputs", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "403": {"description": "The account is deactivated", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Response body when an Application is not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications", "Rollouts"]}
```
