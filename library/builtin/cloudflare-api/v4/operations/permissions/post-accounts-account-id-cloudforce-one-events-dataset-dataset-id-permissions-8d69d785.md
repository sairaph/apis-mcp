---
title: Create a permission for dataset
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-dataset-id-permis-5810850b
path: operations/permissions
description: Create a permission
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/permissions
operation_ids:
    - post_PermissionCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a permission for dataset

`POST /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/permissions`

Operation ID: `post_PermissionCreate`

Create a permission

## Definition

```yaml
{"operationId": "post_PermissionCreate", "summary": "Create a permission for dataset", "description": "Create a permission", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"role": {"type": "string", "example": "read", "enum": ["read", "write"]}, "subjectId": {"type": "string", "example": "123"}, "subjectType": {"type": "string", "example": "account", "enum": ["account", "group"]}}, "required": ["subjectType", "subjectId", "role"]}}}}, "responses": {"200": {"description": "Returns the created permission.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "resourceId": {"description": "The resource ID this permission applies to account_id or group_id", "type": "string", "example": "08846b9f-dba9-4410-be6f-cf883b5ea8d2"}, "resourceType": {"type": "string", "example": "dataset", "enum": ["dataset"]}, "role": {"type": "string", "example": "read", "enum": ["read", "write"]}, "subjectId": {"type": "string", "example": "123"}, "subjectType": {"type": "string", "example": "account", "enum": ["account", "group"]}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "subjectType", "subjectId", "role", "resourceType", "createdAt", "updatedAt"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Permissions"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
