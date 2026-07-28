---
title: Update Member
page_id: operation-put-accounts-account-id-members-member-id-e95ef4fc
path: operations/account-members
description: Modify an account member.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/members/{member_id}
operation_ids:
    - account-members-update-member
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Member

`PUT /accounts/{account_id}/members/{member_id}`

Operation ID: `account-members-update-member`

Modify an account member.

## Definition

```yaml
{"operationId": "account-members-update-member", "summary": "Update Member", "description": "Modify an account member.", "parameters": [{"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/iam_update-member-with-roles"}, {"$ref": "#/components/schemas/iam_update-member-with-policies"}]}}}}, "responses": {"200": {"description": "Update Member response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_member_response_with_policies"}}}}, "4XX": {"description": "Update Member response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
