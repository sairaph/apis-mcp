---
title: Add column to collection
page_id: operation-post-accounts-account-id-cloudforce-one-v2-collections-collection-id-col-ac96cdfe
path: operations/collections
description: Adds a new column to an existing collection schema. Existing items will have NULL for the new column unless a default value is provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/columns
operation_ids:
    - post_ColumnAdd
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add column to collection

`POST /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/columns`

Operation ID: `post_ColumnAdd`

Adds a new column to an existing collection schema. Existing items will have NULL for the new column unless a default value is provided.

## Definition

```yaml
{"operationId": "post_ColumnAdd", "summary": "Add column to collection", "description": "Adds a new column to an existing collection schema. Existing items will have NULL for the new column unless a default value is provided.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string", "example": "10d79d097895ae7ed7942a2b3832186c"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string", "example": "550e8400-e29b-41d4-a716-446655440000"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"default": {"type": "object", "nullable": true}, "name": {"type": "string", "maxLength": 64, "minLength": 1, "pattern": "^[a-zA-Z0-9_]+$"}, "required": {"type": "boolean", "default": false}, "type": {"type": "string", "enum": ["text", "number", "boolean", "date"]}}, "required": ["name", "type"]}}}}, "responses": {"200": {"description": "Column added successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"type": "string", "format": "uuid"}, "name": {"type": "string"}, "position": {"type": "number"}, "required": {"type": "boolean"}, "type": {"type": "string", "enum": ["text", "number", "boolean", "date"]}}, "required": ["id", "name", "type", "required", "position"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Collection not found"}, "409": {"description": "Column with this name already exists"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
