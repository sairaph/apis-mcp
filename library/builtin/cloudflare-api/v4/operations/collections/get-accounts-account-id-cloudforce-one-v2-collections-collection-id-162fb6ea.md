---
title: Get collection
page_id: operation-get-accounts-account-id-cloudforce-one-v2-collections-collection-id-f78d1284
path: operations/collections
description: Retrieves a single collection by ID with its schema and metadata
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}
operation_ids:
    - get_CollectionGet
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get collection

`GET /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}`

Operation ID: `get_CollectionGet`

Retrieves a single collection by ID with its schema and metadata

## Definition

```yaml
{"operationId": "get_CollectionGet", "summary": "Get collection", "description": "Retrieves a single collection by ID with its schema and metadata", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string"}}], "responses": {"200": {"description": "Collection retrieved successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"columns": {"type": "array", "items": {"properties": {"default": {"type": "object", "nullable": true}, "id": {"type": "string"}, "name": {"type": "string"}, "position": {"type": "number"}, "required": {"type": "boolean"}, "type": {"type": "string"}}, "required": ["id", "name", "type", "required", "position"], "type": "object"}}, "created_at": {"type": "string"}, "created_by": {"type": "string"}, "id": {"type": "string"}, "item_count": {"type": "number"}, "metadata": {"type": "object", "additionalProperties": true}, "name": {"type": "string"}, "status": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "name", "status", "item_count", "created_at", "created_by", "updated_at", "columns"]}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result"]}}}}, "404": {"description": "Collection not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
