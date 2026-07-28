---
title: Get application category
page_id: operation-get-accounts-account-id-resource-library-categories-id-acc76674
path: operations/category
description: Get application category by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/resource-library/categories/{id}
operation_ids:
    - getCategoryById
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get application category

`GET /accounts/{account_id}/resource-library/categories/{id}`

Operation ID: `getCategoryById`

Get application category by ID.

## Definition

```yaml
{"operationId": "getCategoryById", "summary": "Get application category", "description": "Get application category by ID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"type": "string"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "id", "in": "path", "description": "Application category ID.", "required": true, "schema": {"type": "string"}, "example": "0b63249c-95bf-4cc0-a7cc-d7faaaf1dac0"}], "responses": {"200": {"description": "Get application category by id response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/alexandria_get_category_response"}}}}, "4XX": {"description": "Get application category by id response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/alexandria_get_category_response"}, {"$ref": "#/components/schemas/alexandria_api_response_common_failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Category"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.categories.get.by", "x-fern-sdk-method-name": "id"}
```
