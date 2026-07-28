---
title: Deletes a category
page_id: operation-delete-accounts-account-id-cloudforce-one-events-categories-category-id-71a0f3d4
path: operations/category
description: Delete a category by its identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/categories/{category_id}
operation_ids:
    - delete_CategoryDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes a category

`DELETE /accounts/{account_id}/cloudforce-one/events/categories/{category_id}`

Operation ID: `delete_CategoryDelete`

Delete a category by its identifier.

## Definition

```yaml
{"operationId": "delete_CategoryDelete", "summary": "Deletes a category", "description": "Delete a category by its identifier.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "category_id", "in": "path", "description": "Category UUID.", "required": true, "schema": {"description": "Category UUID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns the uuid of the deleted category.", "content": {"application/json": {"schema": {"type": "object", "properties": {"uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Category"], "x-api-token-group": ["Cloudforce One Write"]}
```
