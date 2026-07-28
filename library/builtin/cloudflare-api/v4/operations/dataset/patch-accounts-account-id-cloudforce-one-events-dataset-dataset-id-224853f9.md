---
title: Updates an existing dataset
page_id: operation-patch-accounts-account-id-cloudforce-one-events-dataset-dataset-id-077dbbaa
path: operations/dataset
description: Update an existing dataset by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}
operation_ids:
    - patch_DatasetUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates an existing dataset

`PATCH /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}`

Operation ID: `patch_DatasetUpdate`

Update an existing dataset by its identifier.

## Definition

```yaml
{"operationId": "patch_DatasetUpdate", "summary": "Updates an existing dataset", "description": "Update an existing dataset by its identifier.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"isPublic": {"description": "If true, then anyone can search the dataset. If false, then its limited to the account.", "type": "boolean"}, "name": {"description": "Used to describe the dataset within the account context.", "type": "string", "minLength": 1}}, "required": ["name", "isPublic"]}}}}, "responses": {"200": {"description": "Returns dataset information.", "content": {"application/json": {"schema": {"type": "object", "properties": {"deletedAt": {"type": "string"}, "isPublic": {"type": "boolean", "example": true}, "name": {"type": "string", "example": "friendly dataset name"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "isPublic"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Dataset"], "x-api-token-group": ["Cloudforce One Write"]}
```
