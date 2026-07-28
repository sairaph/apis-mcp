---
title: Get account share by ID
page_id: operation-get-accounts-account-id-shares-share-id-07ff1fda
path: operations/resource-sharing
description: Fetches share by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}
operation_ids:
    - shares-get-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account share by ID

`GET /accounts/{account_id}/shares/{share_id}`

Operation ID: `shares-get-by-id`

Fetches share by ID.

## Definition

```yaml
{"operationId": "shares-get-by-id", "summary": "Get account share by ID", "description": "Fetches share by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"$ref": "#/components/parameters/resource-sharing_include_resources"}, {"$ref": "#/components/parameters/resource-sharing_include_recipient_counts"}], "responses": {"200": {"description": "Get account share response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_response_single"}}}}, "4XX": {"description": "Get account share response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Get account share response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
