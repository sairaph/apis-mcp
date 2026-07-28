---
title: Delete list items
page_id: operation-delete-accounts-account-id-rules-lists-list-id-items-04c48678
path: operations/lists
description: |-
    Removes one or more items from a list.

    This operation is asynchronous. To get current the operation status, invoke the `Get bulk operation status` endpoint with the returned `operation_id`.

    There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rules/lists/{list_id}/items
operation_ids:
    - lists-delete-list-items
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete list items

`DELETE /accounts/{account_id}/rules/lists/{list_id}/items`

Operation ID: `lists-delete-list-items`

Removes one or more items from a list.

This operation is asynchronous. To get current the operation status, invoke the `Get bulk operation status` endpoint with the returned `operation_id`.

There is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.

## Definition

```yaml
{"operationId": "lists-delete-list-items", "summary": "Delete list items", "description": "Removes one or more items from a list.\n\nThis operation is asynchronous. To get current the operation status, invoke the `Get bulk operation status` endpoint with the returned `operation_id`.\n\nThere is a limit of 1 pending bulk operation per account. If an outstanding bulk operation is in progress, the request will be rejected.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_list_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"items": {"type": "array", "items": {"example": {"id": "34b12448945f11eaa1b71c4d701ab86e"}, "properties": {"id": {"description": "Defines the unique ID of the item in the List.", "type": "string", "example": "34b12448945f11eaa1b71c4d701ab86e", "maxLength": 32, "minLength": 32, "x-auditable": true}}, "required": ["id"], "type": "object"}, "minItems": 1}}}}}}, "responses": {"200": {"description": "Delete list items response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_lists-async-response"}}}}, "4XX": {"description": "Delete list items response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_lists-async-response"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
