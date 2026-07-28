---
title: Update a share's recipients
page_id: operation-put-accounts-account-id-shares-share-id-recipients-e7cf5163
path: operations/resource-sharing
description: Changes a share's recipients to match the given list. Returns an error if the share targets an organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/recipients
operation_ids:
    - share-recipients-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a share's recipients

`PUT /accounts/{account_id}/shares/{share_id}/recipients`

Operation ID: `share-recipients-update`

Changes a share's recipients to match the given list. Returns an error if the share targets an organization.

## Definition

```yaml
{"operationId": "share-recipients-update", "summary": "Update a share's recipients", "description": "Changes a share's recipients to match the given list. Returns an error if the share targets an organization.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_update_share_recipients_request"}}}}, "responses": {"204": {"description": "Empty body"}, "4XX": {"description": "Update share recipients failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Update share recipients failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.recipients", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
