---
title: List application categories
page_id: operation-get-accounts-account-id-resource-library-categories-bbf87fae
path: operations/category
description: List application categories.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/resource-library/categories
operation_ids:
    - getCategories
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List application categories

`GET /accounts/{account_id}/resource-library/categories`

Operation ID: `getCategories`

List application categories.

## Definition

```yaml
{"operationId": "getCategories", "summary": "List application categories", "description": "List application categories.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"type": "string"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "limit", "in": "query", "description": "Limit of number of results to return.", "schema": {"type": "integer", "default": 25}}, {"name": "offset", "in": "query", "description": "Offset of results to return.", "schema": {"type": "integer", "default": 0}}], "responses": {"200": {"description": "Get all application categories response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/alexandria_get_categories_response"}}}}, "4XX": {"description": "Get application categories response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/alexandria_get_categories_response"}, {"$ref": "#/components/schemas/alexandria_api_response_common_failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Category"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.categories.get", "x-fern-sdk-method-name": "categories"}
```
