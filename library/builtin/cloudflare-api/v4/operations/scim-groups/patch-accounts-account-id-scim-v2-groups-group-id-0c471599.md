---
title: Patch SCIM Group
page_id: operation-patch-accounts-account-id-scim-v2-groups-group-id-597a1688
path: operations/scim-groups
description: Partially updates a SCIM Group via PATCH operations (RFC 7644 Section 3.5.2). Supports add, remove, and replace operations on `members`, `displayName`, and `externalId`. For system groups (prefixed `cloudflare-v1-`), only member management operations are supported.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/scim/v2/Groups/{group_id}
operation_ids:
    - scim-groups-patch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch SCIM Group

`PATCH /accounts/{account_id}/scim/v2/Groups/{group_id}`

Operation ID: `scim-groups-patch`

Partially updates a SCIM Group via PATCH operations (RFC 7644 Section 3.5.2). Supports add, remove, and replace operations on `members`, `displayName`, and `externalId`. For system groups (prefixed `cloudflare-v1-`), only member management operations are supported.

## Definition

```yaml
{"operationId": "scim-groups-patch", "summary": "Patch SCIM Group", "description": "Partially updates a SCIM Group via PATCH operations (RFC 7644 Section 3.5.2). Supports add, remove, and replace operations on `members`, `displayName`, and `externalId`. For system groups (prefixed `cloudflare-v1-`), only member management operations are supported.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_scim_group_identifier"}}], "requestBody": {"required": true, "content": {"application/scim+json": {"examples": {"add-members": {"summary": "Add members to group", "value": {"Operations": [{"op": "add", "path": "members", "value": [{"value": "023e105f4ecef8ad9ca31a8372d0c353"}]}], "schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}, "remove-member": {"summary": "Remove member from group", "value": {"Operations": [{"op": "remove", "path": "members[value eq \"023e105f4ecef8ad9ca31a8372d0c353\"]"}], "schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}, "replace-members": {"summary": "Replace all members", "value": {"Operations": [{"op": "replace", "path": "members", "value": [{"value": "user-tag-1"}, {"value": "user-tag-2"}]}], "schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}, "update-displayname": {"summary": "Update group name", "value": {"Operations": [{"op": "replace", "path": "displayName", "value": "New Group Name"}], "schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"]}}}, "schema": {"$ref": "#/components/schemas/iam_scim_group_patch_op_request"}}}}, "responses": {"200": {"description": "Patch SCIM Group response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_group"}}}}, "4XX": {"description": "Patch SCIM Group response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.update"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
