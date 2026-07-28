---
title: Get bulk operation status
page_id: operation-get-accounts-account-id-rules-lists-bulk-operations-operation-id-d8069396
path: operations/lists
description: |-
    Gets the current status of an asynchronous operation on a list.

    The `status` property can have one of the following values: `pending`, `running`, `completed`, or `failed`. If the status is `failed`, the `error` property will contain a message describing the error.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rules/lists/bulk_operations/{operation_id}
operation_ids:
    - lists-get-bulk-operation-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get bulk operation status

`GET /accounts/{account_id}/rules/lists/bulk_operations/{operation_id}`

Operation ID: `lists-get-bulk-operation-status`

Gets the current status of an asynchronous operation on a list.

The `status` property can have one of the following values: `pending`, `running`, `completed`, or `failed`. If the status is `failed`, the `error` property will contain a message describing the error.

## Definition

```yaml
{"operationId": "lists-get-bulk-operation-status", "summary": "Get bulk operation status", "description": "Gets the current status of an asynchronous operation on a list.\n\nThe `status` property can have one of the following values: `pending`, `running`, `completed`, or `failed`. If the status is `failed`, the `error` property will contain a message describing the error.", "parameters": [{"name": "operation_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_operation_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "responses": {"200": {"description": "Get bulk operation status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_bulk-operation-response-single"}}}}, "4XX": {"description": "Get bulk operation status response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_bulk-operation-response-single"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit", "Account Filter Lists Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
