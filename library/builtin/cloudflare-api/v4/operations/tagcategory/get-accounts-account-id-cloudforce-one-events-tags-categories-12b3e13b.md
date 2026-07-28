---
title: Lists all tag categories (SoT)
page_id: operation-get-accounts-account-id-cloudforce-one-events-tags-categories-4fe24b04
path: operations/tagcategory
description: Returns all Source-of-Truth tag categories for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/tags/categories
operation_ids:
    - get_TagCategoryList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists all tag categories (SoT)

`GET /accounts/{account_id}/cloudforce-one/events/tags/categories`

Operation ID: `get_TagCategoryList`

Returns all Source-of-Truth tag categories for an account.

## Definition

```yaml
{"operationId": "get_TagCategoryList", "summary": "Lists all tag categories (SoT)", "description": "Returns all Source-of-Truth tag categories for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "search", "in": "query", "schema": {"type": "string"}}], "responses": {"200": {"description": "Returns a list of tag categories.", "content": {"application/json": {"schema": {"type": "object", "properties": {"categories": {"type": "array", "items": {"properties": {"createdAt": {"type": "string"}, "description": {"type": "string"}, "name": {"type": "string", "example": "Actor"}, "updatedAt": {"type": "string"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name"], "type": "object"}}}, "required": ["categories"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["TagCategory"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
