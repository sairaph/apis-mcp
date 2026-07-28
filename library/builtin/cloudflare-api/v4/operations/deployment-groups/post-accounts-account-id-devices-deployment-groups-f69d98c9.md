---
title: Create deployment group
page_id: operation-post-accounts-account-id-devices-deployment-groups-e61292ce
path: operations/deployment-groups
description: Creates a new deployment group. Policy IDs must be unique across all deployment groups. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/deployment-groups
operation_ids:
    - create-deployment-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create deployment group

`POST /accounts/{account_id}/devices/deployment-groups`

Operation ID: `create-deployment-group`

Creates a new deployment group. Policy IDs must be unique across all deployment groups. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "create-deployment-group", "summary": "Create deployment group", "description": "Creates a new deployment group. Policy IDs must be unique across all deployment groups. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_deployment_group_create_request"}}}}, "responses": {"200": {"description": "Creates deployment group response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_deployment_group"}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Deployment Groups"], "x-auditable": true, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.deployment.groups", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
