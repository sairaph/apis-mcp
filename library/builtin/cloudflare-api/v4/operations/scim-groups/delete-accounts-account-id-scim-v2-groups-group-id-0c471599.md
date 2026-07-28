---
title: Delete SCIM Group
page_id: operation-delete-accounts-account-id-scim-v2-groups-group-id-0926f71a
path: operations/scim-groups
description: Deletes a SCIM Group (custom user groups only). System groups backed by Cloudflare permission groups cannot be deleted via SCIM. Returns 204 No Content on success.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/scim/v2/Groups/{group_id}
operation_ids:
    - scim-groups-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete SCIM Group

`DELETE /accounts/{account_id}/scim/v2/Groups/{group_id}`

Operation ID: `scim-groups-delete`

Deletes a SCIM Group (custom user groups only). System groups backed by Cloudflare permission groups cannot be deleted via SCIM. Returns 204 No Content on success.

## Definition

```yaml
{"operationId": "scim-groups-delete", "summary": "Delete SCIM Group", "description": "Deletes a SCIM Group (custom user groups only). System groups backed by Cloudflare permission groups cannot be deleted via SCIM. Returns 204 No Content on success.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_scim_group_identifier"}}], "responses": {"204": {"description": "Delete SCIM Group response (no content)"}, "4XX": {"description": "Delete SCIM Group response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.delete"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
