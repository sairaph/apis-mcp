---
title: Remove Resource Group
page_id: operation-delete-accounts-account-id-iam-resource-groups-resource-group-id-fa159bd0
path: operations/account-resource-groups
description: Remove a resource group from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/iam/resource_groups/{resource_group_id}
operation_ids:
    - account-resource-group-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove Resource Group

`DELETE /accounts/{account_id}/iam/resource_groups/{resource_group_id}`

Operation ID: `account-resource-group-delete`

Remove a resource group from an account.

## Definition

```yaml
{"operationId": "account-resource-group-delete", "summary": "Remove Resource Group", "description": "Remove a resource group from an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "resource_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_resource_group_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Remove Resource Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Remove Member response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Resource Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.resource-group.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
