---
title: Get a list
page_id: operation-get-accounts-account-id-rules-lists-list-id-6ac67dd9
path: operations/lists
description: Fetches the details of a list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}
operation_ids:
    - lists-get-a-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a list

`GET /accounts/{account_id}/rules/lists/{list_id}`

Operation ID: `lists-get-a-list`

Fetches the details of a list.

## Definition

```yaml
{"operationId": "lists-get-a-list", "summary": "Get a list", "description": "Fetches the details of a list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "responses": {"200": {"description": "Get a list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_list-response-collection"}}}}, "4XX": {"description": "Get a list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_list-response-collection"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
