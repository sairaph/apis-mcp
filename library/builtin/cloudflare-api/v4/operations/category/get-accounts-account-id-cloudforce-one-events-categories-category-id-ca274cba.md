---
title: Reads a category
page_id: operation-get-accounts-account-id-cloudforce-one-events-categories-category-id-a532bedd
path: operations/category
description: Retrieve a single category by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/categories/{category_id}
operation_ids:
    - get_CategoryRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reads a category

`GET /accounts/{account_id}/cloudforce-one/events/categories/{category_id}`

Operation ID: `get_CategoryRead`

Retrieve a single category by its identifier.

## Definition

```yaml
{"operationId": "get_CategoryRead", "summary": "Reads a category", "description": "Retrieve a single category by its identifier.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "category_id", "in": "path", "description": "Category UUID.", "required": true, "schema": {"description": "Category UUID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns a category.", "content": {"application/json": {"schema": {"type": "object", "properties": {"killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": "T1234", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": "123", "type": "string"}}, "name": {"type": "string", "example": "name"}, "shortname": {"type": "string", "example": "shortname"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "killChain"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Category"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
