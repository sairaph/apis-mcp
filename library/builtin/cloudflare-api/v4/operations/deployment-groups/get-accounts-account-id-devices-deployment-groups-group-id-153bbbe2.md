---
title: Get deployment group
page_id: operation-get-accounts-account-id-devices-deployment-groups-group-id-b2227d53
path: operations/deployment-groups
description: Fetches a single deployment group by its ID. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/deployment-groups/{group_id}
operation_ids:
    - get-deployment-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get deployment group

`GET /accounts/{account_id}/devices/deployment-groups/{group_id}`

Operation ID: `get-deployment-group`

Fetches a single deployment group by its ID. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "get-deployment-group", "summary": "Get deployment group", "description": "Fetches a single deployment group by its ID. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Gets deployment group response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_deployment_group"}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Deployment Groups"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.deployment.groups", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
