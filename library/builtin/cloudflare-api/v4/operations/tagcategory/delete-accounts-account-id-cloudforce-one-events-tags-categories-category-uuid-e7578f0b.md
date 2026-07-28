---
title: Deletes a tag category (SoT)
page_id: operation-delete-accounts-account-id-cloudforce-one-events-tags-categories-categor-505df440
path: operations/tagcategory
description: Deletes a Source-of-Truth tag category by UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/tags/categories/{category_uuid}
operation_ids:
    - delete_TagCategoryDelete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Deletes a tag category (SoT)

`DELETE /accounts/{account_id}/cloudforce-one/events/tags/categories/{category_uuid}`

Operation ID: `delete_TagCategoryDelete`

Deletes a Source-of-Truth tag category by UUID.

## Definition

```yaml
{"operationId": "delete_TagCategoryDelete", "summary": "Deletes a tag category (SoT)", "description": "Deletes a Source-of-Truth tag category by UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "category_uuid", "in": "path", "description": "Tag Category UUID.", "required": true, "schema": {"description": "Tag Category UUID.", "type": "string"}}], "responses": {"200": {"description": "Returns the uuid of the deleted tag category.", "content": {"application/json": {"schema": {"type": "object", "properties": {"uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["TagCategory"], "x-api-token-group": ["Cloudforce One Write"]}
```
