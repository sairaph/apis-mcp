---
title: Delete a list
page_id: operation-delete-accounts-account-id-rules-lists-list-id-ccff2127
path: operations/lists
description: Deletes a specific list and all its items.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}
operation_ids:
    - lists-delete-a-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a list

`DELETE /accounts/{account_id}/rules/lists/{list_id}`

Operation ID: `lists-delete-a-list`

Deletes a specific list and all its items.

## Definition

```yaml
{"operationId": "lists-delete-a-list", "summary": "Delete a list", "description": "Deletes a specific list and all its items.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_list-delete-response-collection"}}}}, "4XX": {"description": "Delete a list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_list-delete-response-collection"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
