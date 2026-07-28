---
title: Update Resource Group
page_id: operation-put-accounts-account-id-iam-resource-groups-resource-group-id-73efd02c
path: operations/account-resource-groups
description: Modify an existing resource group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/iam/resource_groups/{resource_group_id}
operation_ids:
    - account-resource-group-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Resource Group

`PUT /accounts/{account_id}/iam/resource_groups/{resource_group_id}`

Operation ID: `account-resource-group-update`

Modify an existing resource group.

## Definition

```yaml
{"operationId": "account-resource-group-update", "summary": "Update Resource Group", "description": "Modify an existing resource group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "resource_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_resource_group_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_request_update_resource_group"}}}}, "responses": {"200": {"description": "Update Resource Group response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_resource_group"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Resource Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Resource Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.resource-group.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
