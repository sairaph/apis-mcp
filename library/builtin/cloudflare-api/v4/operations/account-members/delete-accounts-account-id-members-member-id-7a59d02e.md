---
title: Remove Member
page_id: operation-delete-accounts-account-id-members-member-id-b5a55ec1
path: operations/account-members
description: Remove a member from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/members/{member_id}
operation_ids:
    - account-members-remove-member
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove Member

`DELETE /accounts/{account_id}/members/{member_id}`

Operation ID: `account-members-remove-member`

Remove a member from an account.

## Definition

```yaml
{"operationId": "account-members-remove-member", "summary": "Remove Member", "description": "Remove a member from an account.", "parameters": [{"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Remove Member response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Remove Member response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
