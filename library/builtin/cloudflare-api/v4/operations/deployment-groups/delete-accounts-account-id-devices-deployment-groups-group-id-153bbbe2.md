---
title: Delete deployment group
page_id: operation-delete-accounts-account-id-devices-deployment-groups-group-id-0c9d6a70
path: operations/deployment-groups
description: Deletes a deployment group. Associated policies no longer apply and devices stop receiving version targets. This endpoint is in Beta.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/deployment-groups/{group_id}
operation_ids:
    - delete-deployment-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete deployment group

`DELETE /accounts/{account_id}/devices/deployment-groups/{group_id}`

Operation ID: `delete-deployment-group`

Deletes a deployment group. Associated policies no longer apply and devices stop receiving version targets. This endpoint is in Beta.

## Definition

```yaml
{"operationId": "delete-deployment-group", "summary": "Delete deployment group", "description": "Deletes a deployment group. Associated policies no longer apply and devices stop receiving version targets. This endpoint is in Beta.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Deletes deployment group response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"type": "object", "properties": {"id": {"description": "The ID of a deleted deployment group.", "type": "string", "example": "550e8400-e29b-41d4-a716-446655440000", "x-auditable": true}}}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Deployment Groups"], "x-auditable": true, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.deployment.groups", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
