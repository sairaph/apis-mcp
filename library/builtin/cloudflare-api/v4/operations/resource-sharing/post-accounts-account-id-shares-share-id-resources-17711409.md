---
title: Create a new share resource
page_id: operation-post-accounts-account-id-shares-share-id-resources-da5f4f05
path: operations/resource-sharing
description: Adds a resource to an existing share, making it available to share recipients.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/resources
operation_ids:
    - share-resource-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new share resource

`POST /accounts/{account_id}/shares/{share_id}/resources`

Operation ID: `share-resource-create`

Adds a resource to an existing share, making it available to share recipients.

## Definition

```yaml
{"operationId": "share-resource-create", "summary": "Create a new share resource", "description": "Adds a resource to an existing share, making it available to share recipients.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_create_share_resource_request"}}}}, "responses": {"201": {"description": "Share resource created.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_resource_response_single"}}}}, "4XX": {"description": "Create share resource failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Create share resource failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.resources", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
