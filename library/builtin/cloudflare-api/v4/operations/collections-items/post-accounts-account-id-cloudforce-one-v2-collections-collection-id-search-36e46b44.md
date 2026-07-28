---
title: Search items (advanced filtering)
page_id: operation-post-accounts-account-id-cloudforce-one-v2-collections-collection-id-sea-4d8af7a5
path: operations/collections-items
description: Search collection items with advanced filtering. Supports operators (eq, neq, gt, lt, gte, lte, contains, is_empty, is_not_empty) and AND/OR logic for complex queries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/search
operation_ids:
    - post_ItemSearch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search items (advanced filtering)

`POST /accounts/{account_id}/cloudforce-one/v2/collections/{collection_id}/search`

Operation ID: `post_ItemSearch`

Search collection items with advanced filtering. Supports operators (eq, neq, gt, lt, gte, lte, contains, is_empty, is_not_empty) and AND/OR logic for complex queries.

## Definition

```yaml
{"operationId": "post_ItemSearch", "summary": "Search items (advanced filtering)", "description": "Search collection items with advanced filtering. Supports operators (eq, neq, gt, lt, gte, lte, contains, is_empty, is_not_empty) and AND/OR logic for complex queries.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string"}}, {"name": "collection_id", "in": "path", "description": "Collection UUID", "required": true, "schema": {"description": "Collection UUID", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"cursor": {"type": "string"}, "filter": {"description": "Recursive filter supporting AND/OR nesting of conditions. Can be either a leaf condition (field/op/value) or a logical group (and/or array of nested filters).", "type": "object", "oneOf": [{"description": "Leaf filter condition", "properties": {"field": {"description": "Column name to filter on", "type": "string"}, "op": {"description": "Comparison operator", "type": "string", "enum": ["eq", "neq", "gt", "lt", "gte", "lte", "contains", "is_empty", "is_not_empty"]}, "value": {"description": "Value to compare against (not needed for is_empty/is_not_empty)"}}, "required": ["field", "op"], "type": "object"}, {"properties": {"and": {"description": "All conditions must match (AND logic). Each item can recursively contain more filters.", "type": "array", "items": {"description": "Nested filter (can be a leaf condition or another AND/OR group)", "type": "object"}, "minItems": 1}}, "required": ["and"], "type": "object"}, {"properties": {"or": {"description": "At least one condition must match (OR logic). Each item can recursively contain more filters.", "type": "array", "items": {"description": "Nested filter (can be a leaf condition or another AND/OR group)", "type": "object"}, "minItems": 1}}, "required": ["or"], "type": "object"}]}, "limit": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}, "q": {"description": "Case-insensitive substring search across all columns. Matches any column containing the term. No relevance ranking.", "type": "string", "maxLength": 500, "minLength": 1}, "sort": {"type": "object", "properties": {"field": {"description": "Column name to sort by", "type": "string"}, "order": {"type": "string", "default": "asc", "enum": ["asc", "desc"]}}, "required": ["field"]}}}}}}, "responses": {"200": {"description": "Items matching search criteria", "content": {"application/json": {"schema": {"type": "object", "properties": {"pagination": {"type": "object", "properties": {"cursors": {"type": "object", "properties": {"current": {"type": "string"}, "first": {"type": "string"}, "next": {"type": "string", "nullable": true}, "prev": {"type": "string", "nullable": true}}, "required": ["first", "current", "prev", "next"]}, "limit": {"type": "number"}, "page": {"type": "number"}, "total": {"type": "number"}}, "required": ["page", "limit", "total", "cursors"]}, "result": {"type": "array", "items": {"properties": {"created_at": {"type": "string"}, "data": {"type": "object", "additionalProperties": true}, "id": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "data", "created_at", "updated_at"], "type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result", "pagination"]}}}}, "400": {"description": "Invalid filter syntax or cursor"}, "404": {"description": "Collection not found"}}, "security": [{"api_token": []}], "tags": ["Collections - Items"], "x-api-token-group": ["Cloudforce One Write"]}
```
