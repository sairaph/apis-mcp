---
title: Update column
page_id: operation-patch-accounts-account-id-cloudforce-one-v2-collections-collection-id-co-a44a53cc
path: operations/collections
description: Update name, type, required, or position of a column.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/columns/{column_id}
operation_ids:
    - patch_ColumnUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update column

`PATCH /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/columns/{column_id}`

Operation ID: `patch_ColumnUpdate`

Update name, type, required, or position of a column.

## Definition

```yaml
{"operationId": "patch_ColumnUpdate", "summary": "Update column", "description": "Update name, type, required, or position of a column.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "collection_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "column_id", "in": "path", "description": "Column UUID", "required": true, "schema": {"description": "Column UUID", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"description": "New column name (must be unique)", "type": "string"}, "position": {"description": "Column display order", "type": "number"}, "required": {"description": "Whether column is required", "type": "boolean"}, "type": {"description": "Column type: text, number, boolean, or date", "type": "string", "enum": ["text", "number", "boolean", "date"]}}}}}}, "responses": {"200": {"description": "Column updated successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"type": "string"}, "name": {"type": "string"}, "position": {"type": "number"}, "required": {"type": "boolean"}, "type": {"type": "string"}}, "required": ["id", "name", "type", "required", "position"]}}, "required": ["result"]}}}}, "400": {"description": "Invalid request (validation error)"}, "404": {"description": "Collection or column not found"}, "409": {"description": "New column name conflicts with existing column"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
