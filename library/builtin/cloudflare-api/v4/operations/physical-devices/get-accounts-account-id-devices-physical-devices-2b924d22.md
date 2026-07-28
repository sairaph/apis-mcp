---
title: List devices
page_id: operation-get-accounts-account-id-devices-physical-devices-11c2a9de
path: operations/physical-devices
description: Lists WARP devices.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/physical-devices
operation_ids:
    - list-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List devices

`GET /accounts/{account_id}/devices/physical-devices`

Operation ID: `list-devices`

Lists WARP devices.

## Definition

```yaml
{"operationId": "list-devices", "summary": "List devices", "description": "Lists WARP devices.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_cursor"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_sort_by"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_sort_order"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_last_seen_user_email"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_seen_after"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_seen_before"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_per_page"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_search"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_active_registrations"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_id"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_last_seen_registration_policy_id"}, {"$ref": "#/components/parameters/teams-devices_devices_list_devices_param_include"}], "responses": {"200": {"description": "Returns a list of Devices.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}, "example": []}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}, "example": []}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_physical_device"}}, "result_info": {"$ref": "#/components/schemas/teams-devices_cursor_result_info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["result", "success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Physical Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.devices", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
