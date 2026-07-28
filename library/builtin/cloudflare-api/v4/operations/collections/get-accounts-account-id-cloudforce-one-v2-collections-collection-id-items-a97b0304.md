---
title: Query collection items
page_id: operation-get-accounts-account-id-cloudforce-one-v2-collections-collection-id-item-02def810
path: operations/collections
description: Retrieves items from a collection with keyset pagination and optional column-based filters
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items
operation_ids:
    - get_ItemQuery
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Query collection items

`GET /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/items`

Operation ID: `get_ItemQuery`

Retrieves items from a collection with keyset pagination and optional column-based filters

## Definition

```yaml
{"operationId": "get_ItemQuery", "summary": "Query collection items", "description": "Retrieves items from a collection with keyset pagination and optional column-based filters", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string"}}, {"name": "cursor", "in": "query", "description": "Opaque pagination cursor from a previous response. Omit for the first page.", "schema": {"description": "Opaque pagination cursor from a previous response. Omit for the first page.", "type": "string"}}, {"name": "limit", "in": "query", "description": "Items per page", "schema": {"description": "Items per page", "type": "number"}}, {"name": "q", "in": "query", "description": "Case-insensitive substring search across all columns. Matches any column containing the term. No relevance ranking.", "schema": {"description": "Case-insensitive substring search across all columns. Matches any column containing the term. No relevance ranking.", "type": "string", "maxLength": 500, "minLength": 1}}], "responses": {"200": {"description": "Items retrieved successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"pagination": {"type": "object", "properties": {"cursors": {"type": "object", "properties": {"current": {"type": "string"}, "first": {"type": "string"}, "next": {"type": "string", "nullable": true}, "prev": {"type": "string", "nullable": true}}, "required": ["first", "current", "prev", "next"]}, "limit": {"type": "number"}, "page": {"type": "number"}, "total": {"type": "number"}}, "required": ["page", "limit", "total", "cursors"]}, "result": {"type": "array", "items": {"properties": {"created_at": {"type": "string"}, "data": {"type": "object", "additionalProperties": true}, "id": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "data", "created_at", "updated_at"], "type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result", "pagination"]}}}}, "400": {"description": "Invalid cursor"}, "404": {"description": "Collection not found"}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
