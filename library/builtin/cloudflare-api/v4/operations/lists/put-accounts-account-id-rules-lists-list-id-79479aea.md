---
title: Update a list
page_id: operation-put-accounts-account-id-rules-lists-list-id-1be4edf6
path: operations/lists
description: Updates the description of a list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}
operation_ids:
    - lists-update-a-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a list

`PUT /accounts/{account_id}/rules/lists/{list_id}`

Operation ID: `lists-update-a-list`

Updates the description of a list.

## Definition

```yaml
{"operationId": "lists-update-a-list", "summary": "Update a list", "description": "Updates the description of a list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/lists_description"}}}}}}, "responses": {"200": {"description": "Update a list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_list-response-collection"}}}}, "4XX": {"description": "Update a list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_list-response-collection"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
