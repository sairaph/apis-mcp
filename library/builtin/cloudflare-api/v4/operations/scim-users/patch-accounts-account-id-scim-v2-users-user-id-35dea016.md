---
title: Patch SCIM User
page_id: operation-patch-accounts-account-id-scim-v2-users-user-id-31418285
path: operations/scim-users
description: 'Partially updates a SCIM User via PATCH operations (RFC 7644 Section 3.5.2). Supports updating `userName`, `name.givenName`, `name.familyName`, and `active`. Setting `active: false` deprovisions the user (removes them from the account). For IdP compatibility, `emails[type eq "work"].value` is also accepted as an alias for `userName`.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/scim/v2/Users/{user_id}
operation_ids:
    - scim-users-patch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch SCIM User

`PATCH /accounts/{account_id}/scim/v2/Users/{user_id}`

Operation ID: `scim-users-patch`

Partially updates a SCIM User via PATCH operations (RFC 7644 Section 3.5.2). Supports updating `userName`, `name.givenName`, `name.familyName`, and `active`. Setting `active: false` deprovisions the user (removes them from the account). For IdP compatibility, `emails[type eq "work"].value` is also accepted as an alias for `userName`.

## Definition

```yaml
{"operationId": "scim-users-patch", "summary": "Patch SCIM User", "description": "Partially updates a SCIM User via PATCH operations (RFC 7644 Section 3.5.2). Supports updating `userName`, `name.givenName`, `name.familyName`, and `active`. Setting `active: false` deprovisions the user (removes them from the account). For IdP compatibility, `emails[type eq \"work\"].value` is also accepted as an alias for `userName`.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_scim_user_identifier"}}], "requestBody": {"required": true, "content": {"application/scim+json": {"examples": {"path-based": {"summary": "Path-based style", "value": {"Operations": [{"op": "replace", "path": "active", "value": false}], "schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}, "value-object": {"summary": "Value-object style (no path)", "value": {"Operations": [{"op": "replace", "value": {"active": false}}], "schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}}, "schema": {"$ref": "#/components/schemas/iam_scim_patch_op_request"}}}}, "responses": {"200": {"description": "Patch SCIM User response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user"}}}}, "4XX": {"description": "Patch SCIM User response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Users"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.update"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
