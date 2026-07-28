---
title: Create Resource Group
page_id: operation-post-accounts-account-id-iam-resource-groups-392c21d8
path: operations/account-resource-groups
description: Create a new Resource Group under the specified account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/iam/resource_groups
operation_ids:
    - account-resource-group-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Resource Group

`POST /accounts/{account_id}/iam/resource_groups`

Operation ID: `account-resource-group-create`

Create a new Resource Group under the specified account.

## Definition

```yaml
{"operationId": "account-resource-group-create", "summary": "Create Resource Group", "description": "Create a new Resource Group under the specified account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_request_create_resource_group"}}}}, "responses": {"200": {"description": "Add Resource Group response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_resource_group"}}, "type": "object"}]}}}}, "4XX": {"description": "Add Resource Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Resource Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.resource-group.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
