---
title: Get a single application by id
page_id: operation-get-accounts-account-id-containers-applications-application-id-ed83869e
path: operations/applications
description: Returns a single application by id
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}
operation_ids:
    - getApplication
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a single application by id

`GET /accounts/{account_id}/containers/applications/{application_id}`

Operation ID: `getApplication`

Returns a single application by id

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "getApplication", "summary": "Get a single application by id", "description": "Returns a single application by id", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}], "responses": {"200": {"description": "A single application", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_Application"}}, "required": ["result"], "type": "object"}]}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Response body when an Application is not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-api-token-group": ["Workers Containers Write", "Workers Containers Read"]}
```
