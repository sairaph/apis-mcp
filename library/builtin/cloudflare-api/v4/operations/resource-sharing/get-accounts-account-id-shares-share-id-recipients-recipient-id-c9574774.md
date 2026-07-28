---
title: Get share recipient by ID
page_id: operation-get-accounts-account-id-shares-share-id-recipients-recipient-id-743473cd
path: operations/resource-sharing
description: Get share recipient by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/recipients/{recipient_id}
operation_ids:
    - share-recipients-get-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get share recipient by ID

`GET /accounts/{account_id}/shares/{share_id}/recipients/{recipient_id}`

Operation ID: `share-recipients-get-by-id`

Get share recipient by ID.

## Definition

```yaml
{"operationId": "share-recipients-get-by-id", "summary": "Get share recipient by ID", "description": "Get share recipient by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"name": "recipient_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_recipient_id"}}, {"$ref": "#/components/parameters/resource-sharing_include_resources"}], "responses": {"200": {"description": "Get account share recipient response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_recipient_response_single"}}}}, "4XX": {"description": "Get account share recipient response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Get account share recipient response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.recipients", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
