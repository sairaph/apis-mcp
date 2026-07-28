---
title: Delete a share recipient
page_id: operation-delete-accounts-account-id-shares-share-id-recipients-recipient-id-a71f98fd
path: operations/resource-sharing
description: Deletion is not immediate, an updated share recipient object with a new status will be returned.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/shares/{share_id}/recipients/{recipient_id}
operation_ids:
    - share-recipient-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a share recipient

`DELETE /accounts/{account_id}/shares/{share_id}/recipients/{recipient_id}`

Operation ID: `share-recipient-delete`

Deletion is not immediate, an updated share recipient object with a new status will be returned.

## Definition

```yaml
{"operationId": "share-recipient-delete", "summary": "Delete a share recipient", "description": "Deletion is not immediate, an updated share recipient object with a new status will be returned.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"name": "share_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_share_id"}}, {"name": "recipient_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_recipient_id"}}], "responses": {"200": {"description": "Share recipient deleted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_recipient_response_single"}}}}, "4XX": {"description": "Delete share recipient failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "Delete share recipient failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing.recipients", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
