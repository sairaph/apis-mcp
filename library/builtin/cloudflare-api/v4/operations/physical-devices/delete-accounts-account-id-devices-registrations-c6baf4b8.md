---
title: Delete registrations
page_id: operation-delete-accounts-account-id-devices-registrations-48606063
path: operations/physical-devices
description: Deletes a list of WARP registrations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/devices/registrations
operation_ids:
    - delete-registrations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete registrations

`DELETE /accounts/{account_id}/devices/registrations`

Operation ID: `delete-registrations`

Deletes a list of WARP registrations.

## Definition

```yaml
{"operationId": "delete-registrations", "summary": "Delete registrations", "description": "Deletes a list of WARP registrations.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "id", "in": "query", "description": "A list of registration IDs to delete.", "required": true, "schema": {"type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Delete a list of registrations response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_empty_body"}, "result_info": {"$ref": "#/components/schemas/teams-devices_cursor_result_info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Physical Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.registrations", "x-fern-sdk-method-name": "bulk-delete", "x-forge-hidden": true}
```
