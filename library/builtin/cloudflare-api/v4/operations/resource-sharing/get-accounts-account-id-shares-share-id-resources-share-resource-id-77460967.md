---
title: Get share resource by ID
page_id: operation-get-accounts-account-id-shares-share-id-resources-share-resource-id-6b1950d9
path: operations/resource-sharing
description: Get share resource by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/resources/{share_resource_id}
operation_ids:
    - share-resources-get-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get share resource by ID

`GET /accounts/{account_id}/shares/{share_id}/resources/{share_resource_id}`

Operation ID: `share-resources-get-by-id`

Get share resource by ID.

## Definition

```yaml
{"operationId": "share-resources-get-by-id", "summary": "Get share resource by ID", "description": "Get share resource by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"name": "share_resource_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_resource_id"}}], "responses": {"200": {"description": "Get account share resource response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_resource_response_single"}}}}, "4XX": {"description": "Get account share resource response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Get account share resource response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.resources", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
