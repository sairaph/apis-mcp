---
title: Creates a dataset
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-create-43293eb3
path: operations/dataset
description: Create a new dataset in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/create
operation_ids:
    - post_DatasetCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a dataset

`POST /accounts/{account_id}/cloudforce-one/events/dataset/create`

Operation ID: `post_DatasetCreate`

Create a new dataset in the account.

## Definition

```yaml
{"operationId": "post_DatasetCreate", "summary": "Creates a dataset", "description": "Create a new dataset in the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"isPublic": {"description": "If true, then anyone can search the dataset. If false, then its limited to the account.", "type": "boolean"}, "name": {"description": "Used to describe the dataset within the account context.", "type": "string", "minLength": 1}}, "required": ["name", "isPublic"]}}}}, "responses": {"200": {"description": "Returns dataset information.", "content": {"application/json": {"schema": {"type": "object", "properties": {"deletedAt": {"type": "string"}, "isPublic": {"type": "boolean", "example": true}, "name": {"type": "string", "example": "friendly dataset name"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "isPublic"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Dataset"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
