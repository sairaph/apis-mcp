---
title: Update all list items
page_id: operation-put-accounts-account-id-rules-lists-list-id-items-9bd22b6c
path: operations/lists
description: |-
    Removes all existing items from the list and adds the provided items to the list.

    This operation is asynchronous. To get current the operation status, invoke the `Get bulk operation status` endpoint with the returned `operation_id`.

    There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}/items
operation_ids:
    - lists-update-all-list-items
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update all list items

`PUT /accounts/{account_id}/rules/lists/{list_id}/items`

Operation ID: `lists-update-all-list-items`

Removes all existing items from the list and adds the provided items to the list.

This operation is asynchronous. To get current the operation status, invoke the `Get bulk operation status` endpoint with the returned `operation_id`.

There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.

## Definition

```yaml
{"operationId": "lists-update-all-list-items", "summary": "Update all list items", "description": "Removes all existing items from the list and adds the provided items to the list.\n\nThis operation is asynchronous. To get current the operation status, invoke the `Get bulk operation status` endpoint with the returned `operation_id`.\n\nThere is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_items-update-request-collection"}}}}, "responses": {"200": {"description": "Update all list items response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_lists-async-response"}}}}, "4XX": {"description": "Update all list items response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_lists-async-response"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
