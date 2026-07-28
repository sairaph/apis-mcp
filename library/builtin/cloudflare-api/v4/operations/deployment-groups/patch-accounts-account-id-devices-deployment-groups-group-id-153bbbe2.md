---
title: Update deployment group
page_id: operation-patch-accounts-account-id-devices-deployment-groups-group-id-5ea4cb50
path: operations/deployment-groups
description: Updates a deployment group. Returns 409 if any newly added policy IDs already belong to another deployment group. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/devices/deployment-groups/{group_id}
operation_ids:
    - update-deployment-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update deployment group

`PATCH /accounts/{account_id}/devices/deployment-groups/{group_id}`

Operation ID: `update-deployment-group`

Updates a deployment group. Returns 409 if any newly added policy IDs already belong to another deployment group. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "update-deployment-group", "summary": "Update deployment group", "description": "Updates a deployment group. Returns 409 if any newly added policy IDs already belong to another deployment group. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_deployment_group_update_request"}}}}, "responses": {"200": {"description": "Updates deployment group response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_deployment_group"}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Deployment Groups"], "x-auditable": true, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.deployment.groups", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
