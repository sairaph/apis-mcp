---
title: Create a new application
page_id: operation-post-accounts-account-id-containers-applications-8befc99d
path: operations/applications
description: Create a new application. An Application represents an intent to run one or more containers, with the same image, dynamically scheduled based on constraints
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/applications
operation_ids:
    - createApplication
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new application

`POST /accounts/{account_id}/containers/applications`

Operation ID: `createApplication`

Create a new application. An Application represents an intent to run one or more containers, with the same image, dynamically scheduled based on constraints

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "createApplication", "summary": "Create a new application", "description": "Create a new application. An Application represents an intent to run one or more containers, with the same image, dynamically scheduled based on constraints", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_CreateApplicationRequest"}}}}, "responses": {"201": {"description": "A newly created application", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_Application"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Could not create the application because of input/limits reasons, more details in the error code", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "403": {"description": "The account is deactivated", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"]}
```
