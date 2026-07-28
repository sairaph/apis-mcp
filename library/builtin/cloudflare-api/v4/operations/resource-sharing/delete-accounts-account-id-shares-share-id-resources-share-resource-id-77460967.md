---
title: Delete a share resource
page_id: operation-delete-accounts-account-id-shares-share-id-resources-share-resource-id-82171ef8
path: operations/resource-sharing
description: Deletion is not immediate, an updated share resource object with a new status will be returned.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/resources/{share_resource_id}
operation_ids:
    - share-resource-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a share resource

`DELETE /accounts/{account_id}/shares/{share_id}/resources/{share_resource_id}`

Operation ID: `share-resource-delete`

Deletion is not immediate, an updated share resource object with a new status will be returned.

## Definition

```yaml
{"operationId": "share-resource-delete", "summary": "Delete a share resource", "description": "Deletion is not immediate, an updated share resource object with a new status will be returned.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"name": "share_resource_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_resource_id"}}], "responses": {"200": {"description": "Share resource deleted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_resource_response_single"}}}}, "4XX": {"description": "Delete share resource failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Delete share resource failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.resources", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
