---
title: Member Details
page_id: operation-get-accounts-account-id-members-member-id-773eb9c9
path: operations/account-members
description: Get information about a specific member of an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/members/{member_id}
operation_ids:
    - account-members-member-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Member Details

`GET /accounts/{account_id}/members/{member_id}`

Operation ID: `account-members-member-details`

Get information about a specific member of an account.

## Definition

```yaml
{"operationId": "account-members-member-details", "summary": "Member Details", "description": "Get information about a specific member of an account.", "parameters": [{"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_membership_components-schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "Member Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_member_response_with_policies"}}}}, "4XX": {"description": "Member Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
