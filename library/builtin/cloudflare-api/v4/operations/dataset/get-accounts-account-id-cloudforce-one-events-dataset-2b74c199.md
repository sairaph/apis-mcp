---
title: Lists all datasets in an account
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-589aa0a1
path: operations/dataset
description: List all datasets accessible to the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset
operation_ids:
    - get_DatasetList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists all datasets in an account

`GET /accounts/{account_id}/cloudforce-one/events/dataset`

Operation ID: `get_DatasetList`

List all datasets accessible to the account.

## Definition

```yaml
{"operationId": "get_DatasetList", "summary": "Lists all datasets in an account", "description": "List all datasets accessible to the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "includeDeleted", "in": "query", "description": "When true, include soft-deleted datasets in the response. Each item includes a `deletedAt` field (ISO 8601 or null). Default: false.", "schema": {"description": "When true, include soft-deleted datasets in the response. Each item includes a `deletedAt` field (ISO 8601 or null). Default: false.", "type": "boolean"}}], "responses": {"200": {"description": "Returns a list of dataset in an account.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"deletedAt": {"type": "string"}, "isPublic": {"type": "boolean", "example": true}, "name": {"type": "string", "example": "friendly dataset name"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "isPublic"], "type": "object"}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Dataset"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
