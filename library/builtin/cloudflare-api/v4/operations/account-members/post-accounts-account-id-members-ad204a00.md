---
title: Add Member
page_id: operation-post-accounts-account-id-members-81905680
path: operations/account-members
description: Add a user to the list of members for this account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/members
operation_ids:
    - account-members-add-member
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add Member

`POST /accounts/{account_id}/members`

Operation ID: `account-members-add-member`

Add a user to the list of members for this account.

## Definition

```yaml
{"operationId": "account-members-add-member", "summary": "Add Member", "description": "Add a user to the list of members for this account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/iam_create-member-with-roles"}, {"$ref": "#/components/schemas/iam_create-member-with-policies"}]}}}}, "responses": {"200": {"description": "Add Member response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_member_response_with_policies"}}}}, "4XX": {"description": "Add Member response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
