---
title: Creates a new tag category (SoT)
page_id: operation-post-accounts-account-id-cloudforce-one-events-tags-categories-create-47b2cbcc
path: operations/tagcategory
description: Creates a new Source-of-Truth tag category for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/tags/categories/create
operation_ids:
    - post_TagCategoryCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new tag category (SoT)

`POST /accounts/{account_id}/cloudforce-one/events/tags/categories/create`

Operation ID: `post_TagCategoryCreate`

Creates a new Source-of-Truth tag category for an account.

## Definition

```yaml
{"operationId": "post_TagCategoryCreate", "summary": "Creates a new tag category (SoT)", "description": "Creates a new Source-of-Truth tag category for an account.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string"}, "name": {"type": "string", "example": "Actor"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Returns the created tag category.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string"}, "description": {"type": "string"}, "name": {"type": "string", "example": "Actor"}, "updatedAt": {"type": "string"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "409": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["TagCategory"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
