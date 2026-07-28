---
title: List device settings profiles
page_id: operation-get-accounts-account-id-devices-policies-230b14ee
path: operations/devices
description: Fetches a list of the device settings profiles for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policies
operation_ids:
    - devices-list-device-settings-policies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List device settings profiles

`GET /accounts/{account_id}/devices/policies`

Operation ID: `devices-list-device-settings-policies`

Fetches a list of the device settings profiles for an account.

## Definition

```yaml
{"operationId": "devices-list-device-settings-policies", "summary": "List device settings profiles", "description": "Fetches a list of the device settings profiles for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "List device settings profiles response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_device_settings_response_collection"}}}}, "4XX": {"description": "List device settings profiles response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_device_settings_response_collection"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
