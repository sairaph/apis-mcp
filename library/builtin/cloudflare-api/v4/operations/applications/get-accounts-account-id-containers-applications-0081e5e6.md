---
title: List Applications associated with your account
page_id: operation-get-accounts-account-id-containers-applications-9ee1e491
path: operations/applications
description: Lists all the applications that are associated with your account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/applications
operation_ids:
    - listApplications
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Applications associated with your account

`GET /accounts/{account_id}/containers/applications`

Operation ID: `listApplications`

Lists all the applications that are associated with your account

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "listApplications", "summary": "List Applications associated with your account", "description": "Lists all the applications that are associated with your account", "parameters": [{"name": "name", "in": "query", "description": "Filter applications by name", "schema": {"$ref": "#/components/schemas/cc_ApplicationName"}}, {"name": "image", "in": "query", "description": "Filter applications by image", "schema": {"$ref": "#/components/schemas/cc_Image"}}, {"name": "label", "in": "query", "description": "Filter applications by label", "schema": {"type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Get all application associated with your account", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_ListApplications"}}, "required": ["result"], "type": "object"}]}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Applications"]}
```
