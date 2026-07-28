---
title: List share recipients by share ID
page_id: operation-get-accounts-account-id-shares-share-id-recipients-5f72d5d6
path: operations/resource-sharing
description: List share recipients by share ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/recipients
operation_ids:
    - share-recipients-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List share recipients by share ID

`GET /accounts/{account_id}/shares/{share_id}/recipients`

Operation ID: `share-recipients-list`

List share recipients by share ID.

## Definition

```yaml
{"operationId": "share-recipients-list", "summary": "List share recipients by share ID", "description": "List share recipients by share ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"$ref": "#/components/parameters/resource-sharing_include_resources"}, {"$ref": "#/components/parameters/resource-sharing_page"}, {"$ref": "#/components/parameters/resource-sharing_per_page"}], "responses": {"200": {"description": "List account share recipients response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_recipient_response_collection"}}}}, "4XX": {"description": "List account share recipients response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "List account share recipients response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.recipients", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
