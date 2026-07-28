---
title: Resource Group Details
page_id: operation-get-accounts-account-id-iam-resource-groups-resource-group-id-d693cdf1
path: operations/account-resource-groups
description: Get information about a specific resource group in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/resource_groups/{resource_group_id}
operation_ids:
    - account-resource-group-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Resource Group Details

`GET /accounts/{account_id}/iam/resource_groups/{resource_group_id}`

Operation ID: `account-resource-group-details`

Get information about a specific resource group in an account.

## Definition

```yaml
{"operationId": "account-resource-group-details", "summary": "Resource Group Details", "description": "Get information about a specific resource group in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "resource_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_resource_group_identifier"}}], "responses": {"200": {"description": "Resource Group Details response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_resource_group"}}, "type": "object"}]}}}}, "4XX": {"description": "Resource Group Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Resource Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.resource-group.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
