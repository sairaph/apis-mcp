---
title: Update a permission for dataset
page_id: operation-put-accounts-account-id-cloudforce-one-events-dataset-dataset-id-permiss-2cfdca05
path: operations/permissions
description: Update a permission
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/permissions/{grant_id}
operation_ids:
    - put_PermissionUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a permission for dataset

`PUT /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/permissions/{grant_id}`

Operation ID: `put_PermissionUpdate`

Update a permission

## Definition

```yaml
{"operationId": "put_PermissionUpdate", "summary": "Update a permission for dataset", "description": "Update a permission", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}, {"name": "grant_id", "in": "path", "required": true, "schema": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"role": {"type": "string", "example": "read", "enum": ["read", "write"]}}, "required": ["role"]}}}}, "responses": {"200": {"description": "Permission updated successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"type": "string"}, "success": {"type": "boolean"}}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Permissions"], "x-api-token-group": ["Cloudforce One Write"]}
```
