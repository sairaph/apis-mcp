---
title: List collections
page_id: operation-get-accounts-account-id-cloudforce-one-v2-collections-c1794f68
path: operations/collections
description: Retrieves all collections for an account with pagination
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/v2/collections
operation_ids:
    - get_CollectionList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List collections

`GET /accounts/{account_id}/cloudforce-one/v2/collections`

Operation ID: `get_CollectionList`

Retrieves all collections for an account with pagination

## Definition

```yaml
{"operationId": "get_CollectionList", "summary": "List collections", "description": "Retrieves all collections for an account with pagination", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID", "required": true, "schema": {"description": "Account ID", "type": "string"}}, {"name": "page", "in": "query", "description": "Page number", "schema": {"description": "Page number", "type": "number"}}, {"name": "limit", "in": "query", "description": "Items per page", "schema": {"description": "Items per page", "type": "number"}}], "responses": {"200": {"description": "Collections retrieved successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"pagination": {"type": "object", "properties": {"limit": {"type": "number"}, "page": {"type": "number"}, "total": {"type": "number"}}, "required": ["page", "limit", "total"]}, "result": {"type": "array", "items": {"properties": {"created_at": {"type": "string"}, "created_by": {"type": "string"}, "id": {"type": "string"}, "item_count": {"type": "number"}, "metadata": {"type": "object", "additionalProperties": true}, "name": {"type": "string"}, "status": {"type": "string"}, "updated_at": {"type": "string"}}, "required": ["id", "name", "status", "item_count", "created_at", "created_by", "updated_at"], "type": "object"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "result", "pagination"]}}}}}, "security": [{"api_token": []}], "tags": ["Collections"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
