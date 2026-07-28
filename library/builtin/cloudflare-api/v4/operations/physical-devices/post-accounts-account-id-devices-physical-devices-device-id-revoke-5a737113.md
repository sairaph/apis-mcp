---
title: Revoke device registrations
page_id: operation-post-accounts-account-id-devices-physical-devices-device-id-revoke-d286e13d
path: operations/physical-devices
description: Revokes all WARP registrations associated with the specified device.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/devices/physical-devices/{device_id}/revoke
operation_ids:
    - revoke-device
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke device registrations

`POST /accounts/{account_id}/devices/physical-devices/{device_id}/revoke`

Operation ID: `revoke-device`

Revokes all WARP registrations associated with the specified device.

## Definition

```yaml
{"operationId": "revoke-device", "summary": "Revoke device registrations", "description": "Revokes all WARP registrations associated with the specified device.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "device_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Revoke device registrations response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/teams-devices_v4_response_message"}}, "result": {"$ref": "#/components/schemas/teams-devices_empty_body"}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["success", "errors", "messages"]}}}}}, "security": [{"api_token": []}], "tags": ["Physical Devices"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.devices", "x-fern-sdk-method-name": "revoke", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation revokes a device's session destructively."}
```
