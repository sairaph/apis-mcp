---
title: List Accounts
page_id: operation-get-accounts-4dc31298
path: operations/accounts
description: List all accounts you have ownership or verified access to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts
operation_ids:
    - accounts-list-accounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Accounts

`GET /accounts`

Operation ID: `accounts-list-accounts`

List all accounts you have ownership or verified access to.

## Definition

```yaml
{"operationId": "accounts-list-accounts", "summary": "List Accounts", "description": "List all accounts you have ownership or verified access to.", "parameters": [{"name": "name", "in": "query", "schema": {"description": "Name of the account.", "type": "string", "example": "example.com", "readOnly": true}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order results.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List Accounts response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_response_collection_accounts"}}}}, "4XX": {"description": "List Accounts response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
