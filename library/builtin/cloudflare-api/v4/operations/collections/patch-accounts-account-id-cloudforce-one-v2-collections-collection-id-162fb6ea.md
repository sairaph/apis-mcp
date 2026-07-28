---
title: Update collection
page_id: operation-patch-accounts-account-id-cloudforce-one-v2-collections-collection-id-30b868f2
path: operations/collections
description: Updates collection name and/or metadata. Schema (columns) cannot be modified.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}
operation_ids:
    - patch_CollectionUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update collection

`PATCH /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}`

Operation ID: `patch_CollectionUpdate`

Updates collection name and/or metadata. Schema (columns) cannot be modified.

## Definition

```yaml
{"operationId": "patch_CollectionUpdate", "summary": "Update collection", "description": "Updates collection name and/or metadata. Schema (columns) cannot be modified.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"metadata": {"type": "object", "properties": {"description": {"description": "Collection description", "type": "string"}, "project_id": {"description": "Project ID", "type": "string"}, "tags": {"type": "array", "items": {"type": "string"}}}}, "name": {"description": "Collection name", "type": "string"}}}}}}, "responses": {"200": {"description": "Collection updated successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"created_at": {"type": "string"}, "created_by": {"type": "string"}, "id": {"type": "string"}, "item_count": {"type": "number"}, "metadata": {"type": "object", "additionalProperties": true}, "name": {"type": "string"}, "status": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "name", "status", "item_count", "created_at", "created_by", "updated_at"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "404": {"description": "Collection not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write"]}
```
