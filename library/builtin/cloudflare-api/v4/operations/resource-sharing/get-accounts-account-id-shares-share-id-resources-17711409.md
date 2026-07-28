---
title: List share resources by share ID
page_id: operation-get-accounts-account-id-shares-share-id-resources-b2ff7be6
path: operations/resource-sharing
description: List share resources by share ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/resources
operation_ids:
    - share-resources-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List share resources by share ID

`GET /accounts/{account_id}/shares/{share_id}/resources`

Operation ID: `share-resources-list`

List share resources by share ID.

## Definition

```yaml
{"operationId": "share-resources-list", "summary": "List share resources by share ID", "description": "List share resources by share ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"$ref": "#/components/parameters/resource-sharing_resource_status"}, {"$ref": "#/components/parameters/resource-sharing_resource_type"}, {"$ref": "#/components/parameters/resource-sharing_page"}, {"$ref": "#/components/parameters/resource-sharing_per_page"}], "responses": {"200": {"description": "List account share resources response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_resource_response_collection"}}}}, "4XX": {"description": "List account share resources response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "List account share resources response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.resources", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
