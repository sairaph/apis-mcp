---
title: Create a new share
page_id: operation-post-accounts-account-id-shares-acb5455a
path: operations/resource-sharing
description: Creates a new resource share for sharing Cloudflare resources with other accounts or organizations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/shares
operation_ids:
    - share-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new share

`POST /accounts/{account_id}/shares`

Operation ID: `share-create`

Creates a new resource share for sharing Cloudflare resources with other accounts or organizations.

## Definition

```yaml
{"operationId": "share-create", "summary": "Create a new share", "description": "Creates a new resource share for sharing Cloudflare resources with other accounts or organizations.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_create_share_request"}}}}, "responses": {"201": {"description": "Share created.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_response_single"}}}}, "4XX": {"description": "Create share failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Create share failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
