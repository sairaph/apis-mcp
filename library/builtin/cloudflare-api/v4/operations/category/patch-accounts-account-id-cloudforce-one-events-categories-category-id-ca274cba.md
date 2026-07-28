---
title: Updates a category
page_id: operation-patch-accounts-account-id-cloudforce-one-events-categories-category-id-121112e7
path: operations/category
description: Update an existing category by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/categories/{category_id}
operation_ids:
    - patch_CategoryUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates a category

`PATCH /accounts/{account_id}/cloudforce-one/events/categories/{category_id}`

Operation ID: `patch_CategoryUpdate`

Update an existing category by its identifier.

## Definition

```yaml
{"operationId": "patch_CategoryUpdate", "summary": "Updates a category", "description": "Update an existing category by its identifier.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "category_id", "in": "path", "description": "Category UUID.", "required": true, "schema": {"description": "Category UUID.", "type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": "T1234", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": "123", "type": "string"}}, "name": {"type": "string", "example": "name"}, "shortname": {"type": "string", "example": "shortname"}}}}}}, "responses": {"200": {"description": "Returns the updated category.", "content": {"application/json": {"schema": {"type": "object", "properties": {"killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": "T1234", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": "123", "type": "string"}}, "name": {"type": "string", "example": "name"}, "shortname": {"type": "string", "example": "shortname"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "killChain"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Category"], "x-api-token-group": ["Cloudforce One Write"]}
```
