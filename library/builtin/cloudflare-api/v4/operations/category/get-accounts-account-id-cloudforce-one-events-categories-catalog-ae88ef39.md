---
title: Lists categories
page_id: operation-get-accounts-account-id-cloudforce-one-events-categories-catalog-11049eb2
path: operations/category
description: List all categories stored in the account catalog.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/categories/catalog
operation_ids:
    - get_CategoryListComplete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists categories

`GET /accounts/{account_id}/cloudforce-one/events/categories/catalog`

Operation ID: `get_CategoryListComplete`

List all categories stored in the account catalog.

## Definition

```yaml
{"operationId": "get_CategoryListComplete", "summary": "Lists categories", "description": "List all categories stored in the account catalog.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "responses": {"200": {"description": "Returns a list of categories.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"killChain": {"type": "number"}, "mitreAttack": {"type": "array", "items": {"example": "T1234", "type": "string"}}, "mitreCapec": {"type": "array", "items": {"example": "123", "type": "string"}}, "name": {"type": "string", "example": "name"}, "shortname": {"type": "string", "example": "shortname"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["uuid", "name", "killChain"], "type": "object"}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Category"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
