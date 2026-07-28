---
title: List all application versions
page_id: operation-get-accounts-account-id-containers-applications-application-id-versions-8dea6b1b
path: operations/applications
description: Returns all versions for this application
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/applications/{application_id}/versions
operation_ids:
    - listApplicationVersions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List all application versions

`GET /accounts/{account_id}/containers/applications/{application_id}/versions`

Operation ID: `listApplicationVersions`

Returns all versions for this application

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "listApplicationVersions", "summary": "List all application versions", "description": "Returns all versions for this application", "parameters": [{"name": "application_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cc_ApplicationID"}}], "responses": {"200": {"description": "List application versions", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cc_ApplicationVersion"}}}, "required": ["result"], "type": "object"}]}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "Response body when an Application is not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"], "x-api-token-group": ["Workers Containers Write", "Workers Containers Read"]}
```
