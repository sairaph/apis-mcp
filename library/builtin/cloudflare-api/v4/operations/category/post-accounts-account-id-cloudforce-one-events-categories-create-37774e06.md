---
title: Creates a new category
page_id: operation-post-accounts-account-id-cloudforce-one-events-categories-create-b70b78e8
path: operations/category
description: Create a new event category for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/categories/create
operation_ids:
    - post_CategoryCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new category

`POST /accounts/{account_id}/cloudforce-one/events/categories/create`

Operation ID: `post_CategoryCreate`

Create a new event category for the account.

## Definition

```yaml
{"operationId": "post_CategoryCreate", "summary": "Creates a new category", "description": "Create a new event category for the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": "T1234", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": "123", "type": "string"}}, "name": {"type": "string", "example": "name"}, "shortname": {"type": "string", "example": "shortname"}}, "required": ["name", "killChain"]}}}}, "responses": {"200": {"description": "Returns the created category.", "content": {"application/json": {"schema": {"type": "object", "properties": {"killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": "T1234", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": "123", "type": "string"}}, "name": {"type": "string", "example": "name"}, "shortname": {"type": "string", "example": "shortname"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "killChain"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Category"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
