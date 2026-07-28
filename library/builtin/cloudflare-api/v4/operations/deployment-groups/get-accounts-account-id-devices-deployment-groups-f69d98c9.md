---
title: List deployment groups
page_id: operation-get-accounts-account-id-devices-deployment-groups-4496e84d
path: operations/deployment-groups
description: Lists all deployment groups for an account. Use deployment groups to assign target WARP client versions to specific devices. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/deployment-groups
operation_ids:
    - list-deployment-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List deployment groups

`GET /accounts/{account_id}/devices/deployment-groups`

Operation ID: `list-deployment-groups`

Lists all deployment groups for an account. Use deployment groups to assign target WARP client versions to specific devices. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "list-deployment-groups", "summary": "List deployment groups", "description": "Lists all deployment groups for an account. Use deployment groups to assign target WARP client versions to specific devices. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "page", "in": "query", "description": "The page number to return.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "The maximum number of deployment groups to return per page.", "schema": {"type": "integer", "default": 50, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "Lists deployment group response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_deployment_group"}}, "result_info": {"$ref": "#/components/schemas/teams-devices_pagination_info"}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "result_info", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Deployment Groups"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.deployment.groups", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
