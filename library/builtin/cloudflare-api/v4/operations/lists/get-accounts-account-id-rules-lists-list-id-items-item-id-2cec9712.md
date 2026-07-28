---
title: Get a list item
page_id: operation-get-accounts-account-id-rules-lists-list-id-items-item-id-772cec03
path: operations/lists
description: Fetches a list item in the list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}/items/{item_id}
operation_ids:
    - lists-get-a-list-item
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a list item

`GET /accounts/{account_id}/rules/lists/{list_id}/items/{item_id}`

Operation ID: `lists-get-a-list-item`

Fetches a list item in the list.

## Definition

```yaml
{"operationId": "lists-get-a-list-item", "summary": "Get a list item", "description": "Fetches a list item in the list.", "parameters": [{"name": "item_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_item_id"}}, {"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "responses": {"200": {"description": "Get a list item response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_item-response-single"}}}}, "4XX": {"description": "Get a list item response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_item-response-single"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit", "Account Filter Lists Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
