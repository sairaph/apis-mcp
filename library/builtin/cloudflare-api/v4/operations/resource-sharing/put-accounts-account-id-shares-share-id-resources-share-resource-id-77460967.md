---
title: Update a share resource
page_id: operation-put-accounts-account-id-shares-share-id-resources-share-resource-id-a31fd82e
path: operations/resource-sharing
description: Update is not immediate, an updated share resource object with a new status will be returned.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/resources/{share_resource_id}
operation_ids:
    - share-resource-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a share resource

`PUT /accounts/{account_id}/shares/{share_id}/resources/{share_resource_id}`

Operation ID: `share-resource-update`

Update is not immediate, an updated share resource object with a new status will be returned.

## Definition

```yaml
{"operationId": "share-resource-update", "summary": "Update a share resource", "description": "Update is not immediate, an updated share resource object with a new status will be returned.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"name": "share_resource_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_resource_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_update_share_resource_request"}}}}, "responses": {"200": {"description": "Share resource updated.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_resource_response_single"}}}}, "4XX": {"description": "Update share resource failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Update share resource failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.resources", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
