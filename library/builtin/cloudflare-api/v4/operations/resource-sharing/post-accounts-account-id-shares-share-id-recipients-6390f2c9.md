---
title: Create a new share recipient
page_id: operation-post-accounts-account-id-shares-share-id-recipients-f06dc41b
path: operations/resource-sharing
description: Adds a recipient to a resource share, granting them access to the shared resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/recipients
operation_ids:
    - share-recipient-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new share recipient

`POST /accounts/{account_id}/shares/{share_id}/recipients`

Operation ID: `share-recipient-create`

Adds a recipient to a resource share, granting them access to the shared resources.

## Definition

```yaml
{"operationId": "share-recipient-create", "summary": "Create a new share recipient", "description": "Adds a recipient to a resource share, granting them access to the shared resources.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_create_share_recipient_request"}}}}, "responses": {"201": {"description": "Share recipient created.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_recipient_response_single"}}}}, "4XX": {"description": "Create share recipient failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Create share recipient failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.recipients", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
