---
title: List Resource Groups
page_id: operation-get-accounts-account-id-iam-resource-groups-a5cfe6af
path: operations/account-resource-groups
description: List all the resource groups for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/resource_groups
operation_ids:
    - account-resource-group-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Resource Groups

`GET /accounts/{account_id}/iam/resource_groups`

Operation ID: `account-resource-group-list`

List all the resource groups for an account.

## Definition

```yaml
{"operationId": "account-resource-group-list", "summary": "List Resource Groups", "description": "List all the resource groups for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "id", "in": "query", "schema": {"description": "ID of the resource group to be fetched.", "allOf": [{"$ref": "#/components/schemas/iam_resource_group_identifier"}]}}, {"name": "name", "in": "query", "schema": {"description": "Name of the resource group to be fetched.", "type": "string", "example": "NameOfTheResourceGroup"}}], "responses": {"200": {"description": "List Resource Groups response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/iam_resource_group"}}}, "type": "object"}], "title": "List of resource groups in the account."}}}}, "4XX": {"description": "List Resource Groups response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Resource Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.resource-group.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
