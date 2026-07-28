---
title: Get the default device settings profile
page_id: operation-get-accounts-account-id-devices-policy-600fa562
path: operations/devices
description: Fetches the default device settings profile for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/devices/policy
operation_ids:
    - devices-get-default-device-settings-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the default device settings profile

`GET /accounts/{account_id}/devices/policy`

Operation ID: `devices-get-default-device-settings-policy`

Fetches the default device settings profile for an account.

## Definition

```yaml
{"operationId": "devices-get-default-device-settings-policy", "summary": "Get the default device settings profile", "description": "Fetches the default device settings profile for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "responses": {"200": {"description": "Get the default device settings profile response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_default_device_settings_response"}}}}, "4XX": {"description": "Get the default device settings profile response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_default_device_settings_response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Devices"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.policies.default", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
