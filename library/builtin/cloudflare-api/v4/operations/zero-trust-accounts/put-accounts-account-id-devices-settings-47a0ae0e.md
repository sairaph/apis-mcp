---
title: Update device settings for a Zero Trust account
page_id: operation-put-accounts-account-id-devices-settings-99c1216e
path: operations/zero-trust-accounts
description: Updates the current device settings for a Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/devices/settings
operation_ids:
    - zero-trust-accounts-update-device-settings-for-the-zero-trust-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update device settings for a Zero Trust account

`PUT /accounts/{account_id}/devices/settings`

Operation ID: `zero-trust-accounts-update-device-settings-for-the-zero-trust-account`

Updates the current device settings for a Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-update-device-settings-for-the-zero-trust-account", "summary": "Update device settings for a Zero Trust account", "description": "Updates the current device settings for a Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/teams-devices_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings"}}}}, "responses": {"200": {"description": "Update device settings for a Zero Trust account response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings-response"}}}}, "4XX": {"description": "Update device settings for a Zero Trust account response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/teams-devices_zero-trust-account-device-settings-response"}, {"$ref": "#/components/schemas/teams-devices_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.devices.settings", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
