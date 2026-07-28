---
title: Get device settings profile by ID
page_id: operation-get-accounts-account-id-devices-policy-policy-id-006164c6
path: operations/devices
description: Fetches a device settings profile by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy/{policy_id}
operation_ids:
    - devices-get-device-settings-policy-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get device settings profile by ID

`GET /accounts/{account_id}/devices/policy/{policy_id}`

Operation ID: `devices-get-device-settings-policy-by-id`

Fetches a device settings profile by ID.

## Definition

```yaml
{"operationId": "devices-get-device-settings-policy-by-id", "summary": "Get device settings profile by ID", "description": "Fetches a device settings profile by ID.", "parameters": [{"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_schemas-uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get device settings profile by ID response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_device_settings_response"}}}}, "4XX": {"description": "Get device settings profile by ID response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_device_settings_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.custom", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
